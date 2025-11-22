package ldb

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/lontten/ldb/v2/utils"
)

// ListPage 查询分页
func (b *SqlBuilder[T]) ScanPage(dest *[]T) (dto PageResult[T], err error) {
	db := b.db
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	if err = ctx.err; err != nil {
		return
	}
	if b.pageConfig == nil {
		err = errors.New("no set pageConfig")
		return
	}
	var total int64
	var pageSize = b.pageConfig.pageSize
	var pageIndex = b.pageConfig.pageIndex

	ctx.initScanDestList(dest)
	if err = ctx.err; err != nil {
		return
	}

	b.initSelectSql()

	var countSql = b.countField
	if countSql == "" {
		countSql = "*"
	}

	countSql = "select count(" + countSql + ") " + b.otherSqlBuilder.String()

	dialect.getSql(countSql)
	ctx.originalArgs = b.otherSqlArgs
	ctx.printSql()

	if !ctx.noRun {
		if b.fakeTotalNum > 0 {
			total = b.fakeTotalNum
		} else {
			rows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
			if err != nil {
				return dto, err
			}
			defer func(rows *sql.Rows) {
				utils.PanicErr(rows.Close())
			}(rows)
			for rows.Next() {
				box := reflect.ValueOf(&total).Interface()
				err = rows.Scan(box)
				if err != nil {
					return dto, err
				}
			}
		}
	}

	// 计算总页数
	var pageNum = total / pageSize
	if total%pageSize != 0 {
		pageNum++
	}

	var selectSql = b.query + " limit ? offset ?"
	var offset = (pageIndex - int64(1)) * pageSize
	args := append(b.args, pageSize, offset)

	dialect.getSql(selectSql)
	ctx.originalArgs = args
	ctx.printSql()

	if ctx.noRun {
		return dto, nil
	}
	if b.noGetList {
		dto = PageResult[T]{
			List:      *dest,
			PageSize:  pageSize,
			PageNum:   pageNum,
			PageIndex: pageIndex,
			Total:     total,
			HasMore:   total > pageSize*pageIndex,
		}
		return dto, nil
	}

	listRows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
	if err != nil {
		return
	}

	_, err = ctx.Scan(listRows)
	if err != nil {
		return
	}

	dto = PageResult[T]{
		List:      *dest,
		PageSize:  pageSize,
		PageNum:   pageNum,
		PageIndex: pageIndex,
		Total:     total,
		HasMore:   total > pageSize*pageIndex,
	}
	return dto, nil
}

// ListPage 查询分页
func (b *SqlBuilder[T]) ScanPageP(dest *[]*T) (dto PageResultP[T], err error) {
	db := b.db
	dialect := db.getDialect()
	ctx := dialect.getCtx()
	if err = ctx.err; err != nil {
		return
	}
	if b.pageConfig == nil {
		err = errors.New("no set pageConfig")
		return
	}
	var total int64
	var pageSize = b.pageConfig.pageSize
	var pageIndex = b.pageConfig.pageIndex

	ctx.initScanDestList(dest)
	if err = ctx.err; err != nil {
		return
	}

	b.initSelectSql()

	var countSql = b.countField
	if countSql == "" {
		countSql = "*"
	}

	countSql = "select count(" + countSql + ") " + b.otherSqlBuilder.String()

	dialect.getSql(countSql)
	ctx.originalArgs = b.otherSqlArgs
	ctx.printSql()

	if !ctx.noRun {
		if b.fakeTotalNum > 0 {
			total = b.fakeTotalNum
		} else {
			rows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
			if err != nil {
				return dto, err
			}
			defer func(rows *sql.Rows) {
				utils.PanicErr(rows.Close())
			}(rows)
			for rows.Next() {
				box := reflect.ValueOf(&total).Interface()
				err = rows.Scan(box)
				if err != nil {
					return dto, err
				}
			}
		}
	}

	// 计算总页数
	var pageNum = total / pageSize
	if total%pageSize != 0 {
		pageNum++
	}

	var selectSql = b.query + " limit ? offset ?"
	var offset = (pageIndex - int64(1)) * pageSize
	args := append(b.args, pageSize, offset)

	dialect.getSql(selectSql)
	ctx.originalArgs = args
	ctx.printSql()

	if ctx.noRun {
		return dto, nil
	}
	if b.noGetList {
		dto = PageResultP[T]{
			List:      *dest,
			PageSize:  pageSize,
			PageNum:   pageNum,
			PageIndex: pageIndex,
			Total:     total,
			HasMore:   total > pageSize*pageIndex,
		}
		return dto, nil
	}

	listRows, err := db.query(ctx.dialectSql, ctx.originalArgs...)
	if err != nil {
		return
	}

	_, err = ctx.Scan(listRows)
	if err != nil {
		return
	}

	dto = PageResultP[T]{
		List:      *dest,
		PageSize:  pageSize,
		PageNum:   pageNum,
		PageIndex: pageIndex,
		Total:     total,
		HasMore:   total > pageSize*pageIndex,
	}
	return dto, nil
}
