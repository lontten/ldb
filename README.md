# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/ctt/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-30 04:06:41 UTC

> 说明：数值越低性能越好，±表示波动范围

# Go 基准测试报告

## 环境信息

| 参数 | 值 |
|------|----|
| goos | linux |
| goarch | amd64 |
| pkg | github.com/lontten/ldb/v2/benchmark |
| cpu | AMD EPYC 7763 64-Core Processor |

## Select 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | **2.596m ±  1%** 🏆 | **422.2Ki ± 0%** 🏆 | **18.80k ± 0%** 🏆 |
| gormT | 2.668m ±  1% | 708.3Ki ± 0% | 18.82k ± 0% |
| ldb | 3.786m ±  2% | 1.496Mi ± 0% | 35.92k ± 0% |
| xorm | 5.049m ±  3% | 1.957Mi ± 2% | 51.84k ± 0% |

## Delete 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 686.1µ ±  4% | 5.454Ki ± 0% | 62.00 ± 0% |
| gormT | 669.9µ ±  8% | 6.321Ki ± 0% | 70.00 ± 0% |
| ldb | **362.8µ ±  5%** 🏆 | **2.983Ki ± 0%** 🏆 | **61.00 ± 0%** 🏆 |
| xorm | 367.5µ ±  1% | 5.048Ki ± 0% | 133.0 ± 0% |

## Insert 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 707.8µ ±  7% | 6.605Ki ± 0% | 87.00 ± 0% |
| gormT | 674.4µ ±  3% | 7.505Ki ± 0% | 95.00 ± 0% |
| ldb | **369.3µ ± 12%** 🏆 | 10.00Ki ± 0% | 145.0 ± 0% |
| xorm | 545.4µ ±  6% | **3.993Ki ± 0%** 🏆 | **83.00 ± 0%** 🏆 |

## First 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 191.5µ ±  1% | **4.970Ki ± 0%** 🏆 | **92.00 ± 0%** 🏆 |
| gormT | **190.4µ ±  2%** 🏆 | 5.836Ki ± 0% | 100.0 ± 0% |
| ldb | 209.5µ ±  1% | 10.01Ki ± 0% | 178.0 ± 0% |
| xorm | 382.2µ ±  1% | 6.103Ki ± 0% | 152.0 ± 0% |

## SelectNuller 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | **2.358m ±  1%** 🏆 | **320.6Ki ± 0%** 🏆 | **13.80k ± 0%** 🏆 |
| gormT | 2.461m ±  0% | 551.8Ki ± 0% | 13.82k ± 0% |
| ldb | 2.956m ±  1% | 831.4Ki ± 0% | 18.91k ± 0% |
| xorm | 3.419m ±  2% | 1.526Mi ± 6% | 35.84k ± 0% |

## Update 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 686.7µ ±  6% | 8.221Ki ± 0% | 93.00 ± 0% |
| gormT | 687.8µ ±  6% | 7.604Ki ± 0% | **75.00 ± 0%** 🏆 |
| ldb | **350.4µ ±  6%** 🏆 | 6.741Ki ± 0% | 105.0 ± 0% |
| xorm | 557.1µ ±  9% | **4.407Ki ± 0%** 🏆 | 116.0 ± 0% |

> 🏆 表示该指标的最佳性能（最小值）
<!-- BENCHMARK_RESULTS_END -->

