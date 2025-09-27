package benchmark

import (
	"context"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func setupTest_BenchmarkInsert(b *testing.B) {
	CleanUsers(b)

	b.Cleanup(func() {
		CleanUsers(b)
	})
}
func BenchmarkInsert_ldb(b *testing.B) {
	setupTest_BenchmarkInsert(b)

	user := User{
		Id:   0,
		Name: "tom",
		Age:  12,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		user.Id = 0
		_, err := ldb.Insert(DB, &user)
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkInsert_gorm(b *testing.B) {
	setupTest_BenchmarkInsert(b)

	user := User{
		Id:   0,
		Name: "tom",
		Age:  12,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		user.Id = 0
		err := GDB.Create(&user).Error
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkInsert_gormT(b *testing.B) {
	ctx := context.Background()
	setupTest_BenchmarkInsert(b)

	user := User{
		Id:   0,
		Name: "tom",
		Age:  12,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		user.Id = 0
		err := gorm.G[User](GDB).Create(ctx, &user)
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkInsert_xorm(b *testing.B) {
	setupTest_BenchmarkInsert(b)

	user := User{
		Id:   0,
		Name: "tom",
		Age:  12,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		user.Id = 0
		_, err := XDB.Insert(user)
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}
