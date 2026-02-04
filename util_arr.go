package ldb

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lontten/ldb/v2/utils"
)

// processArrArg 通用数组参数处理函数
// 输入: arg any - 可以是切片、数组、单个值
func processArrArg(arg any) (bool, []any) {
	val := reflect.ValueOf(arg)
	_, val, _ = basePtrValue(val)
	kind := val.Kind()

	switch kind {
	case reflect.Slice, reflect.Array:
		length := val.Len()

		// 处理空切片/数组
		if length == 0 {
			return true, []any{}
		}

		// 预构建占位符字符串
		args := make([]any, 0, length)

		for i := 0; i < length; i++ {
			args = append(args, val.Index(i).Interface())
		}
		return true, args
	default:
		return false, []any{arg}
	}
}

// 处理query、args，
// 扩展 in 参数
func processNativeExIn(query string, args ...any) (string, []any, error) {
	if len(args) == 0 {
		return query, args, nil
	}

	num, matches := SplitQueryByArgs(query)

	if num != len(args) {
		return "", []any{}, ErrArgsLen
	}

	var newArgs = make([]any, 0, len(args))
	var result strings.Builder

	var index = 0
	for _, match := range matches {
		if match == "?" {
			var arg = args[index]
			index++
			result.WriteString(match)
			newArgs = append(newArgs, arg)
		} else if match == "(?)" {
			var arg = args[index]
			isNil := utils.IsNil(arg)
			if isNil {
				return "", []any{}, fmt.Errorf("invalid use of Native: argument for field '%s' is nil", query)
			}

			index++
			_, argsE := processArrArg(arg)
			length := len(argsE)

			if length == 0 {
				result.WriteString(match)
				newArgs = append(newArgs, nil)
			} else {
				result.WriteString("(" + gen(length) + ")")
				newArgs = append(newArgs, argsE...)
			}
		} else {
			result.WriteString(match)
		}
	}
	return result.String(), newArgs, nil
}
