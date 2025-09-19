package benchmark

import (
	"os"

	"github.com/lontten/ldb/v2"
)

var DB ldb.Engine

func init() {
	conf := ldb.PgConf{
		Host:     os.Getenv("localhost"),
		Port:     os.Getenv("5432"),
		DbName:   os.Getenv("benchmark_test"),
		User:     os.Getenv("postgres"),
		Password: os.Getenv("postgres"),
	}
	db, err := ldb.Connect(conf, nil)
	if err != nil {
		panic(err)
	}
	DB = db
}
