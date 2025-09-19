#!/usr/bin/env python3

import sys

def main():
    # 读取 benchstat 输出（从标准输入）
    lines = [line.strip() for line in sys.stdin if line.strip()]
    if not lines:
        print("无基准测试数据")
        return

    # 提取表头和数据行
    header = lines[0].split()
    data = [line.split() for line in lines[1:]]

    # 生成 Markdown 表格
    print("| " + " | ".join(header) + " |")
    print("| " + " | ".join(["---"] * len(header)) + " |")
    for row in data:
        print("| " + " | ".join(row) + " |")

if __name__ == "__main__":
    main()