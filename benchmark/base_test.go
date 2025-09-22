package benchmark

import (
	"testing"

	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/ldb/v2"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	CreatedAt types.LocalDateTime
}

func (User) TableConf() *ldb.TableConf {
	return new(ldb.TableConf).
		Table("users").
		PrimaryKeys("id").
		AutoPrimaryKey("id")
}

func (User) TableName() string {
	return "users"
}

func InsertUsers(count int, b *testing.B) {
	if count <= 0 {
		return // 无需插入
	}

	const batchSize = 1000 // 固定每批插入 1000 条
	totalBatches := (count + batchSize - 1) / batchSize

	for batch := 0; batch < totalBatches; batch++ {
		start := batch * batchSize
		end := start + batchSize
		if end > count {
			end = count
		}

		users := make([]User, 0, end-start)
		for i := start; i < end; i++ {
			users = append(users, User{
				Id:    int64(i + 1), // ID从1开始
				Name:  "tom",
				Email: "xx@xx.com",
			})
		}

		if err := GDB.Create(&users).Error; err != nil {
			b.Fatalf("batch insert failed (batch %d): %v", batch, err)
		}
	}
}
func CleanUsers(b *testing.B) {
	_, err := ldb.Exec(DB, "TRUNCATE TABLE users")
	if err != nil {
		b.Fatalf("清理数据失败: %v", err)
	}
}
