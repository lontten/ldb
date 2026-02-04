package ldb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestWhereBuilderNative1(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Native("o.name = ?", "a")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("o.name = ?", query)
	as.Equal([]any{"a"}, args)
}

func TestWhereBuilderNative2(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().
		Native("o.name = ?", "a").
		Native("u.age = ?", 1)

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("o.name = ? AND u.age = ?", query)
	as.Equal([]any{"a", 1}, args)
}

func TestWhereBuilderNativeIn1(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Native("o.name in (?)", "a")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("o.name in (?)", query)
	as.Equal([]any{"a"}, args)
}

func TestWhereBuilderNativeIn2(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Native("o.name in (?)", "a", "b")

	query, _, err := w1.toSql(engine.getDialect().parse)
	as.ErrorIs(err, ErrArgsLen)
	as.Equal("", query)
}

func TestWhereBuilderNativeIn3(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Native("o.name in (?)", []string{"a", "b"})

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("o.name in (?,?)", query)
	as.Equal([]any{"a", "b"}, args)
}

func TestWhereBuilderNativeIn4(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Native("o.name in (?) and u.age = ?", []string{"a", "b"}, 22)

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("o.name in (?,?) and u.age = ?", query)
	as.Equal([]any{"a", "b", 22}, args)
}

func TestWhereBuilderNativeIn5(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Native("o.name in (?) and u.age = ? or s.role in(?)", []string{"a", "b"}, 22, "c")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("o.name in (?,?) and u.age = ? or s.role in(?)", query)
	as.Equal([]any{"a", "b", 22, "c"}, args)
}

func TestWhereBuilderNativeIn6(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Native("o.name in (?) and u.age = ? or s.role in(?)", "c", 22, []string{"a", "b"})

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("o.name in (?) and u.age = ? or s.role in(?,?)", query)
	as.Equal([]any{"c", 22, "a", "b"}, args)
}
