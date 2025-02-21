package main

import (
	"example/dbinit"
	"fmt"
	"github.com/lontten/lcore/types"
	"github.com/lontten/ldb"
)

func Del() {
	var u = User{
		Id:   nil,
		Name: types.NewString("xxx"),
	}
	fmt.Println(u)
	var m = make(map[string]any)
	m["a"] = 1
	m["b"] = "bb"
	m["c"] = nil

	num, err := ldb.Delete[User](dbinit.DB, ldb.W().PrimaryKey(1), ldb.E().ShowSql())
	fmt.Println(num)
	fmt.Println(err)
}
