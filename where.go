package ldb

import (
	"github.com/lontten/ldb/field"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type whereTokenType int

const (
	and whereTokenType = iota
	or
	native
)

// and(wheres)
// or(wheres)
// and(clause)
type whereToken struct {
	Type   whereTokenType
	wheres []whereToken
	clause Clause
}

type Clause struct {
	Type    clauseType
	query   string
	argsNum int
}

type WhereBuilder struct {
	primaryKeyValue       []any // 主键值列表
	filterPrimaryKeyValue []any // 主键值列表,过滤

	// 所有的and 组合成一个or放在 andWheres
	// 原因：当 and or 组合时，每条or都是独立的，and是组合使用的，有些反逻辑，为了使最后组成的sql更加已读，
	// 这里把所有and组合成一个or，和其他or联合使用。
	wheres []whereToken
	args   []any

	andWheres []whereToken
	andArgs   []any
}

func W() *WhereBuilder {
	return &WhereBuilder{}
}

type parseFun func(c Clause) (string, error)

func (w *WhereBuilder) toWhereToken() ([]whereToken, []any) {
	if len(w.andWheres) == 0 {
		return w.wheres, w.args
	}
	return append(w.wheres, whereToken{
		Type:   or,
		wheres: w.andWheres,
	}), append(w.args, w.andArgs...)
}

/*
*
各个语句之间的and or关系和具体的数据库无关，直接在这里实现，parse。
每个语句的具体sql生成和数据库有关，但是不需要其他参数，例如orm_config  orm_context 等，
所以，生成具体sql的方法 toSql 直接接受 外界传过来的 parseFun 处理函数，代码结构比较简单，
不然，whereBuilder 里面要有 dialecter 的两种实现，代码结构复杂
*/
func (w *WhereBuilder) toSql(f parseFun) (string, []any, error) {
	tokens, args := w.toWhereToken()
	sql, err := parse(tokens, f)
	if err != nil {
		return "", nil, err
	}
	return sql, args, nil
}

func parse(wts []whereToken, f parseFun) (string, error) {
	sb := strings.Builder{}
	start := true
	for _, wt := range wts {
		switch wt.Type {
		case native:
			result, err := f(wt.clause)
			if err != nil {
				return "", errors.Wrap(err, "parse native where")
			}
			if !start {
				sb.WriteString(" AND ")
			}
			if start {
				start = false
			}
			sb.WriteString(result)
		case and:
			if len(wt.wheres) == 0 {
				continue
			}
			result, err := parse(wt.wheres, f)
			if err != nil {
				return "", errors.Wrap(err, "parse native where")
			}

			if !start {
				sb.WriteString(" AND ")
			}
			if start {
				start = false
			}
			isMore := len(wt.wheres) > 1
			if isMore {
				sb.WriteString("(")
			}
			sb.WriteString(result)
			if isMore {
				sb.WriteString(")")
			}
		case or:
			if len(wt.wheres) == 0 {
				continue
			}
			result, err := parse(wt.wheres, f)
			if err != nil {
				return "", errors.Wrap(err, "parse native where")
			}
			if !start {
				sb.WriteString(" OR ")
			}
			if start {
				start = false
			}
			isMore := len(wt.wheres) > 1
			if isMore {
				sb.WriteString("(")
			}
			sb.WriteString(result)
			if isMore {
				sb.WriteString(")")
			}
		default:
			return "", errors.New("unknown where token type")
		}
	}

	return sb.String(), nil
}

// ------------------------------------------
func (w *WhereBuilder) And(wb *WhereBuilder, condition ...bool) *WhereBuilder {
	if wb == nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}
	tokens, args := wb.toWhereToken()

	w.andWheres = append(w.andWheres, whereToken{
		Type:   and,
		wheres: tokens,
	})
	w.andArgs = append(w.andArgs, args...)
	return w
}

func (w *WhereBuilder) Or(wb *WhereBuilder, condition ...bool) *WhereBuilder {
	if wb == nil {
		return w
	}
	for _, b := range condition {
		if !b {
			return w
		}
	}
	tokens, args := wb.toWhereToken()

	w.wheres = append(w.wheres, whereToken{
		Type:   or,
		wheres: tokens,
	})
	w.args = append(w.args, args...)
	return w
}
func (w *WhereBuilder) fieldValue(name string, v field.Value, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}

	switch v.Type {
	case field.None:
		break
	case field.Null:
		w.andWheres = append(w.andWheres, whereToken{
			Type: native,
			clause: Clause{
				Type:  IsNull,
				query: name,
			},
		})
		break
	case field.Now:
		w.andWheres = append(w.andWheres, whereToken{
			Type: native,
			clause: Clause{
				Type:  Eq,
				query: name,
			},
		})
		w.andArgs = append(w.andArgs, time.Now())
		break
	case field.UnixSecond:
		w.andWheres = append(w.andWheres, whereToken{
			Type: native,
			clause: Clause{
				Type:  Eq,
				query: name,
			},
		})
		w.andArgs = append(w.andArgs, strconv.Itoa(time.Now().Second()))
		break

	case field.UnixMilli:
		w.andWheres = append(w.andWheres, whereToken{
			Type: native,
			clause: Clause{
				Type:  Eq,
				query: name,
			},
		})
		w.andArgs = append(w.andArgs, strconv.FormatInt(time.Now().UnixMilli(), 10))
		break
	case field.UnixNano:
		w.andWheres = append(w.andWheres, whereToken{
			Type: native,
			clause: Clause{
				Type:  Eq,
				query: name,
			},
		})
		w.andArgs = append(w.andArgs, strconv.FormatInt(time.Now().UnixNano(), 10))
		break
	case field.Val:
		w.andWheres = append(w.andWheres, whereToken{
			Type: native,
			clause: Clause{
				Type:  Eq,
				query: name,
			},
		})
		w.andArgs = append(w.andArgs, v.Value)
		break
	case field.Expression:
		w.andWheres = append(w.andWheres, whereToken{
			Type: native,
			clause: Clause{
				Type:  Eq,
				query: name,
			},
		})
		w.andArgs = append(w.andArgs, v.Value)
		break
	default:
		break
	}
	return w
}

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

func (w *WhereBuilder) PrimaryKey(v ...any) *WhereBuilder {
	w.primaryKeyValue = v
	return w
}

func (w *WhereBuilder) FilterPrimaryKey(v ...any) *WhereBuilder {
	w.filterPrimaryKeyValue = v
	return w
}

//------------------model/map/id-end------------------
//------------------eq------------------

func (w *WhereBuilder) Eq(query string, arg any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	arg = getFieldInterZero(reflect.ValueOf(arg))
	if arg == nil {
		return w
	}
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  Eq,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg)
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
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:    In,
			query:   query,
			argsNum: argsLen,
		},
	})
	w.andArgs = append(w.andArgs, args...)
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

	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:    NotIn,
			query:   query,
			argsNum: argsLen,
		},
	})
	w.andArgs = append(w.andArgs, args...)
	return w
}

func (w *WhereBuilder) NotEq(query string, arg any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	arg = getFieldInterZero(reflect.ValueOf(arg))
	if arg == nil {
		return w
	}
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  Neq,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg)
	return w
}

func (w *WhereBuilder) Contains(query string, arg any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	arg = getFieldInterZero(reflect.ValueOf(arg))
	if arg == nil {
		return w
	}
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  Contains,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg)
	return w
}

// 小于
func (w *WhereBuilder) Less(query string, arg any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	arg = getFieldInterZero(reflect.ValueOf(arg))
	if arg == nil {
		return w
	}
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  Less,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg)
	return w
}

// 小于等于
func (w *WhereBuilder) LessEq(query string, arg any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	arg = getFieldInterZero(reflect.ValueOf(arg))
	if arg == nil {
		return w
	}
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  LessEq,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg)
	return w
}

// 大于
func (w *WhereBuilder) Greater(query string, arg any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	arg = getFieldInterZero(reflect.ValueOf(arg))
	if arg == nil {
		return w
	}
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  Greater,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg)
	return w
}

// 大于等于
func (w *WhereBuilder) GreaterEq(query string, arg any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	arg = getFieldInterZero(reflect.ValueOf(arg))
	if arg == nil {
		return w
	}
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  GreaterEq,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg)
	return w
}

func (w *WhereBuilder) Between(query string, arg1, arg2 any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}

	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  Between,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg1, arg2)
	return w
}

func (w *WhereBuilder) Arg(arg any, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}
	w.andArgs = append(w.andArgs, arg)
	return w
}

func (w *WhereBuilder) Args(args ...any) *WhereBuilder {
	w.andArgs = append(w.andArgs, args...)
	return w
}

func (w *WhereBuilder) IsNull(query string, condition ...bool) *WhereBuilder {
	for _, b := range condition {
		if !b {
			return w
		}
	}

	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
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

	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
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
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
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

	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  NotBetween,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg1, arg2)
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
	arg = getFieldInterZero(reflect.ValueOf(arg))
	if arg == nil {
		return w
	}
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  Neq,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, arg)
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
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  Like,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, "%"+*arg+"%")
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
	w.andWheres = append(w.andWheres, whereToken{
		Type: native,
		clause: Clause{
			Type:  NotLike,
			query: query,
		},
	})
	w.andArgs = append(w.andArgs, "%"+*arg+"%")
	return w
}
