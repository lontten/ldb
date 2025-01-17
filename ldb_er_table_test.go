package ldb

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, "new sqlmock error")
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectExec("delete from user where id = ? ").
		WithArgs(1).
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   0,
		Name: "",
	}
	num, err := Insert(engine, u)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")

	mock.ExpectExec("update user set name = 'kk' where id = ? ").
		WithArgs(1).
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	num, err = Exec(engine, "update user set name = 'kk' where id = ? ", 1)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")

	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
}
func TestDel(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, "new sqlmock error")
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectExec("delete from user where id = ? ").
		WithArgs(1).
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	num, err := Exec(engine, "delete from user where id = ?", 1)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")

	mock.ExpectExec("update user set name = 'kk' where id = ? ").
		WithArgs(1).
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	num, err = Exec(engine, "update user set name = 'kk' where id = ? ", 1)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")

	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
}
