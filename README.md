# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/ctt/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-27 02:29:41 UTC

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
| gorm | 655.3µ ± 2% | 5.454Ki ± 0% | 62.00 ± 0% |
| gormT | 655.3µ ± 2% | 6.321Ki ± 0% | 70.00 ± 0% |
| ldb | **354.8µ ± 1%** 🏆 | **2.983Ki ± 0%** 🏆 | **61.00 ± 0%** 🏆 |
| xorm | 362.3µ ± 1% | 5.048Ki ± 0% | 133.0 ± 0% |

## Insert 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 689.1µ ± 4% | 6.605Ki ± 0% | 87.00 ± 0% |
| gormT | 667.9µ ± 3% | 7.505Ki ± 0% | 95.00 ± 0% |
| ldb | **385.1µ ± 4%** 🏆 | 10.01Ki ± 0% | 145.0 ± 0% |
| xorm | 553.0µ ± 6% | **3.993Ki ± 0%** 🏆 | **83.00 ± 0%** 🏆 |

## SelectFirst 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 189.5µ ± 1% | **4.969Ki ± 0%** 🏆 | **92.00 ± 0%** 🏆 |
| gormT | **189.2µ ± 1%** 🏆 | 5.836Ki ± 0% | 100.0 ± 0% |
| ldb | 201.7µ ± 1% | 9.679Ki ± 0% | 173.0 ± 0% |
| xorm | 383.7µ ± 1% | 6.103Ki ± 0% | 152.0 ± 0% |

## Select 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | **2.288m ± 1%** 🏆 | **320.6Ki ± 0%** 🏆 | 13.80k ± 0% |
| gormT | 2.353m ± 1% | 551.8Ki ± 0% | 13.82k ± 0% |
| ldb | 2.644m ± 3% | 596.3Ki ± 0% | **10.90k ± 0%** 🏆 |
| xorm | 3.375m ± 5% | 1.535Mi ± 1% | 35.84k ± 0% |

## Update 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 674.7µ ± 4% | 6.032Ki ± 0% | **66.00 ± 0%** 🏆 |
| gormT | 680.1µ ± 1% | 7.040Ki ± 0% | 76.00 ± 0% |
| ldb | **377.8µ ± 4%** 🏆 | 6.741Ki ± 0% | 105.0 ± 0% |
| xorm | 568.7µ ± 4% | **4.406Ki ± 0%** 🏆 | 116.0 ± 0% |

> 🏆 表示该指标的最佳性能（最小值）
<!-- BENCHMARK_RESULTS_END -->

