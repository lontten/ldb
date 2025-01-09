package main

import (
	"example/dbinit"
	"fmt"
	"github.com/lontten/ldb"
	"github.com/lontten/ldb/types"
)

func First() {
	var u = User{
		Id:   nil,
		Name: types.NewString("xxx"),
	}
	fmt.Println(u)
	var m = make(map[string]any)
	m["a"] = 1
	m["b"] = "bb"
	m["c"] = nil
	eq := ldb.W().Eq("abc", "xxx")

	num, err := ldb.First[User](dbinit.DB, ldb.W().Or(eq).
		Model(u), ldb.E().ShowSql().SkipSoftDelete())
	fmt.Println(num)
	fmt.Println(err)
}

func First2() {
	one, err := ldb.First[User](dbinit.DB, ldb.W().Eq("name", "fjakdsf").
		IsNotNull("name"), ldb.E().ShowSql())
	fmt.Println(one)
	fmt.Println(err)
}

func List() {
	num, err := ldb.List[User](dbinit.DB, ldb.W().Eq("id", 1), ldb.E().ShowSql())
	fmt.Println(num)
	fmt.Println(err)
}

func GetOrInsert() {
	var u = User{
		Name: types.NewString("kb"),
		Age:  types.NewInt(33),
	}
	d, err := ldb.GetOrInsert[User](dbinit.DB, ldb.W().Eq("name", "kb"), &u, ldb.E().ShowSql())
	Log(d)
	Log(u)
	fmt.Println(err)
}

func InsertOrHas() {
	var u = User{
		Name: types.NewString("kc"),
		Age:  types.NewInt(33),
	}
	has, err := ldb.InsertOrHas(dbinit.DB, ldb.W().Eq("name", "kc"), &u, ldb.E().ShowSql())
	fmt.Println(has)
	Log(u)
	fmt.Println(err)

}
