package ldb

import (
	"fmt"

	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/ldb/v2/utils"
)

// BetweenDateTimeOfDate
// 用 Date类型，去查询 DateTime 字段
func (w *WhereBuilder) BetweenDateTimeOfDate(query string, dateBegin, dateEnd *types.LocalDate, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, c := range condition {
		if !c {
			return w
		}
	}

	if dateBegin != nil {
		if dateEnd != nil {
			if dateEnd.Before(*dateBegin) {
				w.err = fmt.Errorf("invalid use of BetweenDateTimeOfDate: dateEnd is before dateBegin for query '%s'", query)
				return w
			}
		}
		dateTimeBegin := dateBegin.ToDateTime()
		w.Gte(query, dateTimeBegin)
		return w
	}

	if dateEnd != nil {
		dateTimeEnd := dateEnd.Add(types.Duration().Day(1)).ToDateTime()
		w.Lt(query, dateTimeEnd)
		return w
	}

	return w
}

//	Between
//
// 双闭区间 [a, b],  x >= a AND x <= b
// 时间类型数据不要用 Between
func (w *WhereBuilder) Between(query string, arg1, arg2 any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	has1 := !utils.IsNil(arg1)
	has2 := !utils.IsNil(arg2)

	if has1 {
		if has2 {
			w.andWheres = append(w.andWheres, WhereBuilder{
				clause: &Clause{
					Type:  Between,
					query: query,
					args:  []any{arg1, arg2},
				},
			})
			return w
		}
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  GreaterEq,
				query: query,
				args:  []any{arg1},
			},
		})
		return w
	}
	if has2 {
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  LessEq,
				query: query,
				args:  []any{arg2},
			},
		})
		return w
	}
	return w
}

// NotBetween
// 落在 [a, b] 双闭区间外的记录，即只保留满足 x < a 或 x > b 的数据
func (w *WhereBuilder) NotBetween(query string, arg1, arg2 any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	has1 := !utils.IsNil(arg1)
	has2 := !utils.IsNil(arg2)

	if has1 {
		if has2 {
			w.andWheres = append(w.andWheres, WhereBuilder{
				clause: &Clause{
					Type:  NotBetween,
					query: query,
					args:  []any{arg1, arg2},
				},
			})
			return w
		}
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  Less,
				query: query,
				args:  []any{arg1},
			},
		})
		return w
	}
	if has2 {
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  Greater,
				query: query,
				args:  []any{arg2},
			},
		})
		return w
	}
	return w
}
