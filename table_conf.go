package ldb

import (
	"reflect"
)

type LdbTabler interface {
	TableConf() *TableConf
}

type TableConf struct {
	tableName                *string  // 表名
	primaryKeyColumnNames    []string // 主键字段名列表
	autoPrimaryKeyColumnName *string  // 自增主键字段名
	otherAutoColumnName      []string // 其他自动生成字段名列表
	indexs                   []Index  // 数据库索引列表
}
type Index struct {
	Name      string   // 索引名称
	Unique    bool     // 是否唯一
	Columns   []string // 索引列
	IndexType string   // 索引类型
	Comment   string   // 索引注释
}

// api
func (c *TableConf) Table(name string) *TableConf {
	c.tableName = &name
	return c
}

// api
func (c *TableConf) PrimaryKeys(name ...string) *TableConf {
	c.primaryKeyColumnNames = name
	return c
}

// api
func (c *TableConf) AutoPrimaryKey(name string) *TableConf {
	c.autoPrimaryKeyColumnName = &name
	return c
}

// api
// OtherAutoField 其他自动生成字段
// 例如：
// 自增字段、虚拟列、计算列、默认值，等
// 在insert时，可以设置返回这些字段
func (c *TableConf) OtherAutoColumn(name ...string) *TableConf {
	c.otherAutoColumnName = name
	return c
}

var TableConfCache = map[reflect.Type]TableConf{}

func getTableConf(v reflect.Value) *TableConf {
	n, has := TableConfCache[v.Type()]
	if has {
		return &n
	}
	method := v.MethodByName("TableConf")
	if !method.IsValid() || method.IsZero() {
		return nil
	}

	values := method.Call(nil)

	if len(values) != 1 {
		return nil
	}
	value := values[0]
	if value.IsNil() {
		return nil
	}
	tc, ok := value.Interface().(*TableConf)
	if !ok {
		return nil
	}
	return tc
}

func getTableName(v reflect.Value) *string {
	tc := getTableConf(v)
	if tc == nil {
		return nil
	}
	return tc.tableName
}

func getPrimaryKeyColumnNames(v reflect.Value) []string {
	tc := getTableConf(v)
	if tc == nil {
		return nil
	}
	return tc.primaryKeyColumnNames
}
