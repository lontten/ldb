# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/ctt/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-29 03:27:38 UTC

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
| gorm | **2.574m ±  1%** 🏆 | **422.2Ki ± 0%** 🏆 | **18.80k ± 0%** 🏆 |
| gormT | 2.647m ±  1% | 708.3Ki ± 0% | 18.82k ± 0% |
| xorm | 5.032m ±  4% | 1.961Mi ± 1% | 51.84k ± 0% |

## Delete 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 688.6µ ±  3% | 5.454Ki ± 0% | 62.00 ± 0% |
| gormT | 697.0µ ±  1% | 6.322Ki ± 0% | 70.00 ± 0% |
| ldb | 383.3µ ± 12% | **2.983Ki ± 0%** 🏆 | **61.00 ± 0%** 🏆 |
| xorm | **361.2µ ±  1%** 🏆 | 5.048Ki ± 0% | 133.0 ± 0% |

## Insert 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 725.4µ ±  5% | 6.606Ki ± 0% | 87.00 ± 0% |
| gormT | 733.9µ ±  4% | 7.505Ki ± 0% | 95.00 ± 0% |
| ldb | **407.2µ ± 11%** 🏆 | 10.08Ki ± 0% | 148.0 ± 0% |
| xorm | 596.8µ ±  6% | **3.993Ki ± 0%** 🏆 | **83.00 ± 0%** 🏆 |

## First 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 189.2µ ±  1% | **4.969Ki ± 0%** 🏆 | **92.00 ± 0%** 🏆 |
| gormT | **187.9µ ±  0%** 🏆 | 5.836Ki ± 0% | 100.0 ± 0% |
| ldb | 205.1µ ±  0% | 10.36Ki ± 0% | 188.0 ± 0% |
| xorm | 381.2µ ±  0% | 6.103Ki ± 0% | 152.0 ± 0% |

## SelectNuller 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | **2.347m ±  2%** 🏆 | **320.6Ki ± 0%** 🏆 | **13.80k ± 0%** 🏆 |
| gormT | 2.439m ±  1% | 551.8Ki ± 0% | 13.82k ± 0% |
| ldb | 2.965m ±  1% | 1.110Mi ± 0% | 25.91k ± 0% |
| xorm | 3.440m ±  2% | 1.530Mi ± 6% | 35.84k ± 0% |

## Update 操作性能比较

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| gorm | 717.4µ ±  1% | 8.221Ki ± 0% | 93.00 ± 0% |
| gormT | 692.8µ ±  7% | 7.604Ki ± 0% | **75.00 ± 0%** 🏆 |
| ldb | **398.1µ ±  4%** 🏆 | 6.741Ki ± 0% | 105.0 ± 0% |
| xorm | 571.0µ ±  7% | **4.406Ki ± 0%** 🏆 | 116.0 ± 0% |

> 🏆 表示该指标的最佳性能（最小值）
<!-- BENCHMARK_RESULTS_END -->

