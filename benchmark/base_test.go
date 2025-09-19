package benchmark

import (
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
