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
	//QueryOneT2()
	//QueryListT2()
	//
	//QueryOneT()
	//QueryList2()
	//
	//Prepare4()
	//time.Sleep(1 * time.Hour)
	//TableInsert()
	//Build1()
	//Build2()
	//Del()
	//First()
	//First2()
	//QueryOneT()
	//List()
	//Insertttt()
	//TableUpdate()
	//GetOrInsert()
	//InsertOrHas()
	//TableInsert3()
	//QueryBuild()
	//Has()
	//TableUpdate2()
}
