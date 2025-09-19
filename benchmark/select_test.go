package benchmark

import (
	"fmt"
	"testing"

	"github.com/lontten/ldb/v2"
)

func BenchmarkSelect_ldb(b *testing.B) {
	for i := 0; i < 100; i++ {
		u := User{
			Id:    0,
			Name:  "tom",
			Email: fmt.Sprintf("xx%d@xx.com", i),
		}

		_, err := ldb.Insert(DB, u)
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

	b.Cleanup(func() {
		_, err := ldb.Exec(DB, "DELETE FROM users WHERE 1=1")
		if err != nil {
			b.Logf("测试后清理数据失败: %v", err)
		}
	})

	b.ResetTimer()

	// 执行b.N次（基准测试核心循环）
	for i := 0; i < b.N; i++ {
		_, err := ldb.List[User](DB, ldb.W().Eq("1", 1))
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

}

func BenchmarkSelect_gorm(b *testing.B) {
	for i := 0; i < 100; i++ {
		u := User{
			Id:    0,
			Name:  "tom",
			Email: fmt.Sprintf("xx%d@xx.com", i),
		}

		_, err := ldb.Insert(DB, u)
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

	b.Cleanup(func() {
		_, err := ldb.Exec(DB, "DELETE FROM users WHERE 1=1")
		if err != nil {
			b.Logf("测试后清理数据失败: %v", err)
		}
	})

	b.ResetTimer()

	// 执行b.N次（基准测试核心循环）
	for i := 0; i < b.N; i++ {
		_, err := ldb.List[User](DB, ldb.W().Eq("1", 1))
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

}
