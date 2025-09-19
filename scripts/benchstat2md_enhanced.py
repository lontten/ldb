#!/usr/bin/env python3
"""
将 benchstat 输出转换为格式良好的 Markdown 表格
"""

import sys
import re

def parse_benchstat():
    """解析 benchstat 输出"""
    lines = [line.rstrip() for line in sys.stdin if line.strip()]

    # 检测是否有基准测试数据
    if not lines or "time/op" not in lines[0]:
        return None, None, None

    # 提取表头
    headers = [h.strip() for h in lines[0].split() if h.strip()]

    # 解析数据行
    data = []
    current_package = None

    for line in lines[1:]:
        if not line.strip():
            continue

        # 检测包名行
        if line.startswith(' ') or line.startswith('\t'):
            # 这是数据行
            parts = re.split(r'\s{2,}', line.strip())
            if len(parts) >= 4:
                row = {"package": current_package}
                row["name"] = parts[0]
                row["time/op"] = parts[1]
                row["delta"] = parts[2] if len(parts) > 2 else ""
                row["pval"] = parts[3] if len(parts) > 3 else ""
                data.append(row)
        else:
            # 这是包名行
            current_package = line.strip()

    return headers, data, current_package

def format_markdown_table(headers, data):
    """生成格式化的 Markdown 表格"""
    if not headers or not data:
        return "没有可用的基准测试数据"

    # 创建表格头
    md_table = "| 包名 | 测试名称 | 时间/操作 | 变化 | P值 |\n"
    md_table += "|------|----------|-----------|------|-----|\n"

    # 填充数据行
    for row in data:
        md_table += f"| {row.get('package', '')} | {row.get('name', '')} | {row.get('time/op', '')} | {row.get('delta', '')} | {row.get('pval', '')} |\n"

    return md_table

def format_detailed_report(data):
    """生成详细报告"""
    report = "## 基准测试详细分析\n\n"

    # 按包分组
    packages = {}
    for row in data:
        pkg = row.get('package', 'unknown')
        if pkg not in packages:
            packages[pkg] = []
        packages[pkg].append(row)

    # 为每个包生成分析
    for pkg, tests in packages.items():
        report += f"### {pkg}\n\n"
        report += "| 测试名称 | 时间/操作 | 变化 | 显著性 |\n"
        report += "|----------|-----------|------|--------|\n"

        for test in tests:
            # 添加表情符号表示性能变化
            delta = test.get('delta', '')
            trend = "➡️"
            if delta.startswith('-'):
                trend = "✅"
            elif delta.startswith('+'):
                trend = "❌"

            report += f"| {test.get('name', '')} | {test.get('time/op', '')} | {delta} | {trend} |\n"

        report += "\n"

    return report

def main():
    # 读取标准输入
    input_text = sys.stdin.read()

    # 解析 benchstat 输出
    headers, data, pkg = parse_benchstat(input_text)

    if not data:
        print("未能解析基准测试数据")
        return

    # 生成 Markdown 报告
    md_output = "# Go 基准测试报告\n\n"
    md_output += "## 测试概览\n\n"
    md_output += format_markdown_table(headers, data)
    md_output += "\n"
    md_output += format_detailed_report(data)

    # 添加总结部分
    md_output += "## 总结\n\n"
    md_output += "- ✅ 表示性能提升\n"
    md_output += "- ❌ 表示性能下降\n"
    md_output += "- ➡️ 表示无明显变化\n"

    print(md_output)

if __name__ == "__main__":
    main()