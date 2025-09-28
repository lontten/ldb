package ldb

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
)

// 创建 row 返回数据，字段 对应的 struct 字段的box
// 返回值 box, vp, v
// box	struct 的 字段box列表
// vp	struct 的 引用
// v	struct 的 值
func createColBoxNew(base reflect.Type, cfLink map[string]compC) (box []any, vp, v reflect.Value, fun func() error) {
	vp = reflect.New(base)
	v = reflect.Indirect(vp)
	tP := vp.Interface()

	colBox, fun := createColBox(v, tP, cfLink)
	return colBox, vp, v, fun
}

// 创建 row 返回数据，字段 对应的 struct 字段的box
// 返回值 box, vp, v
// box	struct 的 字段 引用列表
// vp	struct 的 引用 Value
// v	struct 的 值   Value
func createColBox(v reflect.Value, tP any, cfLink map[string]compC) (box []any, fun func() error) {
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
		rb := new(sql.RawBytes)
		box[f.columnIndex] = rb
		if f.columnName == "" { // "" 表示此列不接收
			continue
		}
		field := v.FieldByName(f.fieldName)

		if f.canNull {
			box[f.columnIndex] = field.Addr().Interface()
			continue
		}

		converters = append(converters, func(f reflect.Value, rawBytes *sql.RawBytes, c compC) func() error {
			return func() error {
				if rawBytes == nil || len(*rawBytes) == 0 {
					f.SetZero()
					return nil
				}
				if c.isScanner {
					newVal := reflect.New(f.Type())
					scanner := newVal.Interface().(sql.Scanner)
					if err := scanner.Scan(*rawBytes); err != nil {
						return err
					}
					f.Set(newVal.Elem())
					return nil
				}

				strValue := string(*rawBytes)

				switch c.kind {
				case reflect.String:
					f.SetString(strValue)
					return nil

				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					i, err := strconv.ParseInt(strValue, 10, 64)
					if err != nil {
						return fmt.Errorf("解析整数失败 (值: %q): %w", strValue, err)
					}
					f.SetInt(i)
					return nil

				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					i, err := strconv.ParseUint(strValue, 10, 64)
					if err != nil {
						return fmt.Errorf("解析无符号整数失败 (值: %q): %w", strValue, err)
					}
					f.SetUint(i)
					return nil

				case reflect.Float32, reflect.Float64:
					i, err := strconv.ParseFloat(strValue, 64)
					if err != nil {
						return fmt.Errorf("解析浮点数失败 (值: %q): %w", strValue, err)
					}
					f.SetFloat(i)
					return nil

				case reflect.Bool:
					b, err := strconv.ParseBool(strValue)
					if err != nil {
						return fmt.Errorf("解析布尔值失败 (值: %q): %w", strValue, err)
					}
					f.SetBool(b)
					return nil

				default:
					return fmt.Errorf("不支持的类型转换: 无法将 []byte 转换为 %s", c.kind)
				}
			}
		}(field, rb, f))
	}
	return
}

// sql返回 row 字段下标 对应的  struct 字段名（""表示不接收该列数据）
type ColIndex2FieldNameMap []string
