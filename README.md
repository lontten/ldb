# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-19 09:39:48 UTC

> 说明：数值越低性能越好，±表示95%置信区间波动范围

# Go 基准测试报告

## 环境信息

| 参数 | 值 |
|------|----|
| goos | linux |
| goarch | amd64 |
| pkg | github.com/lontten/ldb/v2/benchmark |
| cpu | AMD EPYC 7763 64-Core Processor |

## InsertOrUpdate

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| ldb | 316.1µ ± | 4.673Ki ± 0% | 84.00 ± 0% |
| gorm | 315.5µ ± | 4.673Ki ± 0% | 84.00 ± 0% |

## Insert

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| ldb | 313.9µ ± | 4.673Ki ± 0% | 84.00 ± 0% |
| gorm | 634.8µ ± | 6.964Ki ± 0% | 94.00 ± 0% |

## Select

| 实现 | sec/op | B/op | allocs/op |
|------|--------|------|-----------|
| ldb | 419.5µ ± 13% | 49.21Ki ± 0% | 1.121k ± 0% |


<!-- BENCHMARK_RESULTS_END -->

