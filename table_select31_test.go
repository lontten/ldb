package ldb

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type User31 struct {
	Id int64
}

func (User31) TableConf() *TableConfContext {
	return TableConf("t_user").
		PrimaryKeys("id").
		AutoColumn("id")
}
func TestFirst31_mysql(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT id FROM t_user WHERE id IN (?) ORDER BY id DESC LIMIT 1;")).
		WithArgs(11).
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(11),
		)

	user, err := First[User21](engine, W().PrimaryKey(11), E().ShowSql())
	as.Nil(err)
	as.Equal(int64(11), user.Id, "id error")
}

func TestFirst31_pg(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT id FROM t_user WHERE id IN ($1) ORDER BY id DESC LIMIT 1;")).
		WithArgs(11).
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(11),
		)

	user, err := First[User21](engine, W().PrimaryKey(11), E().ShowSql())
	as.Nil(err)
	as.Equal(int64(11), user.Id, "id error")
}
