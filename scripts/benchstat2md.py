#!/usr/bin/env python3

import sys
import re

def main():
    # 读取输入并过滤空行
    lines = [line.rstrip() for line in sys.stdin if line.strip()]
    if not lines:
        print("无基准测试数据")
        return

    # 存储测试用例数据：{测试用例: {指标类型: 数值}}
    # 指标类型：sec/op（耗时）、B/op（内存）、allocs/op（分配次数）
    test_data = {}
    current_metric = None  # 当前处理的指标类型

    # 遍历所有行，提取数据
    for line in lines:
        # 识别指标组标题（sec/op、B/op、allocs/op）
        if 'sec/op' in line:
            current_metric = 'sec/op'
        elif 'B/op' in line:
            current_metric = 'B/op'
        elif 'allocs/op' in line:
            current_metric = 'allocs/op'
        # 处理测试用例行（如 Insert-4、Select-4）
        elif re.match(r'^(Insert|Select)-\d+', line) or line.startswith('geomean'):
            # 用正则分割字段（支持多空格和±符号）
            parts = re.split(r'\s+', line.strip())
            test_case = parts[0]  # 测试用例名称（如 Insert-4）
            value = ' '.join(parts[1:])  # 指标值（如 323.0µ ± 11%）

            # 初始化测试用例数据
            if test_case not in test_data:
                test_data[test_case] = {}
            # 存储当前指标值
            test_data[test_case][current_metric] = value

    # 按测试用例输出结果（只保留 Insert-4、Select-4，跳过 geomean）
    for test_case in ['Insert-4', 'Select-4']:
        if test_case not in test_data:
            continue
        data = test_data[test_case]
        # 输出格式：测试用例标题 + 耗时 | 内存 | 分配次数
        print(f"### {test_case}")
        print(f"{data.get('sec/op', 'N/A')}     {data.get('B/op', 'N/A')}     {data.get('allocs/op', 'N/A')}")
        print()  # 空行分隔

if __name__ == "__main__":
    main()