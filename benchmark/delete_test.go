package benchmark

import (
	"context"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func setupTest_BenchmarkDelete(b *testing.B) {
	InsertUsers(50000, b)

	b.Cleanup(func() {
		CleanUsers(b)
	})

}
func BenchmarkDelete_ldb(b *testing.B) {
	setupTest_BenchmarkDelete(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := ldb.Delete[User](DB, ldb.W().Eq("id", i))
		if err != nil {
			b.Fatalf("delete failed: %v", err)
		}
	}

}

func BenchmarkDelete_gorm(b *testing.B) {
	setupTest_BenchmarkDelete(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := GDB.Delete(&User{}, i).Error
		if err != nil {
			b.Fatalf("delete failed: %v", err)
		}
	}

}

func BenchmarkDelete_gormT(b *testing.B) {
	ctx := context.Background()
	setupTest_BenchmarkDelete(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gorm.G[User](GDB).Where("id = ?", i).Delete(ctx)
		if err != nil {
			b.Fatalf("delete failed: %v", err)
		}
	}

}

func BenchmarkDelete_xorm(b *testing.B) {
	setupTest_BenchmarkDelete(b)

	user := new(User)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := XDB.ID(i).Delete(user)
		if err != nil {
			b.Fatalf("delete failed: %v", err)
		}
	}

}
