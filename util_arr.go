package ldb

import (
	"reflect"
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
