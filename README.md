# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-19 05:44:39 UTC
测试环境: Linux (AMD EPYC 7763 64-Core Processor)

> 说明：数值越低性能越好，±表示95%置信区间波动范围

```
goos: linux
goarch: amd64
pkg: github.com/lontten/ldb/v2/benchmark
cpu: AMD EPYC 7763 64-Core Processor                
               │ benchmark_results.txt │
               │        sec/op         │
Insert_mysql-4             471.5µ ± 5%

               │ benchmark_results.txt │
               │         B/op          │
Insert_mysql-4            4.673Ki ± 0%

               │ benchmark_results.txt │
               │       allocs/op       │
Insert_mysql-4              84.00 ± 0%
```

测试时间: 2025-09-19 05:40:04 UTC
测试环境: Linux (AMD EPYC 7763 64-Core Processor)

> 说明：数值越低性能越好，±表示95%置信区间波动范围

```
goos: linux
goarch: amd64
pkg: github.com/lontten/ldb/v2/benchmark
cpu: AMD EPYC 7763 64-Core Processor                
               │ benchmark_results.txt │
               │        sec/op         │
Insert_mysql-4             350.1µ ± 1%

               │ benchmark_results.txt │
               │         B/op          │
Insert_mysql-4            4.673Ki ± 0%

               │ benchmark_results.txt │
               │       allocs/op       │
Insert_mysql-4              84.00 ± 0%
```

测试时间: 2025-09-19 05:36:47 UTC

```
goos: linux
goarch: amd64
pkg: github.com/lontten/ldb/v2/benchmark
cpu: AMD EPYC 7763 64-Core Processor                
               │ benchmark_results.txt │
               │        sec/op         │
Insert_mysql-4             412.5µ ± 8%

               │ benchmark_results.txt │
               │         B/op          │
Insert_mysql-4            4.673Ki ± 0%

               │ benchmark_results.txt │
               │       allocs/op       │
Insert_mysql-4              84.00 ± 0%
```
<!-- BENCHMARK_RESULTS_END -->

