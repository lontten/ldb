package ldb

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

// 插入时，唯一索引冲突，处理策略-更新

func TestInsert5_mysql(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO t_user (id, name) VALUES (?, ?) AS new ON DUPLICATE KEY UPDATE id = new.id, name = new.name;")).
		WithArgs(1, "tom").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   1,
		Name: "tom",
	}
	num, err := Insert(engine, u, E().ShowSql().WhenDuplicateKey().DoUpdate())
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}

func TestInsert5_mysql2(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO t_user (id, name, birthday, age) VALUES (?, ?, NOW(), NULL) AS new ON DUPLICATE KEY UPDATE name = ?;")).
		WithArgs(1, "tom", "n").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   1,
		Name: "tom",
	}
	num, err := Insert(engine, u, E().
		ShowSql().
		SetNow("birthday").SetNull("age").
		WhenDuplicateKey().DoUpdate(Set().Set("name", "n")),
	)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}

func TestInsert5_mysql3(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &MysqlConf{})

	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO t_user (id, name, birthday, age) VALUES (?, ?, NOW(), NULL) AS new ON DUPLICATE KEY UPDATE name = ?, name2 = ?;")).
		WithArgs(1, "tom", "n", "n2").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   1,
		Name: "tom",
	}
	num, err := Insert(engine, u, E().
		ShowSql().
		SetNow("birthday").SetNull("age").
		WhenDuplicateKey().DoUpdate(Set().Set("name", "n").Set("name2", "n2")),
	)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}

func TestInsert5_pg(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO t_user (id, name) VALUES ($1, $2) ON CONFLICT () DO UPDATE SET id = EXCLUDED.id, name = EXCLUDED.name;")).
		WithArgs(1, "tom").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   1,
		Name: "tom",
	}
	num, err := Insert(engine, u, E().ShowSql().WhenDuplicateKey().DoUpdate())
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}

func TestInsert5_pg2(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO t_user (id, name, birthday, age) VALUES ($1, $2, NOW(), NULL) ON CONFLICT () DO UPDATE SET name = $3;")).
		WithArgs(1, "tom", "n").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   1,
		Name: "tom",
	}
	num, err := Insert(engine, u, E().
		ShowSql().
		SetNow("birthday").SetNull("age").
		WhenDuplicateKey().DoUpdate(Set().Set("name", "n")),
	)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}

func TestInsert5_pg3(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectExec(regexp.QuoteMeta(
		"INSERT INTO t_user (id, name, birthday, age) VALUES ($1, $2, NOW(), NULL) ON CONFLICT () DO UPDATE SET name = $3, name2 = $4;")).
		WithArgs(1, "tom", "n", "n2").
		WillReturnError(nil).
		WillReturnResult(sqlmock.NewResult(0, 1))

	var u = User{
		Id:   1,
		Name: "tom",
	}
	num, err := Insert(engine, u, E().
		ShowSql().
		SetNow("birthday").SetNull("age").
		WhenDuplicateKey().DoUpdate(Set().Set("name", "n").Set("name2", "n2")),
	)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
}
