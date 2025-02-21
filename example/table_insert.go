package main

import (
	"encoding/json"
	"errors"
	"example/dbinit"
	"fmt"
	"github.com/lontten/lcore/types"
	"github.com/lontten/ldb"
	return_type "github.com/lontten/ldb/return-type"
	"time"
)

func TableInsert() {
	var user = User{
		Name: types.NewString(time.Now().String()),
	}
	num, err := ldb.Insert(dbinit.DB, &user, ldb.E().
		SetNull("abc").
		TableName("t_user").
		ReturnType(return_type.PrimaryKey).
		WhenDuplicateKey("name").DoUpdate().
		ShowSql(),
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

func TableInsert2() {
	var user = User{
		Name: types.NewString("abc"),
	}
	num, err := ldb.Insert(dbinit.DB, &user, new(ldb.ExtraContext).
		TableName("t_user2").
		SetNull("uuid").
		WhenDuplicateKey().DoUpdate(ldb.Set().Set("user_state", 1).SetNull("name")).
		ShowSql(),
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

func TableInsert3() {
	tx, err := dbinit.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			tx.Rollback()
			return
		}
	}()
	var user = User{
		Name: types.NewString("99"),
	}
	num, err := ldb.Insert(tx, &user, ldb.E().
		ShowSql(),
	)
	err = errors.New("db tx")
	if err != nil {
		panic(err)
	}

	fmt.Println(num)
	Log(user)

	var user2 = User{
		Name: types.NewString("000"),
	}
	num, err = ldb.Insert(tx, &user2, ldb.E().
		ShowSql(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(num)
	Log(user2)
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
