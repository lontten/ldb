package ldb

import (
	"fmt"
	"reflect"

	"github.com/lontten/ldb/v2/utils"
)

//------------------model/map/id------------------

// 过滤 软删除
func (w *WhereBuilder) Model(v any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	list := getStructCV(reflect.ValueOf(v))
	for _, cv := range list {
		if cv.isSoftDel || cv.isZero {
			continue
		}
		w.fieldValue(cv.columnName, cv.value)
	}
	return w
}

func (w *WhereBuilder) Map(v any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	list := getMapCV(reflect.ValueOf(v))
	for _, cv := range list {
		w.fieldValue(cv.columnName, cv.value)
	}
	return w
}
func (w *WhereBuilder) PrimaryKey(args ...any) *WhereBuilder {
	if w.err != nil {
		return w
	}
	argsLen := len(args)
	if argsLen == 0 {
		return w
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type: PrimaryKeys,
			args: args,
		},
	})
	return w
}

func (w *WhereBuilder) FilterPrimaryKey(args ...any) *WhereBuilder {
	if w.err != nil {
		return w
	}
	argsLen := len(args)
	if argsLen == 0 {
		return w
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type: FilterPrimaryKeys,
			args: args,
		},
	})
	return w
}

//------------------model/map/id-end------------------
//------------------eq------------------

// Eq
// x = ?
func (w *WhereBuilder) Eq(query string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	isNil := utils.IsNil(arg)
	if isNil {
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  IsNull,
				query: query,
			},
		})
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  Eq,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// NotEq
// x <> ?
func (w *WhereBuilder) NotEq(query string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	isNil := utils.IsNil(arg)
	if isNil {
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  IsNotNull,
				query: query,
			},
		})
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  Neq,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// BoolIn
// IN (?)
func (w *WhereBuilder) BoolIn(condition bool, query string, args ...any) *WhereBuilder {
	if w.err != nil {
		return w
	}
	if !condition {
		return w
	}
	argsLen := len(args)
	if argsLen == 0 {
		return w
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  In,
			query: query,
			args:  args,
		},
	})
	return w
}

// In
// IN (?)
func (w *WhereBuilder) In(query string, args ...any) *WhereBuilder {
	if w.err != nil {
		return w
	}
	argsLen := len(args)
	if argsLen == 0 {
		return w
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  In,
			query: query,
			args:  args,
		},
	})
	return w
}

// BoolNotIn
// NOT IN (?)
func (w *WhereBuilder) BoolNotIn(condition bool, query string, args ...any) *WhereBuilder {
	if w.err != nil {
		return w
	}
	if !condition {
		return w
	}
	argsLen := len(args)
	if argsLen == 0 {
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  NotIn,
			query: query,
			args:  args,
		},
	})
	return w
}

// NotIn
// NOT IN (?)
func (w *WhereBuilder) NotIn(query string, args ...any) *WhereBuilder {
	if w.err != nil {
		return w
	}
	argsLen := len(args)
	if argsLen == 0 {
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  NotIn,
			query: query,
			args:  args,
		},
	})
	return w
}

// Contains
// pg 独有
// [1] @< [1,2]
func (w *WhereBuilder) Contains(query string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	isNil := utils.IsNil(arg)
	if isNil {
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  Contains,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// Lt
// x < a
func (w *WhereBuilder) Lt(query string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	isNil := utils.IsNil(arg)
	if isNil {
		w.err = fmt.Errorf("invalid use of Lt: argument for query '%s' is nil.", query)
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  Less,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// Lte
// x <= a
func (w *WhereBuilder) Lte(query string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	isNil := utils.IsNil(arg)
	if isNil {
		w.err = fmt.Errorf("invalid use of Lte: argument for query '%s' is nil.", query)
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  LessEq,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// Gt
// x > a
func (w *WhereBuilder) Gt(query string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	isNil := utils.IsNil(arg)
	if isNil {
		w.err = fmt.Errorf("invalid use of Gt: argument for query '%s' is nil.", query)
		return w
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  Greater,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// Gte
// x >= a
func (w *WhereBuilder) Gte(query string, arg any, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}
	isNil := utils.IsNil(arg)
	if isNil {
		w.err = fmt.Errorf("invalid use of Gte: argument for query '%s' is nil.", query)
		return w
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  GreaterEq,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// IsNull
// x IS NULL
func (w *WhereBuilder) IsNull(query string, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  IsNull,
			query: query,
		},
	})
	return w
}

// IsNotNull
// x IS NOT NULL
func (w *WhereBuilder) IsNotNull(query string, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  IsNotNull,
			query: query,
		},
	})
	return w
}

// IsFalse
// x IS FALSE
func (w *WhereBuilder) IsFalse(query string, condition ...bool) *WhereBuilder {
	if w.err != nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  IsFalse,
			query: query,
		},
	})
	return w
}
