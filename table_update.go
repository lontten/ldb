package ldb

import (
	"github.com/lontten/ldb/sqltype"
	"github.com/lontten/ldb/utils"
)

func Update(db Engine, wb *WhereBuilder, dest any, extra ...*ExtraContext) (int64, error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)
	ctx.sqlType = sqltype.Update

	ctx.initModelDest(dest)
	ctx.initConf() //初始化表名，主键，自增id

	ctx.initColumnsValue() //初始化cv
	ctx.initColumnsValueExtra()

	ctx.initColumnsValueSoftDel() // 软删除

	if ctx.err != nil {
		return 0, ctx.err
	}
	ctx.wb.And(wb)

	dialect.tableUpdateGen()
	if ctx.hasErr() {
		return 0, ctx.err
	}
	sqlStr := dialect.getSql()
	if ctx.showSql {
		utils.PrintSql(sqlStr, ctx.args...)
	}
	if ctx.noRun {
		return 0, nil
	}
	exec, err := db.exec(sqlStr, ctx.args...)
	if err != nil {
		return 0, err
	}
	return exec.RowsAffected()
}
