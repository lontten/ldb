package ldb

func (b *SqlBuilder[T]) ScanOne(dest any) (rowsNum int64, err error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	if ctx.hasErr() {
		return 0, ctx.err
	}

	ctx.initScanDestOne(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}

	b.initSelectSql()
	dialect.getSql(b.query)
	ctx.originalArgs = b.args
	ctx.printSql()

	if ctx.noRun {
		return 0, nil
	}

	rows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
	if err != nil {
		return 0, err
	}
	return ctx.ScanLn(rows)
}

func (b *SqlBuilder[T]) ScanList(dest *[]T) (rowsNum int64, err error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	if ctx.hasErr() {
		return 0, ctx.err
	}

	ctx.initScanDestList(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}
	b.initSelectSql()

	dialect.getSql(b.query)
	ctx.originalArgs = b.args
	ctx.printSql()

	if ctx.noRun {
		return 0, nil
	}
	rows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
	if err != nil {
		return 0, err
	}
	return ctx.Scan(rows)
}

func (b *SqlBuilder[T]) ScanListP(dest *[]*T) (rowsNum int64, err error) {
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	db := b.db
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	if ctx.hasErr() {
		return 0, ctx.err
	}

	ctx.initScanDestList(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}
	b.initSelectSql()

	dialect.getSql(b.query)
	ctx.originalArgs = b.args
	ctx.printSql()

	if ctx.noRun {
		return 0, nil
	}
	rows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
	if err != nil {
		return 0, err
	}
	return ctx.Scan(rows)
}
