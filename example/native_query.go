package main

import (
	"example/dbinit"
	"fmt"

	"github.com/lontten/ldb/v2"
)

func QueryOne() {
	var ka Ka
	num, err := ldb.NativeQueryScan(dbinit.DB, "select * from t_ka where id=?", 2).ScanOne(&ka)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	fmt.Println(*ka.Id)
	fmt.Println(*ka.Name)
}
func QueryOne1() {
	var n int
	num, err := ldb.NativeQueryScan(dbinit.DB, "select 1").ScanOne(&n)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	fmt.Println(n)
}

func QueryList() {
	var list []User
	num, err := ldb.NativeQueryScan(dbinit.DB, "select * from t_user where id>1 limit 1").ScanList(&list)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	for _, ka := range list {
		fmt.Println(*ka.Id)
		fmt.Println(*ka.Name)
	}
}

func QueryList2() {
	var list []User
	num, err := ldb.NativeQueryScan(dbinit.DB, "select * from t_user where id>1").ScanList(&list)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)

	for _, ka := range list {
		fmt.Println(*ka.Id)
		fmt.Println(*ka.Name)
	}
}
