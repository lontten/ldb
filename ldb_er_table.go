package ldb

import (
	"fmt"
	"github.com/lontten/lcore/types"
	"github.com/lontten/ldb/sqltype"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

// ------------------------------------Insert--------------------------------------------

// Insert 插入或者根据主键冲突更新
func Insert(db Engine, v any, extra ...*ExtraContext) (num int64, err error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)
	ctx.sqlType = sqltype.Insert
	ctx.sqlIsQuery = true

	ctx.initModelDest(v) //初始化参数
	ctx.initConf()       //初始化表名，主键，自增id

	ctx.initColumnsValue() //初始化cv
	ctx.initColumnsValueExtra()
	ctx.initColumnsValueSoftDel() // 软删除

	dialect.tableInsertGen()
	if ctx.hasErr() {
		return 0, ctx.err
	}

	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return 0, nil
	}

	if ctx.sqlIsQuery {
		rows, err := db.query(sql, ctx.args...)
		if err != nil {
			return 0, err
		}
		return ctx.ScanLnT(rows)
	}

	exec, err := db.exec(sql, ctx.args...)
	if err != nil {
		return 0, err
	}
	if ctx.needLastInsertId {
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

//------------------------------------Delete--------------------------------------------

func Delete[T any](db Engine, wb *WhereBuilder, extra ...*ExtraContext) (int64, error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)
	ctx.sqlType = sqltype.Delete
	ctx.sqlIsQuery = false

	dest := new(T)
	ctx.initScanDestOneT(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}
	ctx.initConf() //初始化表名，主键，自增id

	ctx.initColumnsValueSoftDel()

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableDelGen()
	if ctx.hasErr() {
		return 0, ctx.err
	}
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return 0, nil
	}
	exec, err := db.exec(sql, ctx.args...)
	if err != nil {
		return 0, err
	}
	return exec.RowsAffected()
}

//------------------------------------Update--------------------------------------------

func Update(db Engine, wb *WhereBuilder, dest any, extra ...*ExtraContext) (int64, error) {
	db = db.init()
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	ctx.initExtra(extra...)
	ctx.sqlType = sqltype.Update
	ctx.sqlIsQuery = false

	ctx.initModelDest(dest)
	ctx.initConf() //初始化表名，主键，自增id

	ctx.initColumnsValue() //初始化cv
	ctx.initColumnsValueExtra()
	ctx.initColumnsValueSoftDel() // 软删除

	ctx.initPrimaryKeyByWhere(wb) // byId(1,2,...)
	if ctx.err != nil {
		return 0, ctx.err
	}
	ctx.wb.And(wb)

	dialect.tableUpdateGen()
	if ctx.hasErr() {
		return 0, ctx.err
	}
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return 0, nil
	}
	exec, err := db.exec(sql, ctx.args...)
	if err != nil {
		return 0, err
	}
	return exec.RowsAffected()
}

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
		ctx.lastSql = " ORDER BY " + strings.Join(ctx.primaryKeyNames, ",")
	}

	ctx.initColumns()
	ctx.initColumnsValueSoftDel()

	ctx.initPrimaryKeyByWhere(wb)
	ctx.wb.And(wb)

	dialect.tableSelectGen()
	if ctx.hasErr() {
		return nil, ctx.err
	}
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(sql, ctx.args...)
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
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(sql, ctx.args...)
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
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(sql, ctx.args...)
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
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return false, nil
	}
	rows, err := db.query(sql, ctx.args...)
	if err != nil {
		return false, err
	}
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
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return 0, nil
	}
	var total int64
	rows, err := db.query(sql, ctx.args...)
	if err != nil {
		return 0, err
	}
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
	ctx.sqlIsQuery = true

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
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(sql, ctx.args...)
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
	ctx.sqlIsQuery = true

	ctx.initModelDest(d) //初始化参数

	ctx.initColumnsValue() //初始化cv
	ctx.initColumnsValueExtra()
	ctx.initColumnsValueSoftDel() // 软删除

	dialect.tableInsertGen()
	if ctx.hasErr() {
		return nil, ctx.err
	}

	sql = dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return nil, nil
	}

	if ctx.sqlIsQuery {
		rows, err = db.query(sql, ctx.args...)
		if err != nil {
			return nil, err
		}
		num, err = ctx.ScanLnT(rows)
		if err != nil {
			return nil, err
		}
	}

	exec, err := db.exec(sql, ctx.args...)
	if err != nil {
		return nil, err
	}
	if ctx.needLastInsertId {
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
	ctx.sqlType = sqltype.Select
	ctx.sqlIsQuery = true

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
	sql := dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return false, nil
	}
	rows, err := db.query(sql, ctx.args...)
	if err != nil {
		return false, err
	}
	if rows.Next() {
		return true, nil
	}

	//------------

	ctx.query.Reset()
	ctx.args = []any{}
	ctx.sqlType = sqltype.Insert
	ctx.sqlIsQuery = true
	ctx.initColumnsValueSoftDel() // 软删除

	dialect.tableInsertGen()
	if ctx.hasErr() {
		return false, ctx.err
	}

	sql = dialect.getSql()
	if ctx.showSql {
		fmt.Println(sql, ctx.args)
	}
	if ctx.noRun {
		return false, nil
	}

	if ctx.sqlIsQuery {
		rows, err = db.query(sql, ctx.args...)
		if err != nil {
			return false, err
		}
		_, err = ctx.ScanLnT(rows)
		if err != nil {
			return false, err
		}
	}

	exec, err := db.exec(sql, ctx.args...)
	if err != nil {
		return false, err
	}
	if ctx.needLastInsertId {
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
