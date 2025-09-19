# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-19 06:59:38 UTC
测试环境: Linux (AMD EPYC 7763 64-Core Processor)

> 说明：数值越低性能越好，±表示95%置信区间波动范围

| goos: | linux |
| --- | --- |
| goarch: | amd64 |
| --- | --- |
| pkg: | github.com/lontten/ldb/v2/benchmark |
| --- | --- |
| cpu: | AMD EPYC 7763 64-Core Processor |
| --- | --- |
| Insert-4 | 394.7µ | ± | 2% |
| Select-4 | 506.5µ | ± | 0% |
| geomean | 447.1µ | | |
| Insert-4 | 4.673Ki | ± | 0% |
| Select-4 | 49.21Ki | ± | 0% |
| geomean | 15.16Ki | | |
| Insert-4 | 84.00 | ± | 0% |
| Select-4 | 1.121k | ± | 0% |
| geomean | 306.9 | | |
<!-- BENCHMARK_RESULTS_END -->

