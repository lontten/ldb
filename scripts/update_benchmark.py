#!/usr/bin/env python3
import datetime
import os
import sys

def get_project_root():
    """获取项目根目录（脚本所在目录的父目录）"""
    script_dir = os.path.dirname(os.path.abspath(__file__))
    return os.path.dirname(script_dir)  # 返回scripts目录的父目录

def generate_benchmark_section():
    """生成基准测试结果段落并保存到文件"""
    project_root = get_project_root()

    # 修复弃用警告：使用时区感知的UTC时间
    current_utc_time = datetime.datetime.now(datetime.timezone.utc).strftime("%Y-%m-%d %H:%M:%S UTC")

    # 创建基准测试部分内容
    section_lines = [
        "## 最新基准测试结果",
        "",
        f"测试时间: {current_utc_time}",
        "",
        "> 说明：数值越低性能越好，±表示波动范围",
        ""
    ]

    # 读取benchmark_summary.md内容（在项目根目录）
    benchmark_summary_path = os.path.join(project_root, "benchmark_summary.md")
    try:
        if os.path.exists(benchmark_summary_path):
            with open(benchmark_summary_path, "r", encoding="utf-8") as f:
                content = f.read().strip()
                if content:
                    section_lines.append(content)
                else:
                    section_lines.append("基准测试数据为空")
        else:
            section_lines.append("警告：未找到 benchmark_summary.md 文件")
            print(f"警告：文件 {benchmark_summary_path} 不存在")
    except Exception as e:
        error_msg = f"读取基准测试摘要文件时出错: {str(e)}"
        section_lines.append(error_msg)
        print(error_msg)

    # 写入到benchmark_section.md（也在项目根目录）
    benchmark_section_path = os.path.join(project_root, "benchmark_section.md")
    try:
        with open(benchmark_section_path, "w", encoding="utf-8") as f:
            # 使用统一的换行符
            f.write("\n".join(section_lines))
        print(f"基准测试部分已生成: {benchmark_section_path}")
    except Exception as e:
        print(f"错误：写入 benchmark_section.md 失败: {str(e)}")
        sys.exit(1)

def update_readme():
    """更新README.md中的基准测试结果部分"""
    project_root = get_project_root()
    readme_path = os.path.join(project_root, "README.md")

    if not os.path.exists(readme_path):
        print(f"错误：未找到 README.md 文件 ({readme_path})")
        return False

    # 读取README.md内容
    try:
        with open(readme_path, "r", encoding="utf-8") as f:
            content = f.read()
    except Exception as e:
        print(f"错误：读取 README.md 失败: {str(e)}")
        return False

    # 检查标记是否存在
    if "<!-- BENCHMARK_RESULTS_START -->" not in content:
        print("错误：在 README.md 中未找到 <!-- BENCHMARK_RESULTS_START --> 标记")
        return False
    if "<!-- BENCHMARK_RESULTS_END -->" not in content:
        print("错误：在 README.md 中未找到 <!-- BENCHMARK_RESULTS_END --> 标记")
        return False

    # 读取要插入的内容
    benchmark_section_path = os.path.join(project_root, "benchmark_section.md")
    if not os.path.exists(benchmark_section_path):
        print(f"错误：未找到 benchmark_section.md 文件 ({benchmark_section_path})")
        return False

    try:
        with open(benchmark_section_path, "r", encoding="utf-8") as f:
            new_content = f.read().strip()
    except Exception as e:
        print(f"错误：读取 benchmark_section.md 失败: {str(e)}")
        return False

    # 构建新的内容
    start_marker = "<!-- BENCHMARK_RESULTS_START -->"
    end_marker = "<!-- BENCHMARK_RESULTS_END -->"

    start_index = content.find(start_marker)
    end_index = content.find(end_marker)

    if start_index == -1 or end_index == -1 or end_index <= start_index:
        print("错误：标记位置无效")
        return False

    # 包含标记本身
    new_content_full = f"{start_marker}\n{new_content}\n{end_marker}"
    before_content = content[:start_index]
    after_content = content[end_index + len(end_marker):]

    final_content = before_content + new_content_full + after_content

    # 写回README.md
    try:
        with open(readme_path, "w", encoding="utf-8") as f:
            f.write(final_content)
        print(f"README.md 已成功更新: {readme_path}")
        return True
    except Exception as e:
        print(f"错误：写入 README.md 失败: {str(e)}")
        return False

def main():
    """主函数"""
    print("开始更新基准测试结果...")

    # 生成基准测试部分
    generate_benchmark_section()

    # 更新README
    success = update_readme()

    if success:
        print("基准测试结果已成功更新到 README.md")
    else:
        print("更新 README.md 失败")
        sys.exit(1)

if __name__ == "__main__":
    main()