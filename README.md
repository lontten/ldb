# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-19 06:48:22 UTC
测试环境: Linux (AMD EPYC 7763 64-Core Processor)

> 说明：数值越低性能越好，±表示95%置信区间波动范围

| goos: | linux |
| --- | --- |
| goarch: | amd64 |
| pkg: | github.com/lontten/ldb/v2/benchmark |
| cpu: | AMD | EPYC | 7763 | 64-Core | Processor |
| │ | benchmark_results.txt | │ |
| │ | sec/op | │ |
| Insert-4 | 357.4µ | ± | 5% |
| Select-4 | 552.7µ | ± | 1% |
| geomean | 444.4µ |
| │ | benchmark_results.txt | │ |
| │ | B/op | │ |
| Insert-4 | 4.673Ki | ± | 0% |
| Select-4 | 49.21Ki | ± | 0% |
| geomean | 15.16Ki |
| │ | benchmark_results.txt | │ |
| │ | allocs/op | │ |
| Insert-4 | 84.00 | ± | 0% |
| Select-4 | 1.121k | ± | 0% |
| geomean | 306.9 |
## 最新

<!-- BENCHMARK_RESULTS_END -->

