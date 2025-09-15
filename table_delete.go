package ldb

import (
	"github.com/lontten/ldb/v2/sqltype"
)

func Delete[T any](db Engine, wb *WhereBuilder, extra ...*ExtraContext) (int64, error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)
	ctx.sqlType = sqltype.Delete

	dest := new(T)
	ctx.initScanDestOneT(dest)
	ctx.initConf() //初始化表名，主键，自增id

	ctx.initColumnsValueSoftDel()

	ctx.wb.And(wb)

	dialect.tableDelGen()
	if ctx.hasErr() {
		return 0, ctx.err
	}

	dialect.getSql()
	dialectSql := ctx.dialectSql
	ctx.printSql()
	if ctx.noRun {
		return 0, nil
	}

	exec, err := db.exec(dialectSql, ctx.args...)
	if err != nil {
		return 0, err
	}
	return exec.RowsAffected()
}
