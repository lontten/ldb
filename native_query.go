package ldb

import "reflect"

type NativeQueryContext[T any] struct {
	db    Engine
	query string
	args  []any
}

func NativeQuery[T any](db Engine, query string, args ...any) *NativeQueryContext[T] {
	return &NativeQueryContext[T]{
		db:    db.init(),
		query: query,
		args:  args,
	}
}
func (q *NativeQueryContext[T]) Convert(c Convert) *NativeQueryContext[T] {
	q.db.getCtx().convertCtx.Add(c)
	return q
}

func (q *NativeQueryContext[T]) One() (t *T, err error) {
	db := q.db
	query := q.query
	args := q.args
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

func (q *NativeQueryContext[T]) List() ([]T, error) {
	db := q.db
	query := q.query
	args := q.args
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
	_, err = ctx.Scan(rows)
	if err != nil {
		return nil, err
	}
	return *dest, nil
}

func (q *NativeQueryContext[T]) ListP() ([]*T, error) {
	db := q.db
	query := q.query
	args := q.args
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
	_, err = ctx.Scan(rows)
	if err != nil {
		return nil, err
	}
	return *dest, nil
}
