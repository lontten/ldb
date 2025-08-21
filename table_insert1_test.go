package ldb

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lontten/lcore/types"
	"github.com/stretchr/testify/assert"
)

type User struct {
	Id       int64
	Name     string
	Name2    string
	Age      int
	Age2     int
	Birthday types.Date
}

func (User) TableConf() *TableConf {
	return new(TableConf).
		Table("t_user").
		PrimaryKeys("id").
		AutoPrimaryKey("id")
}

func TestInsert_mysql(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO t_user (name) VALUES (?);")).
		WithArgs("tom").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   0,
		Name: "tom",
	}
	num, err := Insert(engine, u, E().ShowSql())
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}

func TestInsert_pg(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO t_user (name) VALUES ($1);")).
		WithArgs("tom").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   0,
		Name: "tom",
	}
	num, err := Insert(engine, u, E().ShowSql())
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}
