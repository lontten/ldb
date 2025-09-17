package ldb

import (
	"fmt"

	"github.com/lontten/ldb/v2/sqltype"
	"github.com/lontten/ldb/v2/utils"
)

func Update(db Engine, dest any, wb *WhereBuilder, extra ...*ExtraContext) (int64, error) {
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

	ctx.wb.And(wb)

	dialect.tableUpdateGen()
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

func UpdateByPrimaryKey(db Engine, dest any, extra ...*ExtraContext) (int64, error) {
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

	wb := W()
	for _, name := range ctx.primaryKeyColumnNames {
		find := utils.Find(ctx.columns, name)
		if find == -1 {
			return 0, fmt.Errorf("primaryKey column %s not set value", name)
		} else {
			wb.fieldValue(name, ctx.columnValues[find])
			ctx.columnValues = append(ctx.columnValues[:find], ctx.columnValues[find+1:]...)
			ctx.columns = append(ctx.columns[:find], ctx.columns[find+1:]...)
		}
	}

	ctx.wb.And(wb)

	dialect.tableUpdateGen()
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
