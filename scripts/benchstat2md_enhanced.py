#!/usr/bin/env python3
"""
将 benchstat 输出转换为格式良好的 Markdown 表格
支持多指标类型（sec/op, B/op, allocs/op）
并按操作类型组织数据，比较不同实现
为每个指标的最小值添加明显的高亮显示
"""

import sys
import re

def parse_benchstat(input_text):
    """解析多指标 benchstat 输出"""
    sections = []
    current_section = {}
    lines = input_text.strip().split('\n')

    # 提取环境信息
    env_info = {}
    for line in lines[:4]:
        if ':' in line:
            key, value = line.split(':', 1)
            env_info[key.strip()] = value.strip()

    # 解析各个指标部分
    i = 4  # 跳过前4行环境信息
    while i < len(lines):
        line = lines[i].strip()

        # 检测新部分的开始
        if line.startswith('│') and 'benchmark_results.txt' in line:
            if current_section:  # 保存前一个部分
                sections.append(current_section)

            current_section = {}
            i += 1
            continue

        # 解析指标标题行
        if line.startswith('│') and any(metric in line for metric in ['sec/op', 'B/op', 'allocs/op']):
            metric_match = re.search(r'│\s+([^│]+)\s+│', line)
            if metric_match:
                current_section['metric'] = metric_match.group(1).strip()
            i += 1
            continue

        # 解析数据行 - 改进的解析方法，更好地处理各种单位
        if line and not line.startswith('│') and not line.startswith('─') and not line.startswith('geomean'):
            # 使用更灵活的正则表达式匹配
            # 匹配模式：名称 + 空格 + 数值 + 单位 + 误差
            match = re.match(r'^([\w_\-]+)\s+([\d\.]+[µnmkKMGi]?[B]?)\s*([±]\s*[\d\.%]+)?', line)
            if not match:
                # 尝试另一种模式：名称 + 多个空格 + 任意内容（作为值）
                match = re.match(r'^([\w_\-]+)\s{2,}(.+)', line)

            if match:
                if 'data' not in current_section:
                    current_section['data'] = []

                test_name = match.group(1)
                # 获取完整的值部分
                value_part = line[len(test_name):].strip()

                current_section['data'].append({
                    'name': test_name,
                    'value': value_part
                })

        i += 1

    if current_section:  # 添加最后一个部分
        sections.append(current_section)

    return env_info, sections

def organize_data_by_operation(sections):
    """按操作类型组织数据，比较不同实现"""
    operation_data = {}

    for section in sections:
        metric = section.get('metric', '')
        data = section.get('data', [])

        for item in data:
            test_name = item['name']
            value = item['value']

            # 提取操作类型和实现（如 Insert_ldb -> Insert, ldb）
            if '_' in test_name:
                operation, implementation = test_name.split('_', 1)
                # 去掉线程数后缀（如 -4）
                implementation = re.sub(r'-\d+$', '', implementation)
            else:
                # 如果没有下划线，使用默认实现名称
                operation = re.sub(r'-\d+$', '', test_name)
                implementation = "default"

            if operation not in operation_data:
                operation_data[operation] = {}

            if implementation not in operation_data[operation]:
                operation_data[operation][implementation] = {}

            operation_data[operation][implementation][metric] = value

    return operation_data

def extract_numeric_value(value_str):
    """从字符串中提取数值部分，用于比较，正确处理各种单位"""
    # 提取数值和单位部分
    match = re.match(r'([\d\.]+)\s*([µnmkKMGi]?[B]?)', value_str)
    if not match:
        return float('inf')

    numeric_part = match.group(1)
    unit_part = match.group(2)

    try:
        value = float(numeric_part)
    except ValueError:
        return float('inf')

    # 根据单位进行换算
    unit_multipliers = {
        'n': 1e-9,  # 纳秒
        'µ': 1e-6,  # 微秒
        'm': 1e-3,  # 毫秒
        'k': 1e3,   # 千
        'K': 1e3,   # 千
        'M': 1e6,   # 兆
        'G': 1e9,   # 吉
        'Ki': 1024, # 千字节（二进制）
        'Mi': 1024 * 1024,  # 兆字节
        'Gi': 1024 * 1024 * 1024,  # 吉字节
    }

    # 如果是B/op指标，使用二进制单位换算
    if 'B' in value_str:
        if 'Ki' in value_str:
            value *= unit_multipliers['Ki']
        elif 'Mi' in value_str:
            value *= unit_multipliers['Mi']
        elif 'Gi' in value_str:
            value *= unit_multipliers['Gi']
        elif 'k' in value_str or 'K' in value_str:
            value *= unit_multipliers['k']
        elif 'M' in value_str:
            value *= unit_multipliers['M']
        elif 'G' in value_str:
            value *= unit_multipliers['G']
    else:
        # 对于时间指标，使用标准单位换算
        if unit_part in unit_multipliers:
            value *= unit_multipliers[unit_part]

    return value

def format_markdown_table(env_info, operation_data):
    """生成格式化的 Markdown 表格，并为最小值添加明显的高亮"""
    if not operation_data:
        return "没有可用的基准测试数据"

    # 创建环境信息表
    md_output = "# Go 基准测试报告\n\n"
    md_output += "## 环境信息\n\n"
    md_output += "| 参数 | 值 |\n"
    md_output += "|------|----|\n"
    for key, value in env_info.items():
        md_output += f"| {key} | {value} |\n"
    md_output += "\n"

    # 为每个操作类型创建表格
    for operation, implementations in operation_data.items():
        md_output += f"## {operation} 操作性能比较\n\n"
        md_output += "| 实现 | sec/op | B/op | allocs/op |\n"
        md_output += "|------|--------|------|-----------|\n"

        # 获取所有实现名称并按字母顺序排序
        impl_names = sorted(implementations.keys())

        # 找出每个指标的最小值
        min_sec = float('inf')
        min_b = float('inf')
        min_allocs = float('inf')

        sec_values = {}
        b_values = {}
        allocs_values = {}

        for impl_name in impl_names:
            metrics = implementations[impl_name]

            # 提取数值用于比较
            if 'sec/op' in metrics:
                sec_val = extract_numeric_value(metrics['sec/op'])
                sec_values[impl_name] = sec_val
                min_sec = min(min_sec, sec_val)

            if 'B/op' in metrics:
                b_val = extract_numeric_value(metrics['B/op'])
                b_values[impl_name] = b_val
                min_b = min(min_b, b_val)

            if 'allocs/op' in metrics:
                allocs_val = extract_numeric_value(metrics['allocs/op'])
                allocs_values[impl_name] = allocs_val
                min_allocs = min(min_allocs, allocs_val)

        # 输出表格行，为最小值添加明显标记
        for impl_name in impl_names:
            metrics = implementations[impl_name]
            sec_op = metrics.get('sec/op', 'N/A')
            b_op = metrics.get('B/op', 'N/A')
            allocs_op = metrics.get('allocs/op', 'N/A')

            # 检查是否为最小值并添加明显标记
            if impl_name in sec_values and sec_values[impl_name] == min_sec:
                sec_op = f"**{sec_op}** 🏆"
            if impl_name in b_values and b_values[impl_name] == min_b:
                b_op = f"**{b_op}** 🏆"
            if impl_name in allocs_values and allocs_values[impl_name] == min_allocs:
                allocs_op = f"**{allocs_op}** 🏆"

            md_output += f"| {impl_name} | {sec_op} | {b_op} | {allocs_op} |\n"

        md_output += "\n"

    # 添加图例说明
    md_output += "> 🏆 表示该指标的最佳性能（最小值）\n\n"
    return md_output

def main():
    # 读取标准输入
    input_text = sys.stdin.read()

    # 解析 benchstat 输出
    env_info, sections = parse_benchstat(input_text)

    if not sections:
        print("未能解析基准测试数据")
        return

    # 按操作类型组织数据
    operation_data = organize_data_by_operation(sections)

    # 生成 Markdown 报告
    md_output = format_markdown_table(env_info, operation_data)
    print(md_output)

if __name__ == "__main__":
    main()