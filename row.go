package ldb

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

// 创建 row 返回数据，字段 对应的 struct 字段的box
// 返回值 box, vp, v
// box	struct 的 字段box列表
// vp	struct 的 引用
// v	struct 的 值
func createColBoxNew(base reflect.Type, cfLink map[string]compC, rowColumnTypeMap map[int]rowColumnType) (box []any, vp, v reflect.Value, fun func() error) {
	vp = reflect.New(base)
	v = reflect.Indirect(vp)
	tP := vp.Interface()

	colBox, fun := createColBox(v, tP, cfLink, rowColumnTypeMap)
	return colBox, vp, v, fun
}

// 创建 row 返回数据，字段 对应的 struct 字段的box
// 返回值 box, vp, v
// box	struct 的 字段 引用列表
// vp	struct 的 引用 Value
// v	struct 的 值   Value
func createColBox(v reflect.Value, tP any, cfLink map[string]compC, rowColumnTypeMap map[int]rowColumnType) (box []any, fun func() error) {
	fun = func() error { return nil }
	length := len(cfLink)
	if length == 0 {
		box = make([]any, 1)
		box[0] = tP
		return
	}

	box = make([]any, length)
	var converters []func() error
	fun = func() error {
		for _, f := range converters {
			e := f()
			if e != nil {
				return e
			}
		}
		return nil
	}

	for _, f := range cfLink {
		if f.columnName == "" { // "" 表示此列不接收
			box[f.columnIndex] = new(any)
			continue
		}

		field := v.FieldByName(f.fieldName)
		columnType := rowColumnTypeMap[f.columnIndex]
		// 字段可以接收null 或者 返回值不为null;可以直接用 结构体字段接收
		if f.canNull || columnType.noNull {
			box[f.columnIndex] = field.Addr().Interface()
			continue
		}

		var tmpVal any
		if f.isScanner {
			tmpVal = new(any)
		} else {
			tmpVal = allocDatabaseType(columnType.databaseTypeName)
			if tmpVal == nil {
				tmpVal = allocType(field.Type())
				if tmpVal == nil {
					panic("field not support")
				}
			}
		}
		box[f.columnIndex] = tmpVal

		converters = append(converters, func(f reflect.Value, fieldName string, val any, c compC) func() error {
			return func() error {
				if val == nil {
					f.SetZero()
					return nil
				}

				switch nullVal := val.(type) {
				case *any:
					if nullVal == nil {
						f.SetZero()
						return nil
					} else {
						// 如果 字段实现了 sql.Scanner，但是因为 不是 指针，就创建一个临时指针类型接收，然后再赋值
						newVal := reflect.New(f.Type())
						scanner := newVal.Interface().(sql.Scanner)
						if err := scanner.Scan(*nullVal); err != nil {
							fmt.Println(c.fieldName)
							fmt.Println(c.kind.String())
							return err
						}
						f.Set(newVal.Elem())
						return nil
					}
				case *sql.NullString:
					if nullVal.Valid {
						f.SetString(nullVal.String)
					} else {
						f.SetString("")
					}
				case *sql.NullInt64:
					if nullVal.Valid {
						f.SetInt(nullVal.Int64)
					} else {
						f.SetInt(0)
					}
				case *sql.NullFloat64:
					if nullVal.Valid {
						f.SetFloat(nullVal.Float64)
					} else {
						f.SetFloat(0)
					}
				case *sql.NullBool:
					if nullVal.Valid {
						f.SetBool(nullVal.Bool)
					} else {
						f.SetBool(false)
					}
				case *sql.NullTime:
					if nullVal.Valid {
						f.Set(reflect.ValueOf(nullVal.Time))
					} else {
						f.Set(reflect.ValueOf(time.Time{}))
					}
				case *sql.RawBytes:
					f.SetBytes(*nullVal)
				}

				return nil
			}
		}(field, f.fieldName, tmpVal, f))
	}
	return
}

// sql返回 row 字段下标 对应的  struct 字段名（""表示不接收该列数据）
type ColIndex2FieldNameMap []string
