# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/go-test/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->

## 最新基准测试结果

测试时间: 2025-09-19 03:33:29 UTC

```
goos: linux
goarch: amd64
pkg: github.com/lontten/ldb/v2/benchmark
cpu: AMD EPYC 7763 64-Core Processor                
               │ benchmark_results.txt │
               │        sec/op         │
Insert_mysql-4            412.8µ ± ∞ ¹
¹ need >= 6 samples for confidence interval at level 0.95

               │ benchmark_results.txt │
               │         B/op          │
Insert_mysql-4           4.672Ki ± ∞ ¹
¹ need >= 6 samples for confidence interval at level 0.95

               │ benchmark_results.txt │
               │       allocs/op       │
Insert_mysql-4             84.00 ± ∞ ¹
¹ need >= 6 samples for confidence interval at level 0.95
```

<!-- BENCHMARK_RESULTS_END -->

