package benchmark

import (
	"context"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func setupTest_BenchmarkInsert(b *testing.B) {
	var cleanDataSql = "TRUNCATE TABLE users"
	_, err := ldb.Exec(DB, cleanDataSql)
	if err != nil {
		b.Fatalf("初始化清理数据失败: %v", err)
	}

	b.Cleanup(func() {
		_, err = ldb.Exec(DB, cleanDataSql)
		if err != nil {
			b.Fatalf("测试后清理数据失败: %v", err)
		}
	})
}
func BenchmarkInsert_ldb(b *testing.B) {
	setupTest_BenchmarkInsert(b)

	u := User{
		Id:    0,
		Name:  "tom",
		Email: "xx@xx.com",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.Id = 0
		_, err := ldb.Insert(DB, &u)
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkInsert_gorm(b *testing.B) {
	setupTest_BenchmarkInsert(b)

	u := User{
		Id:    0,
		Name:  "tom",
		Email: "xx@xx.com",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.Id = 0
		err := GDB.Create(&u).Error
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}

func BenchmarkInsert_gormT(b *testing.B) {
	ctx := context.Background()
	setupTest_BenchmarkInsert(b)

	u := User{
		Id:    0,
		Name:  "tom",
		Email: "xx@xx.com",
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.Id = 0
		err := gorm.G[User](GDB).Create(ctx, &u)
		if err != nil {
			b.Fatalf("insert failed: %v", err)
		}
	}

}
