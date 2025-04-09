package dbinit

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lontten/lcore/lcutils"
	"github.com/lontten/ldb"
	"os"
)

//import _ "github.com/jackc/pgx/v5/stdlib"

var DB ldb.Engine

func init() {
	conf := ldb.MysqlConf{
		Host:     os.Getenv("LDB_MYSQL_HOST"),
		Port:     os.Getenv("LDB_MYSQL_PORT"),
		DbName:   os.Getenv("LDB_MYSQL_DB"),
		User:     os.Getenv("LDB_MYSQL_USER"),
		Password: os.Getenv("LDB_MYSQL_PWD"),
		Version:  ldb.MysqlVersion5,
	}
	lcutils.LogJson(conf)
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
