package ldb

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestInsert8_mysql(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO t_user (name, age) VALUES (?, ?) AS new ON DUPLICATE KEY UPDATE name = new.name, age = new.age, id = LAST_INSERT_ID(id);")).
		WithArgs("tom", 22).
		WillReturnResult(sqlmock.NewResult(10, 1))

	var u = User{
		Id:   0,
		Name: "tom",
		Age:  22,
	}
	num, err := Insert(engine, &u, E().ShowSql().WhenDuplicateKey().DoUpdate())
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
	as.Equal(int64(10), u.Id, "id error")
}

func TestInsert8_mysql2(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO t_user (name) VALUES (?) AS new ON DUPLICATE KEY UPDATE name = new.name, id = LAST_INSERT_ID(id);")).
		WithArgs("tom").
		WillReturnResult(sqlmock.NewResult(10, 1))

	var u = User{
		Id:   0,
		Name: "tom",
	}
	num, err := Insert(engine, &u, E().ShowSql().WhenDuplicateKey().DoUpdate())
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
	as.Equal(int64(10), u.Id, "id error")
}

func TestInsert8_pg(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO t_user (name) VALUES ($1) ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name RETURNING id;")).
		WithArgs("tom").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).
			AddRow(10),
		)

	var u = User{
		Id:   0,
		Name: "tom",
	}
	num, err := Insert(engine, &u, E().ShowSql().WhenDuplicateKey().DoUpdate())
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
	as.Equal(int64(10), u.Id, "id error")
}
