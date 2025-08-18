package ldb

type NativeQueryContext struct {
	db    Engine
	query string
	args  []any
}

func NativeQueryScan(db Engine, query string, args ...any) NativeQueryContext {
	return NativeQueryContext{
		db:    db.init(),
		query: query,
		args:  args,
	}
}

func (q NativeQueryContext) ScanOne(dest any) (num int64, err error) {
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
	return ctx.ScanLnT(rows)
}

func (q NativeQueryContext) ScanList(dest any) (num int64, err error) {
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
	return ctx.ScanT(rows)
}
