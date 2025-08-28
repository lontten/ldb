package ldb

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/lontten/lcore/types"
	"github.com/lontten/ldb/sqltype"
	"github.com/lontten/ldb/utils"
	"github.com/pkg/errors"
)

//------------------------------------Select--------------------------------------------

// First 根据条件获取第一个
func First[T any](db Engine, wb *WhereBuilder, extra ...*ExtraContext) (t *T, err error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...) // 表名，set，select配置
	ctx.limit = types.NewInt64(1)

	dest := new(T)
	ctx.initScanDestOneT(dest)
	if ctx.err != nil {
		return nil, ctx.err
	}

	ctx.initConf() //初始化表名，主键，自增id
	if ctx.lastSql == "" {
		if ctx.autoPrimaryKeyColumnName != nil {
			ctx.lastSql = " ORDER BY " + *ctx.autoPrimaryKeyColumnName + " DESC"
		}
	}

	ctx.initColumns()
	ctx.initColumnsValueSoftDel()

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableSelectGen()
	if ctx.hasErr() {
		return nil, ctx.err
	}
	sqlStr := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
		utils.PrintSql(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(sqlStr, ctx.args...)
	if err != nil {
		return nil, err
	}
	num, err := ctx.ScanLnT(rows)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return nil, nil
	}
	return dest, nil
}

func List[T any](db Engine, wb *WhereBuilder, extra ...*ExtraContext) (list []T, err error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...) // 表名，set，select配置

	var dest = &[]T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, false)
	if ctx.err != nil {
		return nil, ctx.err
	}

	ctx.initConf() //初始化表名，主键，自增id
	ctx.initColumns()

	ctx.initColumnsValueSoftDel()

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableSelectGen()
	if ctx.hasErr() {
		return nil, ctx.err
	}
	sqlStr := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(sqlStr, ctx.args...)
	if err != nil {
		return nil, err
	}
	num, err := ctx.ScanT(rows)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return make([]T, 0), nil
	}
	return *dest, nil
}

func ListP[T any](db Engine, wb *WhereBuilder, extra ...*ExtraContext) (list []*T, err error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)

	var dest = &[]*T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, false)
	if ctx.err != nil {
		return nil, ctx.err
	}

	ctx.initConf() //初始化表名，主键，自增id
	ctx.initColumns()
	if len(ctx.modelSelectFieldNames) == 0 {
		ctx.modelSelectFieldNames = ctx.modelAllFieldNames
	}
	ctx.initColumnsValueSoftDel()

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableSelectGen()
	if ctx.hasErr() {
		return nil, ctx.err
	}
	sqlStr := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(sqlStr, ctx.args...)
	if err != nil {
		return nil, err
	}
	num, err := ctx.ScanT(rows)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return make([]*T, 0), nil
	}
	return *dest, nil
}

// Has
func Has[T any](db Engine, wb *WhereBuilder, extra ...*ExtraContext) (t bool, err error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)
	ctx.modelSelectFieldNames = []string{"1"}

	dest := new(T)
	ctx.initScanDestOneT(dest)
	if ctx.err != nil {
		return false, ctx.err
	}

	ctx.initConf() //初始化表名，主键，自增id
	ctx.initColumnsValueSoftDel()

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableSelectGen()
	if ctx.hasErr() {
		return false, ctx.err
	}
	sqlStr := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return false, nil
	}
	rows, err := db.query(sqlStr, ctx.args...)
	if err != nil {
		return false, err
	}
	defer func(rows *sql.Rows) {
		utils.PanicErr(rows.Close())
	}(rows)
	return rows.Next(), nil
}

// Count
func Count[T any](db Engine, wb *WhereBuilder, extra ...*ExtraContext) (t int64, err error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)
	ctx.modelSelectFieldNames = []string{"count(*)"}

	dest := new(T)
	ctx.initScanDestOneT(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}

	ctx.initConf() //初始化表名，主键，自增id
	ctx.initColumnsValueSoftDel()

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableSelectGen()
	if ctx.hasErr() {
		return 0, ctx.err
	}
	sqlStr := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return 0, nil
	}
	var total int64
	rows, err := db.query(sqlStr, ctx.args...)
	if err != nil {
		return 0, err
	}
	defer func(rows *sql.Rows) {
		utils.PanicErr(rows.Close())
	}(rows)
	for rows.Next() {
		box := reflect.ValueOf(&total).Interface()
		err = rows.Scan(box)
		if err != nil {
			return total, err
		}
		return total, nil
	}
	return total, errors.New("rows no data")
}

// GetOrInsert
// d insert 的 对象，
// e 通用设置，select 自定义字段
func GetOrInsert[T any](db Engine, wb *WhereBuilder, d *T, extra ...*ExtraContext) (*T, error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...) // 表名，set，select配置
	ctx.sqlType = sqltype.Select

	dest := new(T)
	ctx.initScanDestOneT(dest)
	if ctx.err != nil {
		return nil, ctx.err
	}

	ctx.initConf() //初始化表名，主键，自增id

	ctx.initColumns()
	ctx.initColumnsValueSoftDel()

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableSelectGen()
	if ctx.hasErr() {
		return nil, ctx.err
	}
	sqlStr := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(sqlStr, ctx.args...)
	if err != nil {
		return nil, err
	}
	num, err := ctx.ScanLnT(rows)
	if err != nil {
		return nil, err
	}
	if num != 0 {
		return dest, nil
	}

	//------------

	ctx.query.Reset()
	ctx.args = []any{}
	ctx.sqlType = sqltype.Insert

	ctx.initModelDest(d) //初始化参数

	ctx.initColumnsValue() //初始化cv
	ctx.initColumnsValueExtra()
	ctx.initColumnsValueSoftDel() // 软删除

	dialect.tableInsertGen()
	if ctx.hasErr() {
		return nil, ctx.err
	}

	sqlStr = dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}

	if ctx.returnAutoPrimaryKey == pkQueryReturn {
		rows, err = db.query(sqlStr, ctx.args...)
		if err != nil {
			return nil, err
		}
		num, err = ctx.ScanLnT(rows)
		if err != nil {
			return nil, err
		}
	}

	exec, err := db.exec(sqlStr, ctx.args...)
	if err != nil {
		return nil, err
	}
	if ctx.returnAutoPrimaryKey == pkFetchReturn {
		id, err := exec.LastInsertId()
		if err != nil {
			return nil, err
		}
		if id > 0 {
			ctx.setLastInsertId(id)
			if ctx.hasErr() {
				return nil, ctx.err
			}
		}
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errors.New("insert affected 0")
	}
	return d, nil
}

// InsertOrHas 根据条件查询是否已存在，不存在则直接插入
// 应用场景：例如添加 后台管理员 时，如果名字已存在，返回名字重复，否者正常添加。
// d insert 的 对象，
// e 通用设置，select 自定义字段
func InsertOrHas(db Engine, wb *WhereBuilder, d any, extra ...*ExtraContext) (bool, error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...) // 表名，set，select配置
	ctx.modelSelectFieldNames = []string{"1"}
	ctx.sqlType = sqltype.Select

	ctx.initModelDest(d) //初始化参数
	ctx.initConf()       //初始化表名，主键，自增id

	ctx.initColumnsValue() //初始化cv
	ctx.initColumnsValueExtra()
	ctx.initColumnsValueSoftDel() // 软删除
	if ctx.err != nil {
		return false, ctx.err
	}

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableSelectGen()
	if ctx.hasErr() {
		return false, ctx.err
	}
	sqlStr := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return false, nil
	}
	rows, err := db.query(sqlStr, ctx.args...)
	if err != nil {
		return false, err
	}
	defer func(rows *sql.Rows) {
		utils.PanicErr(rows.Close())
	}(rows)
	if rows.Next() {
		return true, nil
	}

	//------------

	ctx.query.Reset()
	ctx.args = []any{}
	ctx.sqlType = sqltype.Insert
	ctx.initColumnsValueSoftDel() // 软删除

	dialect.tableInsertGen()
	if ctx.hasErr() {
		return false, ctx.err
	}

	sqlStr = dialect.getSql()
	if ctx.showSql {
		fmt.Println(sqlStr, ctx.args)
	}
	if ctx.noRun {
		return false, nil
	}

	if ctx.returnAutoPrimaryKey == pkQueryReturn {
		rows, err = db.query(sqlStr, ctx.args...)
		if err != nil {
			return false, err
		}
		_, err = ctx.ScanLnT(rows)
		if err != nil {
			return false, err
		}
	}

	exec, err := db.exec(sqlStr, ctx.args...)
	if err != nil {
		return false, err
	}
	if ctx.returnAutoPrimaryKey == pkFetchReturn {
		id, err := exec.LastInsertId()
		if err != nil {
			return false, err
		}
		if id > 0 {
			ctx.setLastInsertId(id)
			if ctx.hasErr() {
				return false, ctx.err
			}
		}
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		return false, err
	}
	if affected == 0 {
		return false, errors.New("insert affected 0")
	}
	return false, nil
}
