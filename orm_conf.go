package ldb

import (
	"reflect"
	"strings"
	"sync"

	"github.com/lontten/ldb/utils"
)

type OrmConf struct {
	//po生成文件目录
	PoDir string
	//是否覆盖，默认true
	IsFileOverride bool

	//作者
	Author string
	//是否开启ActiveRecord模式,默认false
	IsActiveRecord bool

	IdType int

	//表名
	//TableNameFun >  tag > TableNamePrefix
	TableNamePrefix string
	TableNameFun    func(t reflect.Value, dest any) string

	//主键 默认为id
	PrimaryKeyNames   []string
	PrimaryKeyNameFun func(v reflect.Value, dest any) []string

	//多租户
	TenantIdFieldName    string                      //多租户的  租户字段名 空字符串极为不启用多租户
	TenantIdValueFun     func() any                  //租户的id值，获取函数
	TenantIgnoreTableFun func(tableName string) bool //该表是否忽略多租户，true忽略该表，即没有多租户
}

var typeTableNameCache = map[reflect.Type]string{}
var typeTableNameMu sync.Mutex

func getTypeTableName(t reflect.Type, tableNamePrefix string) string {
	s, ok := typeTableNameCache[t]
	if ok {
		return s
	}
	typeTableNameMu.Lock()
	defer typeTableNameMu.Unlock()
	s, ok = typeTableNameCache[t]
	if ok {
		return s
	}

	name := t.String()
	index := strings.LastIndex(name, ".")
	if index > 0 {
		name = name[index+1:]
	}
	name = utils.Camel2Case(name)
	if tableNamePrefix != "" {
		name = tableNamePrefix + name
	}
	typeTableNameCache[t] = name
	return name
}

// 不可缓存
// 获取表名
func (c OrmConf) tableName(v reflect.Value, dest any) string {
	// fun
	tableNameFun := c.TableNameFun
	if tableNameFun != nil {
		return tableNameFun(v, dest)
	}

	// tableName
	n := getTableName(v)
	if n != nil {
		return *n
	}

	// structName
	t := v.Type()
	name := getTypeTableName(t, c.TableNamePrefix)
	return name
}

// 不可缓存
// 1.可以PrimaryKeyNames设置主键字段名
// 2.通过表名动态设置主键字段名-fn
func (c OrmConf) primaryKeyColumnNames(v reflect.Value, dest any) []string {
	//fun
	primaryKeyColumnNameFun := c.PrimaryKeyNameFun
	if primaryKeyColumnNameFun != nil {
		return primaryKeyColumnNameFun(v, dest)
	}

	return getPrimaryKeyColumnNames(v)
}

// 获取 rows 返回数据，每个字段index 对应 struct 的字段 名字
func getColIndex2FieldNameMap(columns []string, t reflect.Type) (ColIndex2FieldNameMap, error) {
	if isValuerType(t) {
		return ColIndex2FieldNameMap{}, nil
	}

	colNum := len(columns)
	ci2fm := make([]string, colNum)
	cf := getStructCFMap(t)

	validNum := 0
	for i, column := range columns {
		fieldName, ok := cf[column]
		if !ok {
			ci2fm[i] = ""
			continue
		}
		ci2fm[i] = fieldName
		validNum++
	}

	if colNum == 1 && validNum == 0 {
		return ColIndex2FieldNameMap{}, nil
	}
	return ci2fm, nil
}
