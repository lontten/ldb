package main

import (
	"encoding/json"
	"example/dbinit"
	"fmt"

	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/ldb/v2"
	return_type "github.com/lontten/ldb/v2/return-type"
)

func TableUpdate() {
	var user = User{
		Name: types.NewString("abc"),
	}
	num, err := ldb.Update(dbinit.DB, &user, ldb.W(), ldb.E().
		SetNull("abc").
		TableName("t_user").
		ReturnType(return_type.Auto).
		WhenDuplicateKey("name").DoUpdate().
		ShowSql().NoRun(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	bytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
