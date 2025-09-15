package main

import (
	"example/dbinit"
	"fmt"

	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/ldb/v2"
)

func QueryBuild() {
	var list []User

	var age = types.NewInt(1)
	var name = "abc"
	_, _, err := ldb.QueryBuild(dbinit.DB).ShowSql().
		Select("u.*").
		From("t_user u").
		BoolWhere(age != nil, "u.age", age).
		BoolWhere(name != "", "u.name", name).
		Page(1, 100).
		ScanPage(&list)
	if err != nil {
		panic(err)
	}

	for _, ka := range list {
		fmt.Println(*ka.Id)
		fmt.Println(*ka.Name)
	}
}
