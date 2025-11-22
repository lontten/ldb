package ldb

import (
	"database/sql"
	"fmt"
	"reflect"
)

func (b *SqlBuilder[T]) ScanOne(dest any) (rowsNum int64, err error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return 0, ctx.err
	}

	ctx.initScanDestOne(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}

	b.initSelectSql()
	query := b.query
	args := b.args
	if ctx.showSql {
		fmt.Println(query, args)
	}
	if ctx.noRun {
		return 0, nil
	}

	rows, err := db.query(query, args...)
	if err != nil {
		return 0, err
	}
	return ctx.ScanLn(rows)
}

func (b *SqlBuilder[T]) ScanList(dest any) (rowsNum int64, err error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return 0, ctx.err
	}

	ctx.initScanDestList(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}
	b.initSelectSql()

	query := b.query
	args := b.args

	if ctx.showSql {
		fmt.Println(query, args)
	}
	if ctx.noRun {
		return 0, nil
	}
	rows, err := db.query(query, args...)
	if err != nil {
		return 0, err
	}
	return ctx.Scan(rows)
}

func (b *SqlBuilder[T]) Exec() (sql.Result, error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return nil, ctx.err
	}
	b.initSelectSql()
	if ctx.showSql {
		fmt.Println(b.query, b.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	return db.exec(b.query, b.args...)
}

func (b *SqlBuilder[T]) One() (t *T, err error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	if ctx.hasErr() {
		return nil, ctx.err
	}

	dest := new(T)
	ctx.initScanDestOneT(dest)
	if ctx.err != nil {
		return nil, ctx.err
	}

	b.initSelectSql()

	dialect.getSql(b.query)
	ctx.originalArgs = b.args
	ctx.printSql()

	if ctx.noRun {
		return nil, nil
	}

	rows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
	if err != nil {
		return nil, err
	}
	num, err := ctx.ScanLn(rows)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return nil, nil
	}
	return dest, nil
}

func (b *SqlBuilder[T]) List() (list []T, err error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	if ctx.hasErr() {
		return nil, ctx.err
	}

	var dest = &[]T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, false)
	if ctx.err != nil {
		return nil, ctx.err
	}

	b.initSelectSql()

	dialect.getSql(b.query)
	ctx.originalArgs = b.args
	ctx.printSql()

	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
	if err != nil {
		return nil, err
	}
	num, err := ctx.Scan(rows)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return nil, nil
	}
	return *dest, nil
}

func (b *SqlBuilder[T]) ListP() (list []*T, err error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	if ctx.hasErr() {
		return nil, ctx.err
	}

	var dest = &[]*T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, true)
	if ctx.err != nil {
		return nil, ctx.err
	}

	b.initSelectSql()

	dialect.getSql(b.query)
	ctx.originalArgs = b.args
	ctx.printSql()

	if ctx.noRun {
		return nil, nil
	}
	rows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
	if err != nil {
		return nil, err
	}
	num, err := ctx.Scan(rows)
	if err != nil {
		return nil, err
	}
	if num == 0 {
		return nil, nil
	}
	return *dest, nil
}
