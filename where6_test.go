package ldb

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/lontten/lcore/v2/types"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestWhereBuilderLike1(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Like(types.NewString("k"), "a")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a LIKE ?", query)
	as.Equal([]any{"%k%"}, args)
}

func TestWhereBuilderLike2(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().Like(types.NewString("k"), "a", "b", "c")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a LIKE ? OR b LIKE ? OR c LIKE ?", query)
	as.Equal([]any{"%k%", "%k%", "%k%"}, args)
}

func TestWhereBuilderLike3(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().LikeLeft(types.NewString("k"), "a")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a LIKE ?", query)
	as.Equal([]any{"k%"}, args)
}

func TestWhereBuilderLike4(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().LikeLeft(types.NewString("k"), "a", "b", "c")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a LIKE ? OR b LIKE ? OR c LIKE ?", query)
	as.Equal([]any{"k%", "k%", "k%"}, args)
}

func TestWhereBuilderLike21(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().NoLike(types.NewString("k"), "a")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a NOT LIKE ?", query)
	as.Equal([]any{"%k%"}, args)
}

func TestWhereBuilderLike22(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().NoLike(types.NewString("k"), "a", "b", "c")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a NOT LIKE ? AND b NOT LIKE ? AND c NOT LIKE ?", query)
	as.Equal([]any{"%k%", "%k%", "%k%"}, args)
}

func TestWhereBuilderLike23(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().NoLikeLeft(types.NewString("k"), "a")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a NOT LIKE ?", query)
	as.Equal([]any{"k%"}, args)
}

func TestWhereBuilderLike24(t *testing.T) {
	as := assert.New(t)
	engine := getMockDB(PgConf{})

	w1 := W().NoLikeLeft(types.NewString("k"), "a", "b", "c")

	query, args, err := w1.toSql(engine.getDialect().parse)
	as.Nil(err)
	as.Equal("a NOT LIKE ? AND b NOT LIKE ? AND c NOT LIKE ?", query)
	as.Equal([]any{"k%", "k%", "k%"}, args)
}
