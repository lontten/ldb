package main

import (
	"encoding/json"
	"example/dbinit"
	"fmt"

	"github.com/lontten/ldb/v2"
)

func Build1() {
	var u []User
	num, dto, err := ldb.QueryBuild(dbinit.DB).ShowSql().
		Select("*").
		From("t_user u").
		Native(`
`).
		Where("id > 1").
		Page(1, 10).
		ScanPage(&u)
	fmt.Println(num, dto, err)
	bytes, err := json.Marshal(dto)
	fmt.Println(string(bytes))

	for _, user := range u {
		bytes2, _ := json.Marshal(user)
		fmt.Println(string(bytes2))
	}

}

func Build2() {
	neq := ldb.W().Neq("id", 7)
	var u User
	num, err := ldb.QueryBuild(dbinit.DB).
		Select("id").Select("name").
		From("t_user").
		Where("id = 2").
		WhereBuilder(ldb.W().
			Eq("id", 3).
			Eq("id", 5).
			Or(neq)).
		Limit(2).ShowSql().
		ScanOne(&u)
	fmt.Println(num, err)
	bytes, err := json.Marshal(u)
	fmt.Println(string(bytes))
}
