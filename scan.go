package ldb

import (
	"database/sql"
	"reflect"

	"github.com/lontten/ldb/v2/utils"
	"github.com/pkg/errors"
)

type rowColumnType struct {
	index            int    // 字段再row中的位置
	noNull           bool   // true 字段必定不为null
	databaseTypeName string // 字段-数据库数据类型
}

// ScanLn
// 接收一行结果
// 1.ptr single/comp
// 2.slice- single
func (ctx ormContext) ScanLn(rows *sql.Rows) (num int64, err error) {
	defer func(rows *sql.Rows) {
		utils.PanicErr(rows.Close())
	}(rows)

	num = 0
	t := ctx.destBaseType
	v := ctx.destBaseValue
	tP := ctx.scanDest

	columns, err := rows.Columns()
	if err != nil {
		return
	}

	cfm := make(map[string]compC)
	if ctx.destBaseTypeIsComp {
		cfm = getColIndex2FieldNameMap(columns, t)
	}

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return 0, err
	}
	var rowColumnTypeMap = make(map[int]rowColumnType)
	for i, columnType := range columnTypes {
		nullable, ok := columnType.Nullable()
		rowColumnTypeMap[i] = rowColumnType{
			index:            i,
			databaseTypeName: columnType.DatabaseTypeName(),
			noNull:           ok && !nullable,
		}
	}

	if rows.Next() {
		box, convert := createColBox(v, tP, cfm, rowColumnTypeMap)
		err = rows.Scan(box...)
		if err != nil {
			return
		}
		err = convert()
		if err != nil {
			return
		}
		num++
	}

	if rows.Next() {
		return 0, errors.New("result to many for one")
	}
	return
}

// Scan
// 接收多行结果
func (ctx ormContext) Scan(rows *sql.Rows) (int64, error) {
	defer func(rows *sql.Rows) {
		utils.PanicErr(rows.Close())
	}(rows)

	var num int64 = 0
	t := ctx.destBaseType
	arr := ctx.scanV
	isPtr := ctx.destSliceItemIsPtr

	columns, err := rows.Columns()
	if err != nil {
		return 0, err
	}
	cfm := getColIndex2FieldNameMap(columns, t)

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return 0, err
	}

	var rowColumnTypeMap = make(map[int]rowColumnType)
	for i, columnType := range columnTypes {
		nullable, ok := columnType.Nullable()
		rowColumnTypeMap[i] = rowColumnType{
			index:            i,
			databaseTypeName: columnType.DatabaseTypeName(),
			noNull:           ok && !nullable,
		}
	}

	for rows.Next() {
		box, vp, v, convert := createColBoxNew(t, cfm, rowColumnTypeMap)

		err = rows.Scan(box...)
		if err != nil {
			return 0, err
		}
		err = convert()
		if err != nil {
			return 0, err
		}

		if isPtr {
			arr.Set(reflect.Append(arr, vp))
		} else {
			arr.Set(reflect.Append(arr, v))
		}
		num++
	}
	return num, nil
}
