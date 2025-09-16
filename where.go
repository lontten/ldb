package ldb

import (
	"strconv"
	"time"

	"github.com/lontten/ldb/v2/field"
)

type parseFun func(c Clause) (string, error)

type Clause struct {
	Type  clauseType
	query string
	pks   []string // 多主键情况，主键字段名称列表
	args  []any
}

type WhereBuilder struct {
	not bool

	// 所有的and 组合成一个or放在 andWheres
	// 原因：当 and or 组合时，每条or都是独立的，and是组合使用的，有些反逻辑，为了使最后组成的sql更加易读，
	// 这里把所有and组合成一个or，和其他or联合使用。
	wheres []WhereBuilder

	andWheres []WhereBuilder

	clause *Clause
	err    error
}

func W() *WhereBuilder {
	return &WhereBuilder{}
}

func (w WhereBuilder) has() bool {
	if len(w.wheres) > 0 {
		return true
	}
	if len(w.andWheres) > 0 {
		return true
	}
	return false
}

func (w WhereBuilder) Invalid() bool {
	return len(w.wheres) == 0 && len(w.andWheres) == 0 && w.clause == nil
}

// ------------------------------------------

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
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  IsNull,
				query: name,
			},
		})
		break
	case field.Now:
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  Eq,
				query: name,
				args:  []any{time.Now()},
			},
		})
		break
	case field.UnixSecond:
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  Eq,
				query: name,
				args:  []any{strconv.Itoa(time.Now().Second())},
			},
		})
		break

	case field.UnixMilli:
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  Eq,
				query: name,
				args:  []any{strconv.FormatInt(time.Now().UnixMilli(), 10)},
			},
		})
		break
	case field.UnixNano:
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  Eq,
				query: name,
				args:  []any{strconv.FormatInt(time.Now().UnixNano(), 10)},
			},
		})
		break
	case field.Val:
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  Eq,
				query: name,
				args:  []any{v.Value},
			},
		})
		break
	case field.Expression:
		w.andWheres = append(w.andWheres, WhereBuilder{
			clause: &Clause{
				Type:  Eq,
				query: name,
				args:  []any{v.Value},
			},
		})
		break
	default:
		break
	}
	return w
}
