package ldb

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/lontten/ldb/v2/utils"
	"github.com/pkg/errors"
)

// ScanLnT
// 接收一行结果
// 1.ptr single/comp
// 2.slice- single
func (ctx ormContext) ScanLnT(rows *sql.Rows) (num int64, err error) {
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

	if rows.Next() {
		box, convert := createColBoxT(v, tP, cfm)
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

// ScanT
// 接收多行结果
func (ctx ormContext) ScanT(rows *sql.Rows) (int64, error) {
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
	for rows.Next() {
		box, vp, v := createColBox(t, cfm)

		err = rows.Scan(box...)
		if err != nil {
			fmt.Println(err)
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
