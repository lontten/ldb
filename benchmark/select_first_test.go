package benchmark

import (
	"context"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func setupTest_BenchmarkFirst(b *testing.B) {
	InsertUsers(50000, b)

	b.Cleanup(func() {
		CleanUsers(b)
	})

}
func BenchmarkFirst_ldb(b *testing.B) {
	setupTest_BenchmarkFirst(b)

	b.ResetTimer()

	// 执行b.N次（基准测试核心循环）
	for i := 1; i < b.N; i++ {
		_, err := ldb.First[User](DB, ldb.W().Eq("id", i))
		if err != nil {
			b.Fatalf("select failed: %v", err)
		}
	}

}

func BenchmarkFirst_gorm(b *testing.B) {
	setupTest_BenchmarkFirst(b)

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		var user User
		err := GDB.First(&user, i).Error
		if err != nil {
			b.Fatalf("select failed: %v", err)
		}
	}

}

func BenchmarkFirst_gormT(b *testing.B) {
	ctx := context.Background()
	setupTest_BenchmarkFirst(b)

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		_, err := gorm.G[User](GDB).Where("id = ?", i).First(ctx)
		if err != nil {
			b.Fatalf("select failed: %v", err)
		}
	}

}

func BenchmarkFirst_xorm(b *testing.B) {
	setupTest_BenchmarkFirst(b)

	var user User

	b.ResetTimer()

	for i := 1; i < b.N; i++ {
		_, err := XDB.ID(i).Get(&user)
		if err != nil {
			b.Fatalf("select failed: %v", err)
		}
	}

}
