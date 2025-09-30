package benchmark

import (
	"context"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func setupTest_BenchmarkUpdate(b *testing.B) {
	InsertUsers(50000, b)

	b.Cleanup(func() {
		CleanUsers(b)
	})

}
func BenchmarkUpdate_ldb(b *testing.B) {
	setupTest_BenchmarkUpdate(b)

	upm := User{
		Age: 14,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := ldb.Update(DB, upm, ldb.W().Eq("id", i))
		if err != nil {
			b.Fatalf("update failed: %v", err)
		}
	}

}

func BenchmarkUpdate_gorm(b *testing.B) {
	setupTest_BenchmarkUpdate(b)

	upm := User{
		Age: 14,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		upm.Id = int64(i)
		err := GDB.Save(&upm).Error
		if err != nil {
			b.Fatalf("update failed: %v", err)
		}
	}

}

func BenchmarkUpdate_gormT(b *testing.B) {
	ctx := context.Background()
	setupTest_BenchmarkUpdate(b)

	upm := User{
		Age: 14,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gorm.G[User](GDB).Where("id = ?", i).Updates(ctx, upm)
		if err != nil {
			b.Fatalf("update failed: %v", err)
		}
	}

}

func BenchmarkUpdate_xorm(b *testing.B) {
	setupTest_BenchmarkUpdate(b)

	upm := User{
		Age: 14,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := XDB.ID(i).Update(upm)
		if err != nil {
			b.Fatalf("update failed: %v", err)
		}
	}

}
