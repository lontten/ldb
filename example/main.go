package main

import (
	"encoding/json"
	"fmt"

	"github.com/lontten/ldb/v2"
)

func Log(v any) {
	bytes, _ := json.Marshal(v)
	fmt.Println(string(bytes))
}

type Kaaa struct {
}
type User struct {
	Id   *int    `db:"user_id"`
	Name *string `db:"-"`
	Age  *int    `db:"-"`
	Kaaa `db:"-"`
}

func (u User) TableConf() *ldb.TableConf {
	return new(ldb.TableConf).Table("xjwy_user").
		PrimaryKeys("user_id").
		AutoPrimaryKey("user_id")
}

func main() {
	QueryOne()
}
