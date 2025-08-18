package ldb

import (
	"database/sql"
)

type MysqlVersion int

const (
	MysqlVersionLast MysqlVersion = iota
	MysqlVersion5    MysqlVersion = iota

	MysqlVersion8_0_19
	MysqlVersion8_0_20
	MysqlVersion8Last
)

type MysqlConf struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	Other    string
	Version  MysqlVersion
}

func (c MysqlConf) dialect(ctx *ormContext) Dialecter {
	ctx.insertCanReturn = false
	if c.Version == MysqlVersionLast {
		c.Version = MysqlVersion8Last
	}
	return &MysqlDialect{
		ctx:       ctx,
		dbVersion: c.Version,
	}
}

func (c MysqlConf) open() (*sql.DB, error) {
	dsn := c.User + ":" + c.Password +
		"@tcp(" + c.Host +
		":" + c.Port +
		")/" + c.DbName + "?"

	if c.Other == "" {
		dsn += "charset=utf8mb4&parseTime=True&loc=Local"
	} else {
		dsn += c.Other
	}
	return sql.Open("mysql", dsn)
}
