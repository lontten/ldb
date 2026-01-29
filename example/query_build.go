package main

import (
	"example/dbinit"
	"fmt"

	"github.com/lontten/ldb/v2"
	"github.com/lontten/lutil/logutil"
)

func QueryBuild() {

	list, err := ldb.QueryBuild[User](dbinit.DB).ShowSql().
		Select("u.*").
		From("t_user u").
		Convert(ldb.ConvertRegister("age", func(v *int) any {
			fmt.Println(v, v == nil)
			if v == nil {
				return "kk"
			}
			if *v == 1 {
				return "one"
			}
			return "abc"
		})).
		List()
	if err != nil {
		panic(err)
	}

	for _, v := range list {
		logutil.Log(v)
	}
}
