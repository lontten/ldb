# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-19 10:59:52 UTC

> 说明：数值越低性能越好，±表示波动范围

# Go 基准测试报告

## 环境信息

| 参数 | 值 |
|------|----|
| goos | linux |
| goarch | amd64 |
| pkg | github.com/lontten/ldb/v2/benchmark |
| cpu | AMD EPYC 7763 64-Core Processor |

## Insert

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| ldb | 323.8µ ±  1% | 4.673Ki ± 0% | 84.00 ± 0% |
| gorm | 655.4µ ±  1% | 6.065Ki ± 0% | 86.00 ± 0% |
| gormT | 656.6µ ±  1% | 6.965Ki ± 0% | 94.00 ± 0% |

## Select

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| ldb | 428.8µ ± 12% | 49.21Ki ± 0% | 1.121k ± 0% |
| gorm | 1.729m ±  1% | 398.8Ki ± 0% | 13.05k ± 0% |
| gormT | 1.742m ±  1% | 399.7Ki ± 0% | 13.06k ± 0% |


<!-- BENCHMARK_RESULTS_END -->

