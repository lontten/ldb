package ldb

import (
	"errors"
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
