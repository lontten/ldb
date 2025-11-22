package ldb

import (
	"database/sql"
	"errors"
	"reflect"

	"github.com/lontten/ldb/v2/utils"
)

// CountField 自定义count字段
func (b *SqlBuilder[T]) CountField(field string, conditions ...bool) *SqlBuilder[T] {
	for _, c := range conditions {
		if !c {
			return b
		}
	}
	b.countField = field
	return b
}

// FakerTotalNum 分页时，直接使用 fakeTotalNum，不再查询实际总数
func (b *SqlBuilder[T]) FakerTotalNum(num int64, conditions ...bool) *SqlBuilder[T] {
	for _, c := range conditions {
		if !c {
			return b
		}
	}
	b.fakeTotalNum = num
	return b
}

// NoGetList 分页时，只查询数量，不返回数据列表
func (b *SqlBuilder[T]) NoGetList(conditions ...bool) *SqlBuilder[T] {
	for _, c := range conditions {
		if !c {
			return b
		}
	}
	b.noGetList = true
	return b
}

func (b *SqlBuilder[T]) Page(pageIndex int64, pageSize int64) *SqlBuilder[T] {
	if pageSize < 1 || pageIndex < 1 {
		b.db.getCtx().err = errors.New("pageSize,pageIndex must be greater than 0")
	}
	b.pageConfig = &PageConfig{
		pageSize:  pageSize,
		pageIndex: pageIndex,
	}
	return b
}

type PageConfig struct {
	pageSize  int64
	pageIndex int64
}

type PageResult[T any] struct {
	List      []T   `json:"list"`      // 结果
	PageSize  int64 `json:"pageSize"`  // 每页大小
	PageIndex int64 `json:"pageIndex"` // 当前页码
	Total     int64 `json:"total"`     // 总数
	PageNum   int64 `json:"totalPage"` // 总页数
	HasMore   bool  `json:"hasMore"`   // 是否有更多
}

type PageResultP[T any] struct {
	List      []*T  `json:"list"`      // 结果
	PageSize  int64 `json:"pageSize"`  // 每页大小
	PageIndex int64 `json:"pageIndex"` // 当前页码
	Total     int64 `json:"total"`     // 总数
	PageNum   int64 `json:"totalPage"` // 总页数
	HasMore   bool  `json:"hasMore"`   // 是否有更多
}

// ScanPage 查询分页
func (b *SqlBuilder[T]) ScanPage() (dto PageResult[T], err error) {
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

	var dest = &[]T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, false)
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

// PageP 查询分页
func (b *SqlBuilder[T]) PageP() (dto PageResultP[T], err error) {
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

	var dest = &[]*T{}
	v := reflect.ValueOf(dest).Elem()
	baseV := reflect.ValueOf(new(T)).Elem()
	t := baseV.Type()

	ctx.initScanDestListT(dest, v, baseV, t, true)
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
