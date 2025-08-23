package ldb

import (
	"github.com/lontten/ldb/sqltype"
	"github.com/lontten/ldb/utils"
)

// Insert 插入或者根据主键冲突更新
func Insert(db Engine, v any, extra ...*ExtraContext) (num int64, err error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)
	ctx.sqlType = sqltype.Insert

	ctx.initModelDest(v) //初始化参数
	ctx.initConf()       //初始化表名，主键，自增id

	ctx.initColumnsValue() //初始化cv
	ctx.initTableNameExtra()
	ctx.initColumnsValueExtra()
	ctx.initColumnsValueSoftDel() // 软删除

	dialect.tableInsertGen()
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

	if ctx.returnAutoPrimaryKey == pkQueryReturn {
		rows, err := db.query(sqlStr, ctx.args...)
		if err != nil {
			return 0, err
		}
		return ctx.ScanLnT(rows)
	}

	exec, err := db.exec(sqlStr, ctx.args...)
	if err != nil {
		return 0, err
	}
	if ctx.returnAutoPrimaryKey == pkFetchReturn {
		id, err := exec.LastInsertId()
		if err != nil {
			return 0, err
		}
		if id > 0 {
			ctx.setLastInsertId(id)
			if ctx.hasErr() {
				return 0, ctx.err
			}
		}
	}
	return exec.RowsAffected()
}
