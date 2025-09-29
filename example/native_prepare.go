package main

import (
	"example/dbinit"
	"fmt"

	"github.com/lontten/ldb/v2"
)

func Prepare() {
	stmt, err := ldb.Prepare(dbinit.DB, "delete from t_ka where id= $1 ")
	if err != nil {
		panic(err)
	}
	num, err := stmt.Exec(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
}

func Prepare2() {
	stmt, err := ldb.Prepare(dbinit.DB, "select id from t_ka where id=$1 ")
	if err != nil {
		panic(err)
	}
	var n int64
	num, err := stmt.QueryScan(2).ScanOne(&n)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
	fmt.Println(num)
}

func Prepare3() {
	stmt, err := ldb.Prepare(dbinit.DB, "select id from t_ka where id>$1 ")
	if err != nil {
		panic(err)
	}
	var list []int64
	num, err := stmt.QueryScan(2).ScanList(&list)
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
	fmt.Println(num)
}
