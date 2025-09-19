# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-19 08:51:40 UTC
测试环境: Linux (AMD EPYC 7763 64-Core Processor)

> 说明：数值越低性能越好，±表示95%置信区间波动范围

# Go 基准测试报告

## 环境信息

| 参数 | 值 |
|------|----|
| goos | linux |
| goarch | amd64 |
| pkg | github.com/lontten/ldb/v2/benchmark |
| cpu | AMD EPYC 7763 64-Core Processor |

## Insert_ldb

| 测试名称 | sec/op | B/op | allocs/op |
|----------|--------|------|-----------|
| Insert_ldb-4 | 304.7µ ± | 4.673Ki ± 0% | 84.00 ± 0% |

## Insert_gorm

| 测试名称 | sec/op | B/op | allocs/op |
|----------|--------|------|-----------|
| Insert_gorm-4 | 306.6µ ± | 4.673Ki ± 0% | 84.00 ± 0% |

## Select_ldb

| 测试名称 | sec/op | B/op | allocs/op |
|----------|--------|------|-----------|
| Select_ldb-4 | 592.5µ ± 31% | 49.21Ki ± 0% | 1.121k ± 0% |

## Select_gorm

| 测试名称 | sec/op | B/op | allocs/op |
|----------|--------|------|-----------|
| Select_gorm-4 | 353.9µ ± 15% | 49.21Ki ± 0% | 1.121k ± 0% |


<!-- BENCHMARK_RESULTS_END -->

