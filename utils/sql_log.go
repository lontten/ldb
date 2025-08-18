package utils

import (
	"database/sql/driver"
	"fmt"
)

func PrintSql(sql string, args ...any) {
	fmt.Println("SQL:")
	fmt.Println(sql)
	fmt.Print("参数:")
	for _, arg := range args {
		fmt.Print(DbArg2String(arg), " , ")
	}
}

func DbArg2String(arg any) string {
	if v, ok := arg.(driver.Valuer); ok {
		value, err := v.Value()
		if err != nil {
			return fmt.Sprintf("获取 Value() 出错: %#v", err)
		} else {
			return fmt.Sprintf("%v", value)
		}
	} else {
		return fmt.Sprintf("%v", arg)
	}
}
