#!/usr/bin/env python3
"""
将 benchstat 输出转换为格式良好的 Markdown 表格
支持多指标类型（sec/op, B/op, allocs/op）
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

        # 解析数据行
        if line and not line.startswith('│') and not line.startswith('─') and not line.startswith('geomean'):
            parts = re.split(r'\s{2,}', line)
            if len(parts) >= 2:
                if 'data' not in current_section:
                    current_section['data'] = []

                test_name = parts[0]
                value = parts[1] if len(parts) > 1 else ""
                current_section['data'].append({
                    'name': test_name,
                    'value': value
                })

        i += 1

    if current_section:  # 添加最后一个部分
        sections.append(current_section)

    return env_info, sections

def format_markdown_table(env_info, sections):
    """生成格式化的 Markdown 表格"""
    if not sections:
        return "没有可用的基准测试数据"

    # 创建环境信息表
    md_output = "# Go 基准测试报告\n\n"
    md_output += "## 环境信息\n\n"
    md_output += "| 参数 | 值 |\n"
    md_output += "|------|----|\n"
    for key, value in env_info.items():
        md_output += f"| {key} | {value} |\n"
    md_output += "\n"

    # 为每个指标部分创建表格
    for section in sections:
        metric = section.get('metric', '未知指标')
        data = section.get('data', [])

        md_output += f"## {metric}\n\n"
        md_output += "| 测试名称 | 值 |\n"
        md_output += "|----------|----|\n"

        for item in data:
            md_output += f"| {item['name']} | {item['value']} |\n"

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

    # 生成 Markdown 报告
    md_output = format_markdown_table(env_info, sections)
    print(md_output)

if __name__ == "__main__":
    main()