package dbinit

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lontten/ldb"
)

//import _ "github.com/jackc/pgx/v5/stdlib"

var DB ldb.Engine

func init() {
	conf := ldb.MysqlConf{
		Host:     "127.0.0.1",
		Port:     "3306",
		DbName:   "test",
		User:     "root",
		Password: "123456",
	}
	db, err := ldb.Connect(conf, nil)
	if err != nil {
		panic(err)
	}
	DB = db
}

func init2() {
	conf := ldb.PgConf{
		Host:     "127.0.0.1",
		Port:     "5432",
		DbName:   "test",
		User:     "postgres",
		Password: "123456",
	}
	db, err := ldb.Connect(conf, nil)
	if err != nil {
		panic(err)
	}
	DB = db
}
