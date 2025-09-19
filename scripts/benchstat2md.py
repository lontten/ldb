#!/usr/bin/env python3

import sys
import re

def main():
    lines = [line.rstrip() for line in sys.stdin if line.strip()]
    if not lines:
        print("无基准测试数据")
        return

    # 提取开头的说明信息（测试时间、环境等）
    desc_lines = []
    table_start = 0
    for i, line in enumerate(lines):
        if re.match(r'^goos:|^goarch:|^pkg:|^cpu:|^\s*benchmark', line):
            table_start = i
            break
        desc_lines.append(line)

    # 输出说明信息
    if desc_lines:
        print("\n".join(desc_lines))
        print()  # 空行分隔

    # 处理表格数据（区分键值对行和指标行）
    table_lines = lines[table_start:]
    current_header = None

    for line in table_lines:
        # 处理键值对行（如goos: linux）
        if ':' in line:
            key, value = re.split(r':\s+', line, 1)
            print(f"| {key}: | {value} |")
            print("| --- | --- |")
        # 处理指标分组行（如sec/op、B/op）
        elif re.match(r'^\s*\|?\s*benchmark_results\.txt\s*\|?', line):
            if 'sec/op' in line:
                current_header = ['benchmark', 'sec/op', '±', '波动']
            elif 'B/op' in line:
                current_header = ['benchmark', 'B/op', '±', '波动']
            elif 'allocs/op' in line:
                current_header = ['benchmark', 'allocs/op', '±', '波动']
            # 输出分组表头
            if current_header:
                print()  # 分组间空行
                print("| " + " | ".join(current_header) + " |")
                print("| " + " | ".join(["---"] * len(current_header)) + " |")
        # 处理具体数据行（如Insert-4 357.4µ ± 5%）
        elif re.match(r'^[A-Za-z0-9-]+\s+', line):
            parts = re.split(r'\s+', line.strip())
            # 确保数据字段匹配表头
            if len(parts) >= 4:
                print(f"| {parts[0]} | {parts[1]} | {parts[2]} | {parts[3]} |")
            elif len(parts) == 2:  # 处理geomean行
                print(f"| {parts[0]} | {parts[1]} | | |")

if __name__ == "__main__":
    main()