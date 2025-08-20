package ldb

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestWhereBuilder1(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().Eq("a", 1)

	sql, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a = ?", sql)
	as.Equal([]any{1}, args)
}

func TestWhereBuilder2(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().Eq("a", 1).Eq("b", 2)

	sql, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a = ? AND b = ?", sql)
	as.Equal([]any{1, 2}, args)
}

func TestWhereBuilder3(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().Eq("a", 1)
	w2 := W().Eq("b", 2)

	w0 := w1.And(w2)

	sql, args, err := w0.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a = ? AND b = ?", sql)
	as.Equal([]any{1, 2}, args)
}

func TestWhereBuilder4(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().Eq("a", 1)
	w2 := W().Eq("b", 2)

	w0 := w1.Or(w2)

	sql, args, err := w0.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("(b = ?) OR (a = ?)", sql)
	as.Equal([]any{2, 1}, args)
}

func TestWhereBuilder5(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w01 := W().Eq("a1", "a1").Or(W().Eq("a2", "a2"))
	w02 := W().Eq("b1", "b1").Or(W().Eq("b2", "b2"))

	w0 := W().And(w01).And(w02)

	sql, args, err := w0.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("((a2 = ?) OR (a1 = ?)) AND ((b2 = ?) OR (b1 = ?))", sql)
	as.Equal([]any{"a2", "a1", "b2", "b1"}, args)
}

func TestWhereBuilder6(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w01 := W().Eq("a1", "a1").Or(W().Eq("a2", "a2"))

	w0 := W().Or(w01).Or(w01)

	w00 := W().Or(w0)

	sql, args, err := w00.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("((a2 = ?) OR (a1 = ?)) OR ((a2 = ?) OR (a1 = ?))", sql)
	as.Equal([]any{"a2", "a1", "a2", "a1"}, args)
}

func TestWhereBuilder7(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().Eq("a1", "a1")
	w2 := W().Eq("a2", "a2")
	w3 := W().Eq("a3", "a3")

	w0 := W().And(w1).And(w2).And(w3)

	sql, args, err := w0.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a1 = ? AND a2 = ? AND a3 = ?", sql)
	as.Equal([]any{"a1", "a2", "a3"}, args)
}

func TestWhereBuilder8(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().Or(W())
	w2 := W().And(W())
	w3 := W().Or(W().Or(W()))

	w0 := W().And(w1).And(w2).And(w3)

	sql, args, err := w0.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("", sql)
	as.Nil(args)
}
