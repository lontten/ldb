package benchmark

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lontten/ldb/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB ldb.Engine
var GDB *gorm.DB

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
	}
	DB = db

	gdb, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		fmt.Println("init gdb error:", err.Error())
	}
	GDB = gdb
}
