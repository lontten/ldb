package utils

import "reflect"

func ToSliceValue(v reflect.Value) []reflect.Value {
	l := v.Len()
	arr := make([]reflect.Value, l)
	for i := 0; i < l; i++ {
		arr[i] = v.Index(i)
	}
	return arr
}

func ToSlice(v reflect.Value) []any {
	l := v.Len()
	arr := make([]any, l)
	for i := 0; i < l; i++ {
		arr[i] = v.Index(i).Interface()
	}
	return arr
}

func Contains(list []string, s string) bool {
	for _, a := range list {
		if a == s {
			return true
		}
	}
	return false
}

func Find(list []string, s string) int {
	for i, a := range list {
		if a == s {
			return i
		}
	}
	return -1
}

// HasDuplicate 检测字符串切片中是否存在重复元素
// 返回：bool - 有重复返回true，无重复返回false
func HasDuplicate(list []string) (bool, string) {
	// 定义map记录已出现的字符串，key为字符串，value用空结构体节省内存
	seen := make(map[string]struct{}, len(list))
	// 遍历切片中的每个字符串
	for _, s := range list {
		// 检查当前字符串是否已在map中
		if _, exists := seen[s]; exists {
			return true, s // 存在重复，直接返回true
		}
		seen[s] = struct{}{}
	}
	// 遍历完成无重复，返回false
	return false, ""
}
