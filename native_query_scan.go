package ldb

func (q *NativeQueryContext[T]) ScanOne(dest any) (num int64, err error) {
	db := q.db
	ctx := db.getCtx()
	ctx.initScanDestOne(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}

	query := q.query
	args := q.args

	rows, err := db.query(query, args...)
	if err != nil {
		return 0, err
	}
	return ctx.ScanLn(rows)
}

func (q *NativeQueryContext[T]) ScanList(dest any) (num int64, err error) {
	db := q.db
	ctx := db.getCtx()
	ctx.initScanDestList(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}

	query := q.query
	args := q.args

	rows, err := db.query(query, args...)
	if err != nil {
		return 0, err
	}
	return ctx.Scan(rows)
}
