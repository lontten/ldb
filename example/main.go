package main

import (
	"github.com/lontten/ldb/v2"
)

type User struct {
	Id   *int `db:"id"`
	Name *string
	Age  string
}

func (u User) TableConf() *ldb.TableConfContext {
	return new(ldb.TableConfContext).Table("t_user").
		PrimaryKeys("id").
		AutoColumn("id")
}

func main() {
	QueryBuild()
}
