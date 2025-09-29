package main

import (
	"example/dbinit"
	"fmt"

	"github.com/lontten/ldb/v2"
)

type ExamUser struct {
	Id   int64
	Name string
	Age  int
}

func QueryOne() {
	var user ExamUser
	num, err := ldb.NativeQueryScan(dbinit.DB, `
SELECT 
  CASE 
    WHEN id > 10 THEN NULL  -- 当id大于10时返回null
    WHEN id < 10 THEN id    -- 当id小于10时返回具体id值
  END AS id,
  NULL AS name ,
  age
FROM t_test2
limit 1
`).ScanOne(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	fmt.Println(user.Id)
	fmt.Println(user.Name)
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
