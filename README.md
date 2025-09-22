# ldb

[![Go Test](https://github.com/lontten/ldb/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/lontten/ldb/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/lontten/ldb/branch/ctt/graph/badge.svg)](https://codecov.io/gh/lontten/ldb)

<!-- BENCHMARK_RESULTS_START -->
## 最新基准测试结果

测试时间: 2025-09-25 03:52:19 UTC

> 说明：数值越低性能越好，±表示波动范围

# Go 基准测试报告



## 环境信息



| 参数 | 值 |

|------|----|

| goos | linux |

| goarch | amd64 |

| pkg | github.com/lontten/ldb/v2/benchmark |

| cpu | AMD EPYC 7763 64-Core Processor |



## Delete 操作性能比较



| 实现 | sec/op | B/op | allocs/op |

|------|--------|------|-----------|

| gorm | 607.9µ ±  1% | 5.423Ki ± 0% | 62.00 ± 0% |

| gormT | 608.2µ ±  1% | 6.290Ki ± 0% | 70.00 ± 0% |

| ldb | **288.5µ ±  1%** 🏆 | **2.952Ki ± 0%** 🏆 | **61.00 ± 0%** 🏆 |

| xorm | 362.5µ ±  1% | 4.040Ki ± 0% | 106.0 ± 0% |



## Insert 操作性能比较



| 实现 | sec/op | B/op | allocs/op |

|------|--------|------|-----------|

| gorm | 621.1µ ±  2% | 5.936Ki ± 0% | 81.00 ± 0% |

| gormT | 618.4µ ±  3% | 6.834Ki ± 0% | 89.00 ± 0% |

| ldb | **310.3µ ±  2%** 🏆 | 6.148Ki ± 0% | 113.0 ± 0% |

| xorm | 488.5µ ± 10% | **3.359Ki ± 0%** 🏆 | **71.00 ± 0%** 🏆 |



## SelectFirst 操作性能比较



| 实现 | sec/op | B/op | allocs/op |

|------|--------|------|-----------|

| gorm | **210.5µ ± 17%** 🏆 | **4.496Ki ± 0%** 🏆 | **80.00 ± 0%** 🏆 |

| gormT | 232.2µ ± 10% | 5.371Ki ± 0% | 88.00 ± 0% |

| ldb | 224.2µ ±  8% | 5.858Ki ± 0% | 122.0 ± 0% |

| xorm | 418.5µ ±  5% | 4.711Ki ± 0% | 122.0 ± 0% |



## Select 操作性能比较



| 实现 | sec/op | B/op | allocs/op |

|------|--------|------|-----------|

| gorm | **1.727m ±  4%** 🏆 | **253.5Ki ± 0%** 🏆 | 12.79k ± 0% |

| gormT | 1.742m ± 22% | 399.7Ki ± 0% | 12.81k ± 0% |

| ldb | 1.732m ±  5% | 394.2Ki ± 0% | **9.869k ± 0%** 🏆 |

| xorm | 2.903m ± 25% | 1.031Mi ± 4% | 28.83k ± 0% |



## Update 操作性能比较



| 实现 | sec/op | B/op | allocs/op |

|------|--------|------|-----------|

| gorm | 783.7µ ±  8% | 6.001Ki ± 0% | **66.00 ± 0%** 🏆 |

| gormT | 689.6µ ± 11% | 7.009Ki ± 0% | 76.00 ± 0% |

| ldb | **384.3µ ±  9%** 🏆 | 4.921Ki ± 0% | 92.00 ± 0% |

| xorm | 557.0µ ±  4% | **3.703Ki ± 0%** 🏆 | 100.0 ± 0% |



> 🏆 表示该指标的最佳性能（最小值）




<!-- BENCHMARK_RESULTS_END -->

