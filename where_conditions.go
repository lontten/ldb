package ldb

import (
	"reflect"

	"github.com/lontten/ldb/v2/utils"
)

//------------------model/map/id------------------

// 过滤 软删除
func (w *WhereBuilder) Model(v any, condition ...bool) *WhereBuilder {
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

// 参数为nil，自动跳过条件
func (w *WhereBuilder) Eq(query string, arg any, condition ...bool) *WhereBuilder {
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

func (w *WhereBuilder) In(query string, args ArgArray, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
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

func (w *WhereBuilder) NotIn(query string, args ArgArray, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
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

// 参数为nil，不生成条件
func (w *WhereBuilder) NotEq(query string, arg any, condition ...bool) *WhereBuilder {
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

func (w *WhereBuilder) Contains(query string, arg any, condition ...bool) *WhereBuilder {
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

// 小于
func (w *WhereBuilder) Less(query string, arg any, condition ...bool) *WhereBuilder {
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
			Type:  Less,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// 小于等于
func (w *WhereBuilder) LessEq(query string, arg any, condition ...bool) *WhereBuilder {
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
			Type:  LessEq,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// 大于
func (w *WhereBuilder) Greater(query string, arg any, condition ...bool) *WhereBuilder {
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
			Type:  Greater,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

// 大于等于
func (w *WhereBuilder) GreaterEq(query string, arg any, condition ...bool) *WhereBuilder {
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
			Type:  GreaterEq,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

func (w *WhereBuilder) Between(query string, arg1, arg2 any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  Between,
			query: query,
			args:  []any{arg1, arg2},
		},
	})
	return w
}

//func (w *WhereBuilder) Arg(arg any, condition ...bool) *WhereBuilder {
//	for _, b := range condition {
//		if !b {
//			return w
//		}
//	}
//	w.andArgs = append(w.andArgs, arg)
//	return w
//}
//
//func (w *WhereBuilder) Args(args ...any) *WhereBuilder {
//	w.andArgs = append(w.andArgs, args...)
//	return w
//}

func (w *WhereBuilder) IsNull(query string, condition ...bool) *WhereBuilder {
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

func (w *WhereBuilder) IsNotNull(query string, condition ...bool) *WhereBuilder {
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

func (w *WhereBuilder) IsFalse(query string, condition ...bool) *WhereBuilder {
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
func (w *WhereBuilder) NotBetween(query string, arg1, arg2 any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}

	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  NotBetween,
			query: query,
			args:  []any{arg1, arg2},
		},
	})
	return w
}

// Neq
// 不等于
// 当 arg 为 nil 时，不添加条件
func (w *WhereBuilder) Neq(query string, arg any, condition ...bool) *WhereBuilder {
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
			Type:  Neq,
			query: query,
			args:  []any{arg},
		},
	})
	return w
}

func (w *WhereBuilder) Like(query string, arg *string, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	if arg == nil {
		return w
	}
	if *arg == "" {
		return w
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  Like,
			query: query,
			args:  []any{"%" + *arg + "%"},
		},
	})
	return w
}

func (w *WhereBuilder) NoLike(query string, arg *string, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	if arg == nil {
		return w
	}
	if *arg == "" {
		return w
	}
	w.andWheres = append(w.andWheres, WhereBuilder{
		clause: &Clause{
			Type:  NotLike,
			query: query,
			args:  []any{"%" + *arg + "%"},
		},
	})
	return w
}
