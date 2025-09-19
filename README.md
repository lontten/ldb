# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-19 08:33:35 UTC
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

## sec/op

| 测试名称 | 值 |
|----------|----|
| Insert-4 | 340.7µ ± 1% |
| Select-4 | 576.2µ ± 3% |

## B/op

| 测试名称 | 值 |
|----------|----|
| Insert-4 | 4.673Ki ± 0% |
| Select-4 | 49.21Ki ± 0% |

## allocs/op

| 测试名称 | 值 |
|----------|----|
| Insert-4 | 84.00 ± 0% |
| Select-4 | 1.121k ± 0% |


<!-- BENCHMARK_RESULTS_END -->

