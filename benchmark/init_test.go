package benchmark

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/lontten/ldb/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"xorm.io/xorm"
)

var DB ldb.Engine
var GDB *gorm.DB
var XDB *xorm.Engine

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
		fmt.Println("init ldb error:", err.Error())
	}
	DB = db

	gdb, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		fmt.Println("init gorm error:", err.Error())
	}
	GDB = gdb

	engine, err := xorm.NewEngine("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("init xorm error:", err.Error())
	}
	XDB = engine
}
