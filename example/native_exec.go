package main

import (
	"example/dbinit"
	"fmt"
	"github.com/lontten/ldb"
)

func Exec() {
	num, err := ldb.Exec(dbinit.DB, "delete from t_ka where id=?", 1)
	fmt.Println(num, err)
}
