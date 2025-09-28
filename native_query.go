package ldb

import "reflect"

func StmtQueryOne[T any](db Stmter, args ...any) (*T, error) {
	db = db.init()
	ctx := db.getCtx()

	dest := new(T)

	ctx.initScanDestOneT(dest)
	if ctx.err != nil {
		return nil, ctx.err
	}

	rows, err := db.query(args...)
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

func StmtQueryList[T any](db Stmter, args ...any) ([]T, error) {
	db = db.init()
	ctx := db.getCtx()

	var dest = &[]T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, false)
	if ctx.err != nil {
		return nil, ctx.err
	}

	rows, err := db.query(args...)
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

func StmtQueryListP[T any](db Stmter, args ...any) ([]*T, error) {
	db = db.init()
	ctx := db.getCtx()

	var dest = &[]*T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, true)
	if ctx.err != nil {
		return nil, ctx.err
	}

	rows, err := db.query(args...)
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

func QueryOne[T any](db Engine, query string, args ...any) (*T, error) {
	db = db.init()
	ctx := db.getCtx()

	dest := new(T)

	ctx.initScanDestOneT(dest)
	if ctx.err != nil {
		return nil, ctx.err
	}

	rows, err := db.query(query, args...)
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

func QueryList[T any](db Engine, query string, args ...any) ([]T, error) {
	db = db.init()
	ctx := db.getCtx()

	var dest = &[]T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, false)
	if ctx.err != nil {
		return nil, ctx.err
	}

	rows, err := db.query(query, args...)
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

func QueryListP[T any](db Engine, query string, args ...any) ([]*T, error) {
	db = db.init()
	ctx := db.getCtx()

	var dest = &[]*T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, true)
	if ctx.err != nil {
		return nil, ctx.err
	}

	rows, err := db.query(query, args...)
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
