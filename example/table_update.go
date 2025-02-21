package main

import (
	"encoding/json"
	"example/dbinit"
	"fmt"
	"github.com/lontten/lcore/types"
	"github.com/lontten/ldb"
	return_type "github.com/lontten/ldb/return-type"
)

func TableUpdate() {
	var user = User{
		Name: types.NewString("abc"),
	}
	num, err := ldb.Update(dbinit.DB, ldb.W(), &user, ldb.E().
		SetNull("abc").
		TableName("t_user").
		ReturnType(return_type.PrimaryKey).
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
