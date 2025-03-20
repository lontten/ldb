package ldb

import (
	"database/sql"
)

type MysqlVersion int

const (
	MysqlVersionLast MysqlVersion = iota
	MysqlVersion5
	MysqlVersion8
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
	ctx.dialectNeedLastInsertId = true
	if c.Version == MysqlVersionLast {
		c.Version = MysqlVersion8
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
