package benchmark

import (
	"context"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func setupTest_BenchmarkSelect(b *testing.B) {
	InsertUsersZero(1000, b)

	b.Cleanup(func() {
		CleanUsers(b)
	})

}
func BenchmarkSelect_ldb(b *testing.B) {
	setupTest_BenchmarkSelect(b)

	b.ResetTimer()

	// 执行b.N次（基准测试核心循环）
	for i := 0; i < b.N; i++ {
		_, err := ldb.List[UserZero](DB, ldb.W().Eq("1", 1))
		if err != nil {
			b.Fatalf("select failed: %v", err)
		}
	}

}

func BenchmarkSelect_gorm(b *testing.B) {
	setupTest_BenchmarkSelect(b)

	var users []UserZero

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := GDB.Find(&users).Error
		if err != nil {
			b.Fatalf("select failed: %v", err)
		}
	}

}

func BenchmarkSelect_gormT(b *testing.B) {
	ctx := context.Background()
	setupTest_BenchmarkSelect(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gorm.G[UserZero](GDB).Find(ctx)
		if err != nil {
			b.Fatalf("select failed: %v", err)
		}
	}

}

func BenchmarkSelect_xorm(b *testing.B) {
	setupTest_BenchmarkSelect(b)

	var users []UserZero

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := XDB.Find(&users)
		if err != nil {
			b.Fatalf("select failed: %v", err)
		}
	}

}
