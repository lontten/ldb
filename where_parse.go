package ldb

import (
	"reflect"
	"strings"

	"github.com/lontten/lcore/lcutils"
	"github.com/pkg/errors"
)

/*
*
各个语句之间的and or关系和具体的数据库无关，直接在这里实现，parse。
每个语句的具体sql生成和数据库有关，但是不需要其他参数，例如orm_config  orm_context 等，
所以，生成具体sql的方法 toSql 直接接受 外界传过来的 parseFun 处理函数，代码结构比较简单，
不然，whereBuilder 里面要有 dialecter 的两种实现，代码结构复杂

primaryKeyFieldNames 主键字段名称列表
*/

func (w WhereBuilder) toSql(f parseFun, primaryKeyColumnNames ...string) (string, []any, error) {
	_, sql, args, err := w.parse(f, primaryKeyColumnNames...)
	if err != nil {
		return "", nil, err
	}
	return sql, args, err
}
func (w *WhereBuilder) parsePkClause(primaryKeyColumnNames ...string) error {
	if len(primaryKeyColumnNames) == 0 {
		return ErrNoPk
	}
	var args = w.clause.args
	argsLen := len(args)
	if argsLen == 0 {
		return nil
	}

	// 0未设置；1struct复合主键;2map复合主键;3单主键
	var kindType = 0

	var nw = W()
	for _, arg := range args {
		var nw2 = W()

		var v = reflect.ValueOf(arg)
		_, v, err := basePtrValue(v)
		if err != nil {
			return err
		}
		kind := v.Kind()

		if kind == reflect.Struct {
			if kindType == 0 {
				kindType = 1
			} else {
				if kindType != 1 {
					return ErrTypePkArgs
				}
			}
			list := getStructCV(v)
			for _, cv := range list {
				if cv.isSoftDel || cv.isZero {
					continue
				}
				if lcutils.StrContainsAny(cv.columnName, primaryKeyColumnNames...) {
					nw2.fieldValue(cv.columnName, cv.value)
				}
			}

		} else if kind == reflect.Map {
			if kindType == 0 {
				kindType = 2
			} else {
				if kindType != 2 {
					return ErrTypePkArgs
				}
			}
			list := getMapCV(v)
			for _, cv := range list {
				if cv.isSoftDel || cv.isZero {
					continue
				}
				if lcutils.StrContainsAny(cv.columnName, primaryKeyColumnNames...) {
					nw2.fieldValue(cv.columnName, cv.value)
				}
			}
		} else {
			kindType = 3
		}

		nw.Or(nw2)
	}

	if kindType == 3 {
		if len(primaryKeyColumnNames) != 1 {
			return ErrNeedMultiPk
		}
		nw.In(primaryKeyColumnNames[0], args)
	}

	w.clause = nil
	w.And(nw)
	return nil
}
func (w WhereBuilder) parse(f parseFun, primaryKeyFieldNames ...string) (hasOr bool, sql string, args []any, err error) {
	sb := strings.Builder{}
	var ors = w.wheres
	var ands = w.andWheres
	var allArgs []any

	var orLen = len(ors)
	var andLen = len(ands)

	var logicOrLen = orLen
	if len(ands) > 0 {
		logicOrLen++
	}

	if w.clause != nil {
		var c = *w.clause
		if c.Type == PrimaryKeys || c.Type == FilterPrimaryKeys {
			var _err = w.parsePkClause(primaryKeyFieldNames...)
			if _err != nil {
				err = _err
				return
			}
			return w.parse(f, primaryKeyFieldNames...)
		} else {
			result, _err := f(c)
			if _err != nil {
				err = errors.Wrap(_err, "parse WhereBuilder")
				return
			}
			sb.WriteString(result)
			return false, sb.String(), c.args, nil
		}
	}

	if !w.has() {
		return
	}

	for i, wt := range ors {
		var _hasOr, _sql, _args, _err = wt.parse(f, primaryKeyFieldNames...)
		if _err != nil {
			err = errors.Wrap(_err, "parse WhereBuilder")
			return
		}
		if _hasOr {
			hasOr = true
		}
		allArgs = append(allArgs, _args...)
		if logicOrLen > 1 {
			sb.WriteString("(")
		}
		sb.WriteString(_sql)
		if logicOrLen > 1 {
			sb.WriteString(")")
		}
		if i < orLen-1 {
			sb.WriteString(" OR ")
		}
	}

	var needOr = orLen > 0 && andLen > 0

	if needOr {
		sb.WriteString(" OR ")
		sb.WriteString("(")
		hasOr = true
	}

	for i, wt := range ands {
		var _hasOr, _sql, _args, _err = wt.parse(f, primaryKeyFieldNames...)
		if _err != nil {
			err = errors.Wrap(_err, "parse WhereBuilder")
			return
		}
		if _hasOr {
			hasOr = true
		}
		allArgs = append(allArgs, _args...)
		if andLen > 1 && hasOr {
			sb.WriteString("(")
		}
		sb.WriteString(_sql)
		if andLen > 1 && hasOr {
			sb.WriteString(")")
		}
		if i < andLen-1 {
			sb.WriteString(" AND ")
		}
	}

	if needOr {
		sb.WriteString(")")
	}
	return hasOr, sb.String(), allArgs, nil
}
