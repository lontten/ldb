#!/usr/bin/env python3

import sys
import re

def main():
    # 读取输入并过滤空行
    lines = [line.rstrip() for line in sys.stdin if line.strip()]
    if not lines:
        print("无基准测试数据")
        return

    # 1. 处理系统信息（goos、goarch、pkg、cpu）
    sys_info = []
    data_start = 0
    for i, line in enumerate(lines):
        # 匹配系统信息行（如 goos: linux）
        if re.match(r'^(goos|goarch|pkg|cpu):\s+', line):
            key, value = re.split(r':\s+', line, 1)
            sys_info.append((key, value))
        else:
            data_start = i  # 找到数据开始的位置
            break

    # 输出系统信息表格
    if sys_info:
        print("### 测试环境信息")
        for key, value in sys_info:
            print(f"| {key}: | {value} |")
            print("| --- | --- |")
        print()  # 空行分隔

    # 2. 处理指标组（sec/op、B/op、allocs/op）
    # 分组标识和表头映射
    groups = {
        'sec/op': '每次操作耗时',
        'B/op': '每次操作内存分配',
        'allocs/op': '每次操作内存分配次数'
    }
    current_group = None
    group_data = []

    # 遍历数据行（从系统信息之后开始）
    for line in lines[data_start:]:
        # 检测分组标题行（如 "│        sec/op         │"）
        for marker, title in groups.items():
            if marker in line:
                # 如果之前有分组数据，先输出
                if current_group and group_data:
                    print_group(current_group, group_data)
                    group_data = []
                current_group = title
                break
        else:
            # 非标题行，收集数据（过滤空行和分隔符）
            if line.strip() and not re.match(r'^\s*│', line):
                group_data.append(line.strip())

    # 输出最后一个分组
    if current_group and group_data:
        print_group(current_group, group_data)

def print_group(title, data_lines):
    """输出一个指标组的Markdown表格"""
    print(f"### {title}")
    # 表头
    print("| 基准测试 | 数值 | ± | 波动范围 |")
    print("| -------- | ---- | - | -------- |")
    # 处理每行数据
    for line in data_lines:
        # 用正则分割（支持多空格分隔）
        parts = re.split(r'\s+', line)
        if len(parts) == 2:
            # geomean行（没有±和波动范围）
            print(f"| {parts[0]} | {parts[1]} | | |")
        elif len(parts) >= 4:
            # 普通行（包含±和波动范围）
            print(f"| {parts[0]} | {parts[1]} | {parts[2]} | {parts[3]} |")

if __name__ == "__main__":
    main()