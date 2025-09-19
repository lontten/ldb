package benchmark

import (
	"testing"

	"github.com/lontten/ldb/v2"
)

func BenchmarkInsert_mysql(b *testing.B) {
	// 定义测试对象
	u := User{
		Id:   0,
		Name: "tom",
	}

	// 重置计时器（排除初始化时间）
	b.ResetTimer()

	// 执行b.N次（基准测试核心循环）
	for i := 0; i < b.N; i++ {
		_, err := ldb.Insert(DB, u, ldb.E().ShowSql())
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

}
