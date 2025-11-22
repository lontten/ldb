package ldb

func (q *StmtQueryContext[T]) ScanOne(dest *T) (int64, error) {
	db := q.db
	ctx := db.getCtx()

	ctx.initScanDestOne(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}

	rows, err := q.db.query(q.args...)
	if err != nil {
		return 0, err
	}
	return ctx.ScanLn(rows)
}

func (q *StmtQueryContext[T]) ScanList(dest *[]T) (int64, error) {
	db := q.db
	ctx := db.getCtx()

	ctx.initScanDestList(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}
	rows, err := q.db.query(q.args...)
	if err != nil {
		return 0, err
	}
	return ctx.ScanLn(rows)
}

func (q *StmtQueryContext[T]) ScanListP(dest *[]T) (int64, error) {
	db := q.db
	ctx := db.getCtx()

	ctx.initScanDestList(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}
	rows, err := q.db.query(q.args...)
	if err != nil {
		return 0, err
	}
	return ctx.ScanLn(rows)
}
