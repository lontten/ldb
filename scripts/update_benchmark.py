#!/usr/bin/env python3
import datetime
import os

def generate_benchmark_section():
    """生成基准测试结果段落并保存到文件"""
    # 获取当前UTC时间
    current_utc_time = datetime.datetime.utcnow().strftime("%Y-%m-%d %H:%M:%S UTC")

    # 创建基准测试部分内容
    section_lines = [
        "## 最新基准测试结果",
        "",
        f"测试时间: {current_utc_time}",
        "",
        "> 说明：数值越低性能越好，±表示波动范围",
        ""
    ]

    # 读取benchmark_summary.md内容
    if os.path.exists("benchmark_summary.md"):
        with open("benchmark_summary.md", "r", encoding="utf-8") as f:
            section_lines.extend(f.readlines())
    else:
        section_lines.append("警告：未找到benchmark_summary.md文件")

    # 写入到benchmark_section.md
    with open("benchmark_section.md", "w", encoding="utf-8") as f:
        f.write("\n".join(section_lines))

def update_readme():
    """更新README.md中的基准测试结果部分"""
    if not os.path.exists("README.md"):
        print("错误：未找到README.md文件")
        return

    # 读取README.md内容
    with open("README.md", "r", encoding="utf-8") as f:
        lines = f.readlines()

    # 查找标记位置
    start_idx = None
    end_idx = None
    for i, line in enumerate(lines):
        if "<!-- BENCHMARK_RESULTS_START -->" in line:
            start_idx = i
        if "<!-- BENCHMARK_RESULTS_END -->" in line:
            end_idx = i

    if start_idx is None or end_idx is None:
        print("错误：在README.md中未找到完整的标记")
        return

    # 读取要插入的内容
    if not os.path.exists("benchmark_section.md"):
        print("错误：未找到benchmark_section.md文件")
        return

    with open("benchmark_section.md", "r", encoding="utf-8") as f:
        new_content = f.readlines()

    # 构建新的内容列表
    new_lines = []
    # 添加开始标记前的内容
    new_lines.extend(lines[:start_idx + 1])
    # 添加新的基准测试内容
    new_lines.extend(new_content)
    # 添加结束标记及之后的内容
    new_lines.extend(lines[end_idx:])

    # 写回README.md
    with open("README.md", "w", encoding="utf-8") as f:
        f.writelines(new_lines)

def main():
    generate_benchmark_section()
    update_readme()
    print("基准测试结果已成功更新到README.md")

if __name__ == "__main__":
    main()
