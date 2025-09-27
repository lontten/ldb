# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/ctt/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-27 02:10:33 UTC

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
| gorm | 613.9µ ± 2% | 5.423Ki ± 0% | 62.00 ± 0% |
| gormT | 615.0µ ± 2% | 6.290Ki ± 0% | 70.00 ± 0% |
| ldb | **312.8µ ± 4%** 🏆 | **2.952Ki ± 0%** 🏆 | **61.00 ± 0%** 🏆 |
| xorm | 354.7µ ± 1% | 4.040Ki ± 0% | 106.0 ± 0% |

## Insert 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 630.9µ ± 2% | 5.936Ki ± 0% | 81.00 ± 0% |
| gormT | 635.0µ ± 2% | 6.834Ki ± 0% | 89.00 ± 0% |
| ldb | **333.3µ ± 4%** 🏆 | 7.633Ki ± 0% | 124.0 ± 0% |
| xorm | 491.8µ ± 4% | **3.359Ki ± 0%** 🏆 | **71.00 ± 0%** 🏆 |

## SelectFirst 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 186.4µ ± 1% | **4.497Ki ± 0%** 🏆 | **80.00 ± 0%** 🏆 |
| gormT | **185.9µ ± 1%** 🏆 | 5.371Ki ± 0% | 88.00 ± 0% |
| ldb | 195.6µ ± 1% | 7.669Ki ± 0% | 143.0 ± 0% |
| xorm | 371.4µ ± 1% | 4.711Ki ± 0% | 122.0 ± 0% |

## Select 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | **1.653m ± 2%** 🏆 | **253.5Ki ± 0%** 🏆 | 12.79k ± 0% |
| gormT | 1.692m ± 1% | 399.7Ki ± 0% | 12.81k ± 0% |
| ldb | 1.774m ± 0% | 395.5Ki ± 0% | **9.872k ± 0%** 🏆 |
| xorm | 2.638m ± 3% | 1.066Mi ± 1% | 28.83k ± 0% |

## Update 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 633.4µ ± 0% | 6.001Ki ± 0% | **66.00 ± 0%** 🏆 |
| gormT | 638.4µ ± 1% | 7.009Ki ± 0% | 76.00 ± 0% |
| ldb | **332.1µ ± 2%** 🏆 | 5.054Ki ± 0% | 93.00 ± 0% |
| xorm | 509.3µ ± 5% | **3.703Ki ± 0%** 🏆 | 100.0 ± 0% |

> 🏆 表示该指标的最佳性能（最小值）
<!-- BENCHMARK_RESULTS_END -->

