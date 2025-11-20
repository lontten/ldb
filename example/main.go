package main

import (
	"github.com/lontten/ldb/v2"
)

type User struct {
	Id   *int `db:"id"`
	Name *string
	Age  *int
}

func (u User) TableConf() *ldb.TableConf {
	return new(ldb.TableConf).Table("t_user").
		PrimaryKeys("id").
		AutoPrimaryKey("id")
}

func main() {
	QueryBuild()
}
