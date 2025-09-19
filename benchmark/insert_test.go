package benchmark

import (
	"fmt"
	"testing"

	"github.com/lontten/ldb/v2"
)

func BenchmarkInsert_mysql(b *testing.B) {
	b.ResetTimer()

	// 执行b.N次（基准测试核心循环）
	for i := 0; i < b.N; i++ {
		u := User{
			Id:    0,
			Name:  "tom",
			Email: fmt.Sprintf("xx%d@xx.com", i),
		}

		_, err := ldb.Insert(DB, u, ldb.E().ShowSql())
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

}
