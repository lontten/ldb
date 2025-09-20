package benchmark

import (
	"context"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func BenchmarkUpdate_setupTest(b *testing.B) {
	for i := 0; i < 10000; i++ {
		u := User{
			Id:    int64(i),
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
func BenchmarkUpdate_ldb(b *testing.B) {
	BenchmarkUpdate_setupTest(b)

	b.ResetTimer()

	upm := User{
		Email: "aa@aa.com",
	}

	for i := 0; i < b.N; i++ {
		_, err := ldb.Update[User](DB, upm, ldb.W().Eq("id", i))
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkUpdate_gorm(b *testing.B) {
	BenchmarkUpdate_setupTest(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := GDB.Model(&User{}).Where("id = ?", i).Update("email", "aa@aa.com").Error
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkUpdate_gormT(b *testing.B) {
	ctx := context.Background()
	BenchmarkUpdate_setupTest(b)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := gorm.G[User](GDB).Where("id = ?", i).Update(ctx, "email", "aa@aa.com")
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}
