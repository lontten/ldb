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
	initMysql()
}
func initMysql() {
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

func initPg() {
	conf := ldb.PgConf{
		Host:     os.Getenv("LDB_PG_HOST"),
		Port:     os.Getenv("LDB_PG_PORT"),
		DbName:   os.Getenv("LDB_PG_DB"),
		User:     os.Getenv("LDB_PG_USER"),
		Password: os.Getenv("LDB_PG_PWD"),
	}
	db, err := ldb.Connect(conf, nil)
	if err != nil {
		panic(err)
	}
	DB = db
}
