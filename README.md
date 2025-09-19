# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-19 07:18:14 UTC
测试环境: Linux (AMD EPYC 7763 64-Core Processor)

> 说明：数值越低性能越好，±表示95%置信区间波动范围

### 测试环境信息
| goos: | linux |
| --- | --- |
| goarch: | amd64 |
| --- | --- |
| pkg: | github.com/lontten/ldb/v2/benchmark |
| --- | --- |
| cpu: | AMD EPYC 7763 64-Core Processor |
| --- | --- |

### 每次操作耗时
| 基准测试 | 数值 | ± | 波动范围 |
| -------- | ---- | - | -------- |
| Insert-4 | 323.0µ | ± | 11% |
| Select-4 | 553.6µ | ± | 0% |
| geomean | 422.9µ | | |
### 每次操作内存分配
| 基准测试 | 数值 | ± | 波动范围 |
| -------- | ---- | - | -------- |
| Insert-4 | 4.673Ki | ± | 0% |
| Select-4 | 49.21Ki | ± | 0% |
| geomean | 15.16Ki | | |
### 每次操作内存分配次数
| 基准测试 | 数值 | ± | 波动范围 |
| -------- | ---- | - | -------- |
| Insert-4 | 84.00 | ± | 0% |
| Select-4 | 1.121k | ± | 0% |
| geomean | 306.9 | | |
<!-- BENCHMARK_RESULTS_END -->

