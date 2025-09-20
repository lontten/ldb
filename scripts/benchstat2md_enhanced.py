#!/usr/bin/env python3
"""
将 benchstat 输出转换为格式良好的 Markdown 表格
支持多指标类型（sec/op, B/op, allocs/op）
并按操作类型组织数据，比较不同实现
"""

import sys
import re

def parse_benchstat(input_text):
    """解析多指标 benchstat 输出"""
    sections = []
    lines = input_text.strip().split('\n')

    # 提取环境信息
    env_info = {}
    for line in lines[:4]:
        if ':' in line:
            key, value = line.split(':', 1)
            env_info[key.strip()] = value.strip()

    # 解析基准测试数据部分
    current_section = {}
    i = 4  # 跳过前4行环境信息

    while i < len(lines):
        line = lines[i].strip()

        # 检测新部分的开始（指标标题行）
        if line.startswith('name') and any(metric in line for metric in ['sec/op', 'B/op', 'allocs/op']):
            if current_section:  # 保存前一个部分
                sections.append(current_section)

            # 提取指标类型
            metric_match = re.search(r'(\S+/op)', line)
            current_section = {
                'metric': metric_match.group(1) if metric_match else 'unknown',
                'data': []
            }
            i += 2  # 跳过标题行和分隔行
            continue

        # 解析数据行
        if line and not line.startswith('goos:') and not line.startswith('goarch:'):
            # 使用正则表达式匹配测试名称和值
            parts = re.split(r'\s+', line)
            if len(parts) >= 2:
                test_name = parts[0]
                # 合并剩余部分作为值（可能包含空格，如 "10.5 ± 2%"）
                value = ' '.join(parts[1:])

                current_section['data'].append({
                    'name': test_name,
                    'value': value
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

            # 提取操作类型和实现
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

def format_markdown_table(env_info, operation_data):
    """生成格式化的 Markdown 表格"""
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

        for impl_name, metrics in implementations.items():
            sec_op = metrics.get('sec/op', 'N/A')
            b_op = metrics.get('B/op', 'N/A')
            allocs_op = metrics.get('allocs/op', 'N/A')

            md_output += f"| {impl_name} | {sec_op} | {b_op} | {allocs_op} |\n"

        md_output += "\n"

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