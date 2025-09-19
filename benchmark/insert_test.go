package benchmark

import (
	"context"
	"fmt"
	"testing"

	"github.com/lontten/ldb/v2"
	"gorm.io/gorm"
)

func BenchmarkInsert_ldb(b *testing.B) {
	_, err := ldb.Exec(DB, "DELETE FROM users WHERE 1=1")
	if err != nil {
		b.Fatalf("初始化清理数据失败: %v", err)
	}

	b.Cleanup(func() {
		_, err := ldb.Exec(DB, "DELETE FROM users WHERE 1=1")
		if err != nil {
			b.Logf("测试后清理数据失败: %v", err)
		}
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u := User{
			Id:    0,
			Name:  "tom",
			Email: fmt.Sprintf("xx%d@xx.com", i),
		}

		_, err = ldb.Insert(DB, u)
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

}

func BenchmarkInsert_gorm(b *testing.B) {
	_, err := ldb.Exec(DB, "DELETE FROM users WHERE 1=1")
	if err != nil {
		b.Fatalf("初始化清理数据失败: %v", err)
	}

	b.Cleanup(func() {
		_, err = ldb.Exec(DB, "DELETE FROM users WHERE 1=1")
		if err != nil {
			b.Logf("测试后清理数据失败: %v", err)
		}
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u := User{
			Id:    0,
			Name:  "tom",
			Email: fmt.Sprintf("xx%d@xx.com", i),
		}
		err = GDB.Create(u).Error
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

}

func BenchmarkInsert_gormt(b *testing.B) {
	ctx := context.Background()
	_, err := ldb.Exec(DB, "DELETE FROM users WHERE 1=1")
	if err != nil {
		b.Fatalf("初始化清理数据失败: %v", err)
	}

	b.Cleanup(func() {
		_, err = ldb.Exec(DB, "DELETE FROM users WHERE 1=1")
		if err != nil {
			b.Logf("测试后清理数据失败: %v", err)
		}
	})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u := User{
			Id:    0,
			Name:  "tom",
			Email: fmt.Sprintf("xx%d@xx.com", i),
		}
		err = gorm.G[User](GDB).Create(ctx, &u)
		if err != nil {
			b.Errorf("insert failed: %v", err)
		}
	}

}
