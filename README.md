# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/ctt/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-27 01:48:44 UTC

> 说明：数值越低性能越好，±表示波动范围

# Go 基准测试报告

## 环境信息

| 参数 | 值 |
|------|----|
| goos | linux |
| goarch | amd64 |
| pkg | github.com/lontten/ldb/v2/benchmark |
| cpu | AMD EPYC 7763 64-Core Processor |

## Delete 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 774.0µ ±   ∞ ¹ | 5.423Ki ±  ∞ ¹ | 62.00 ±  ∞ ¹ |
| gormT | 765.9µ ±  5% | 6.290Ki ± 0% | 70.00 ± 0% |
| ldb | 478.2µ ± 13% | **2.952Ki ± 0%** 🏆 | **61.00 ± 0%** 🏆 |
| xorm | **346.2µ ±  1%** 🏆 | 4.040Ki ± 0% | 106.0 ± 0% |

## Insert 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 740.3µ ±  6% | 5.936Ki ± 0% | 81.00 ± 0% |
| gormT | 760.9µ ±  6% | 6.834Ki ± 0% | 89.00 ± 0% |
| ldb | **482.2µ ±  5%** 🏆 | 7.633Ki ± 0% | 124.0 ± 1% |
| xorm | 620.9µ ± 13% | **3.359Ki ± 0%** 🏆 | **71.00 ± 0%** 🏆 |

## SelectFirst 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 185.4µ ±  0% | **4.496Ki ± 0%** 🏆 | **80.00 ± 0%** 🏆 |
| gormT | **183.7µ ±  1%** 🏆 | 5.371Ki ± 0% | 88.00 ± 0% |
| ldb | 194.8µ ±  2% | 7.669Ki ± 0% | 143.0 ± 0% |
| xorm | 364.8µ ±  1% | 4.711Ki ± 0% | 122.0 ± 0% |

## Select 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | **1.643m ±  1%** 🏆 | **253.5Ki ± 0%** 🏆 | 12.79k ± 0% |
| gormT | 1.685m ±  2% | 399.7Ki ± 0% | 12.81k ± 0% |
| ldb | 1.763m ±  1% | 395.5Ki ± 0% | **9.872k ± 0%** 🏆 |
| xorm | 2.642m ±  2% | 1.056Mi ± 5% | 28.83k ± 0% |

## Update 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 741.6µ ±  7% | 6.001Ki ± 0% | **66.00 ± 0%** 🏆 |
| gormT | 771.6µ ±  6% | 7.009Ki ± 0% | 76.00 ± 0% |
| ldb | **476.2µ ±  8%** 🏆 | 5.054Ki ± 0% | 93.00 ± 0% |
| xorm | 686.9µ ±  2% | **3.703Ki ± 0%** 🏆 | 100.0 ± 0% |

> 🏆 表示该指标的最佳性能（最小值）
<!-- BENCHMARK_RESULTS_END -->

