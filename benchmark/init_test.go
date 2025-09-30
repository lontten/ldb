package benchmark

import (
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/lontten/ldb/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, err := ldb.Connect(conf, &ldb.PoolConf{
		MaxIdleCount: 10,               // 设置最大空闲连接数
		MaxOpen:      20,               // 设置最大打开连接数
		MaxLifetime:  30 * time.Minute, // 设置连接的最大生存期
		MaxIdleTime:  10 * time.Minute, // 设置连接的最大空闲时间
	})
	if err != nil {
		fmt.Println("init ldb error:", err.Error())
		return
	}
	DB = db

	gdb, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("init gorm error:", err.Error())
		return
	}
	// 获取底层sql.DB对象以设置连接池
	sqlDB, err := gdb.DB()
	if err != nil {
		fmt.Println("init gorm error:", err.Error())
		return
	}
	// 设置连接池参数
	// 设置最大打开连接数
	sqlDB.SetMaxOpenConns(20)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置连接的最大生存期
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	// 设置连接的最大空闲时间
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	GDB = gdb

	engine, err := xorm.NewEngine("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("init xorm error:", err.Error())
		return
	}
	// 设置最大打开连接数 (0表示无限制)
	engine.SetMaxOpenConns(20)
	// 设置最大空闲连接数
	engine.SetMaxIdleConns(10)
	// 设置连接的最大生存期
	engine.SetConnMaxLifetime(30 * time.Minute)
	// 设置连接的最大空闲时间
	engine.SetConnMaxIdleTime(10 * time.Minute)
	XDB = engine
}
