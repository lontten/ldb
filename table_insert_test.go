package ldb

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectExec("INSERT").
		WithArgs("tom").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   0,
		Name: "tom",
	}
	num, err := Insert(engine, u)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}
