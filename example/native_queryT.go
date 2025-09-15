package main

import (
	"encoding/json"
	"example/dbinit"
	"fmt"
	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/ldb/v2"
)

func QueryOneT() {
	ka, err := ldb.QueryOne[User](dbinit.DB, "select * from t_user where id=?", 2222)
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(ka)
	fmt.Println(string(bytes))
}
func QueryOneT2() {
	ka, err := ldb.QueryOne[types.StringList](dbinit.DB, "select img_list from public.user  where id=$1", 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(*ka)
}

func QueryListT() {
	list, err := ldb.QueryList[Ka](dbinit.DB, "select * from t_ka where id>1")
	if err != nil {
		panic(err)
	}
	for _, ka := range list {
		fmt.Println(*ka.Id)
		fmt.Println(*ka.Name)
	}
}

func QueryListT2() {
	list, err := ldb.QueryListP[Ka](dbinit.DB, "select * from t_ka where id>1")
	if err != nil {
		panic(err)
	}
	for _, ka := range list {
		fmt.Println(*ka.Id)
		fmt.Println(*ka.Name)
	}
}
