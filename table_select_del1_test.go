package ldb

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lontten/ldb/v2/softdelete"
	"github.com/stretchr/testify/assert"
)

type UserDel1 struct {
	Id int64
	softdelete.DeleteTimeNil
}

func (UserDel1) TableConf() *TableConfContext {
	return TableConf("t_user").
		PrimaryKeys("id").
		AutoColumn("id")
}
func TestFirst_del1_mysql(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT id ,deleted_at FROM t_user WHERE deleted_at IS NULL AND id = ? ORDER BY id DESC LIMIT 1;")).
		WithArgs(11).
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(11),
		)

	user, err := First[UserDel1](engine, W().Eq("id", 11), E().ShowSql())
	as.Nil(err)
	as.Equal(int64(11), user.Id, "id error")
}

func TestFirst_del1_pg(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectQuery(regexp.QuoteMeta(
		"SELECT id ,deleted_at FROM t_user WHERE deleted_at IS NULL AND id = $1 ORDER BY id DESC LIMIT 1;")).
		WithArgs(11).
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(11),
		)

	user, err := First[UserDel1](engine, W().Eq("id", 11), E().ShowSql())
	as.Nil(err)
	as.Equal(int64(11), user.Id, "id error")
}
