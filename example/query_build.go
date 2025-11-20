package main

import (
	"example/dbinit"
	"fmt"

	"github.com/lontten/lcore/v2/logutil"
	"github.com/lontten/ldb/v2"
)

func QueryBuild() {
	var list []User

	_, err := ldb.QueryBuild(dbinit.DB).ShowSql().
		Select("u.*").
		From("t_user u").
		Convert(ldb.ConvertRegister("age", func(v *int) any {
			fmt.Println(v, v == nil)
			if v == nil {
				return 9
			}
			if *v == 1 {
				return 99
			}
			return 999
		})).
		ScanList(&list)
	if err != nil {
		panic(err)
	}

	for _, v := range list {
		logutil.Log(v)
	}
}
