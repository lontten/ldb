package benchmark

import (
	"fmt"

	"github.com/lontten/ldb/v2"
)

var DB ldb.Engine

func init() {
	conf := ldb.PgConf{
		Host:     "localhost",
		Port:     "5432",
		DbName:   "benchmark_test",
		User:     "postgres",
		Password: "postgres",
	}
	db, err := ldb.Connect(conf, nil)
	if err != nil {
		fmt.Println("init db error:", err.Error())
		panic(err)
	}
	DB = db
}
