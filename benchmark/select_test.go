package benchmark

import (
	"context"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func setupTest_BenchmarkSelect(b *testing.B) {
	for i := 0; i < 100; i++ {
		u := User{
			Id:    0,
			Name:  "tom",
			Email: "xx@xx.com",
		}

		_, err := ldb.Insert(DB, u)
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

	b.Cleanup(func() {
		_, err := ldb.Exec(DB, "TRUNCATE TABLE users")
		if err != nil {
			b.Fatalf("测试后清理数据失败: %v", err)
		}
	})

}
func BenchmarkSelect_ldb(b *testing.B) {
	setupTest_BenchmarkSelect(b)

	b.ResetTimer()

	// 执行b.N次（基准测试核心循环）
	for i := 0; i < b.N; i++ {
		_, err := ldb.List[User](DB, ldb.W().Eq("1", 1))
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkSelect_gorm(b *testing.B) {
	setupTest_BenchmarkSelect(b)

	var users []User

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := GDB.Find(&users).Error
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkSelect_gormT(b *testing.B) {
	ctx := context.Background()
	setupTest_BenchmarkSelect(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gorm.G[User](GDB).Find(ctx)
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}
