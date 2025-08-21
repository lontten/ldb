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

func TestWhereBuilder9(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().PrimaryKey(1, 2)

	sql, args, err := w1.toSql(engine.getDialect().parse)
	as.ErrorIs(err, ErrNoPk)
	as.Equal("", sql)
	as.Nil(args)
}

func TestWhereBuilder10(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().PrimaryKey(1, 2)

	sql, args, err := w1.toSql(engine.getDialect().parse, "id")
	as.Nil(err)
	as.Equal("id IN (?,?)", sql)
	as.Equal([]any{1, 2}, args)
}

func TestWhereBuilder11(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().PrimaryKey(1)

	sql, args, err := w1.toSql(engine.getDialect().parse, "id")
	as.Nil(err)
	as.Equal("id IN (?)", sql)
	as.Equal([]any{1}, args)
}

func TestWhereBuilder12(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().PrimaryKey(1, struct {
	}{})

	sql, args, err := w1.toSql(engine.getDialect().parse, "id")
	as.ErrorIs(err, ErrTypePkArgs)
	as.Equal("", sql)
	as.Nil(args)
}

func TestWhereBuilder13(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().PrimaryKey(1, 2)

	sql, args, err := w1.toSql(engine.getDialect().parse, "id", "name")
	as.ErrorIs(err, ErrNeedMultiPk)
	as.Equal("", sql)
	as.Nil(args)
}

func TestWhereBuilder14(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().PrimaryKey(struct {
		Id   int
		Name string
	}{
		Id:   1,
		Name: "name",
	})

	sql, args, err := w1.toSql(engine.getDialect().parse, "id", "name")
	as.Nil(err)
	as.Equal("id = ? AND name = ?", sql)
	as.Equal([]any{1, "name"}, args)
}

func TestWhereBuilder15(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	w1 := W().PrimaryKey(struct {
		DocId   int
		DocName string
	}{
		DocId:   1,
		DocName: "name",
	})

	sql, args, err := w1.toSql(engine.getDialect().parse, "doc_id", "doc_name")
	as.Nil(err)
	as.Equal("doc_id = ? AND doc_name = ?", sql)
	as.Equal([]any{1, "name"}, args)
}

func TestWhereBuilder16(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	var m = make(map[string]any)
	m["doc_id"] = 1
	m["doc_name"] = "name"

	w1 := W().PrimaryKey(m)

	sql, args, err := w1.toSql(engine.getDialect().parse, "doc_id", "doc_name")
	as.Nil(err)
	as.Equal("doc_id = ? AND doc_name = ?", sql)
	as.Equal([]any{1, "name"}, args)
}

func TestWhereBuilder17(t *testing.T) {
	as := assert.New(t)
	db, _, err := sqlmock.New()
	as.Nil(err, fmt.Sprintf("failed to open sqlmock database: %s", err))
	defer db.Close()
	engine := MustConnectMock(db, &PgConf{})

	var m = make(map[string]any)
	m["doc_id"] = 1
	m["doc_name"] = "name"

	var m2 = make(map[string]any)
	m2["doc_id"] = 2
	m2["doc_name"] = "name2"

	w1 := W().PrimaryKey(m, m2)

	sql, args, err := w1.toSql(engine.getDialect().parse, "doc_id", "doc_name")
	as.Nil(err)
	as.Equal("(doc_id = ? AND doc_name = ?) OR (doc_id = ? AND doc_name = ?)", sql)
	as.Equal([]any{1, "name", 2, "name2"}, args)
}
