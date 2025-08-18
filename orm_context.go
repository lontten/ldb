package ldb

import (
	"github.com/lontten/ldb/field"
	"github.com/lontten/ldb/insert-type"
	"github.com/lontten/ldb/return-type"
	"github.com/lontten/ldb/softdelete"
	"github.com/lontten/ldb/sqltype"
	"github.com/lontten/ldb/utils"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type tableSqlType int

// d前缀是单表的意思，tableSqlType 只用于单表操作
const (
	dInsert tableSqlType = iota
	dInsertOrUpdate
	dUpdate
	dDelete
	dSelect
	dGetOrInsert
	dHas
	dCount
)

type returnAutoPrimaryKeyType int

// pk前缀是主键的意思
const (
	pkNoReturn    returnAutoPrimaryKeyType = iota
	pkQueryReturn                          // insert 时，可以直接 query 返回
	pkFetchReturn                          // insert 时，不能直接返回，需要手动获取
)

type ormContext struct {
	ormConf *OrmConf
	extra   *ExtraContext

	// model 参数，用于校验字段类型是否合法
	paramModelBaseV reflect.Value

	// dest
	scanDest  any
	scanIsPtr bool
	scanV     reflect.Value // dest 去除ptr的value

	destBaseValue      reflect.Value // list第一个，去除ptr
	destBaseType       reflect.Type
	destBaseTypeIsComp bool
	// scan 为slice时，里面item是否是ptr
	destIsSlice        bool
	destSliceItemIsPtr bool

	log Logger
	err error

	tableSqlType tableSqlType //单表，sql类型crud

	isLgDel bool //是否启用了逻辑删除
	isTen   bool //是否启用了多租户

	// ------------------主键----------------------
	indexs []Index // 索引列表

	insertCanReturn      bool                     // 数据库是否支持 insert时直接返回字段
	returnAutoPrimaryKey returnAutoPrimaryKeyType // 自增主键返回类型

	// 在不支持 insertCanReturn 的数据库中，使用 LastInsertId 返回 自增主键
	autoPrimaryKeyFieldName     *string      // 自增主键字段名
	autoPrimaryKeyFieldIsPtr    bool         // id 对应的model字段 是否是 ptr
	autoPrimaryKeyFieldBaseType reflect.Type // id 对应的model字段 type

	// 只能在 insert时，返回字段，只能支持 insertCanReturn 的数据库，可以返回
	otherAutoFieldNames []string // 其他自动生成字段名列表

	allAutoFieldNames []string // 所有自动生成字段名列表

	// id = 1
	//主键名-列表,这里考虑到多主键
	primaryKeyNames []string
	//主键值-列表
	primaryKeyValues [][]field.Value

	// id != 1 ,使用场景 更新名字时，检查名字重复，排除自己
	//主键名-列表,这里考虑到多主键-排除
	filterPrimaryKeyNames []string
	//主键值-列表-排除
	filterPrimaryKeyValues [][]field.Value

	// ------------------conf----------------------

	insertType     insert_type.InsertType
	returnType     return_type.ReturnType
	softDeleteType softdelete.SoftDelType
	skipSoftDelete bool   // 跳过软删除
	tableName      string //当前表名
	checkParam     bool   // 是否检查参数
	showSql        bool   // 打印sql
	noRun          bool   // 不实际执行

	// ------------------conf-end----------------------

	// ------------------字段名：字段值----------------------

	columns      []string      // 有效字段列表
	columnValues []field.Value // 有效字段- 值

	modelZeroFieldNames      []string       // model 零值字段列表
	modelNoSoftDelFieldNames []string       // model 所有字段列表- 忽略软删除字段
	modelAllFieldNames       []string       // model 所有字段列表
	modelFieldIndexMap       map[string]int // model字段名-index
	modelSelectFieldNames    []string       // model select 字段列表
	// ------------------字段名：字段值-end----------------------

	//------------------scan----------------------
	//true query,false exec
	sqlType sqltype.SqlType

	//要执行的sql语句
	query *strings.Builder
	//参数
	args []any

	started bool

	wb       *WhereBuilder
	whereSql string // WhereBuilder 生成的 where sql
	lastSql  string // 最后拼接的sql
	limit    *int64
	offset   *int64
}

func (ctx *ormContext) setLastInsertId(lastInsertId int64) {
	var vp reflect.Value
	switch ctx.autoPrimaryKeyFieldBaseType.Kind() {
	case reflect.Int8:
		id := int8(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Int16:
		id := int16(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Int32:
		id := int32(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Int64:
		id := int64(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Int:
		id := int(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Uint:
		id := uint(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Uint8:
		id := uint8(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Uint16:
		id := uint16(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Uint32:
		id := uint32(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	case reflect.Uint64:
		id := uint64(lastInsertId)
		vp = reflect.ValueOf(&id)
		break
	default:
		ctx.err = errors.New("last_insert_id field type error")
		return
	}
	f := ctx.destBaseValue.FieldByName(*ctx.autoPrimaryKeyFieldName)
	if ctx.autoPrimaryKeyFieldIsPtr {
		f.Set(vp)
	} else {
		f.Set(reflect.Indirect(vp))
	}
}
func (ctx *ormContext) initExtra(extra ...*ExtraContext) {
	var e *ExtraContext
	if len(extra) > 0 && extra[0] != nil {
		e = extra[0]
	} else {
		e = E()
	}
	// err 上抛到 ormContext
	if e.GetErr() != nil {
		ctx.err = e.GetErr()
		return
	}
	ctx.extra = e
	ctx.insertType = e.insertType
	ctx.returnType = e.returnType
	ctx.showSql = e.showSql
	ctx.noRun = e.noRun
	ctx.skipSoftDelete = e.skipSoftDelete
	ctx.tableName = e.tableName

	ctx.modelSelectFieldNames = e.selectColumns

	if len(e.orderByTokens) > 0 {
		ctx.lastSql += " ORDER BY " + strings.Join(e.orderByTokens, ",")
	}
	if ctx.limit == nil {
		ctx.limit = e.limit
	}
	ctx.offset = e.offset
}

// 初始化 表名,主键，自增id
func (ctx *ormContext) initConf() {
	if ctx.hasErr() {
		return
	}

	v := ctx.destBaseValue

	dest := ctx.scanDest
	t := ctx.destBaseType
	ctx.softDeleteType = utils.GetSoftDelType(t)

	if ctx.tableName == "" {
		tableName := ctx.ormConf.tableName(v, dest)
		ctx.tableName = tableName
	}

	primaryKeys := ctx.ormConf.primaryKeys(v, dest)
	ctx.primaryKeyNames = primaryKeys

	tc := getTableConf(v)
	if tc != nil {
		ctx.autoPrimaryKeyFieldName = tc.autoPrimaryKeyColumnName
		ctx.otherAutoFieldNames = tc.otherAutoColumnName
	}
}

// 获取struct对应的字段名 和 其值，
// slice为全部，一个为非nil字段。
func (ctx *ormContext) initColumnsValue() {
	if ctx.hasErr() {
		return
	}

	cv := getStructCVMap(ctx.destBaseValue)
	ctx.columns = cv.columns
	ctx.columnValues = cv.columnValues

	ctx.modelZeroFieldNames = cv.modelZeroFieldNames
	ctx.modelNoSoftDelFieldNames = cv.modelAllFieldNames
	ctx.modelAllFieldNames = cv.modelAllFieldNames

	if ctx.scanIsPtr && ctx.returnType != return_type.None {
		if ctx.insertCanReturn {
			ctx.returnAutoPrimaryKey = pkQueryReturn
		} else if ctx.autoPrimaryKeyFieldName != nil {
			ctx.returnAutoPrimaryKey = pkFetchReturn
		}
	}

	if ctx.returnAutoPrimaryKey == pkFetchReturn {
		fieldName, ok := cv.modelAllFieldNameMap[*ctx.autoPrimaryKeyFieldName]
		if !ok {
			ctx.err = errors.New("auto_increment field not found")
			return
		}
		ctx.autoPrimaryKeyFieldName = &fieldName

		structField, _ := ctx.destBaseType.FieldByName(fieldName)
		isPtr, baseT := basePtrType(structField.Type)
		ctx.autoPrimaryKeyFieldIsPtr = isPtr
		ctx.autoPrimaryKeyFieldBaseType = baseT
	}
	return
}
func (ctx *ormContext) initColumns() {
	if ctx.hasErr() {
		return
	}
	if len(ctx.modelSelectFieldNames) == 0 {
		columns := getStructCList(ctx.destBaseType)
		ctx.modelSelectFieldNames = columns
	}
	return
}

func (ctx *ormContext) initColumnsValueExtra() {
	if ctx.hasErr() {
		return
	}
	e := ctx.extra
	set := e.set
	if set.hasModel {
		oc := &ormContext{
			ormConf:        ctx.ormConf,
			skipSoftDelete: true,
		}
		oc.initTargetDestOne(set.model) //初始化参数
		oc.initColumnsValue()           //初始化cv

		set.columns = append(set.columns, oc.columns...)
		set.columnValues = append(set.columnValues, oc.columnValues...)
	}
	if ctx.hasErr() {
		return
	}
	for i, column := range e.columns {
		cv := e.columnValues[i]
		if cv.Type == field.Null || cv.Type == field.Now {
			ctx.modelZeroFieldNames = append(ctx.modelZeroFieldNames, column)
		}
		find := utils.Find(ctx.columns, column)
		if find == -1 {
			ctx.columns = append(ctx.columns, column)
			ctx.columnValues = append(ctx.columnValues, cv)
		} else {
			ctx.columnValues[find] = cv
		}
	}
	return
}
func (ctx *ormContext) initColumnsValueSoftDel() {
	if ctx.hasErr() {
		return
	}
	if ctx.skipSoftDelete {
		return
	}

	switch ctx.sqlType {
	case sqltype.Insert:
		value, has := softdelete.SoftDelTypeNoFVMap[ctx.softDeleteType]
		if has {
			ctx.columns = append(ctx.columns, value.Name)
			ctx.columnValues = append(ctx.columnValues, value.ToValue())
		}
		break
	case sqltype.Delete:
		if ctx.softDeleteType != softdelete.None {
			// set
			value, has := softdelete.SoftDelTypeYesFVMap[ctx.softDeleteType]
			if has {
				ctx.columns = append(ctx.columns, value.Name)
				ctx.columnValues = append(ctx.columnValues, value.ToValue())
			}

			// where
			value, has = softdelete.SoftDelTypeNoFVMap[ctx.softDeleteType]
			if has {
				ctx.wb.fieldValue(value.Name, value.ToValue())
			}
		}
		break
	case sqltype.Update:
		if ctx.softDeleteType != softdelete.None {
			// where
			value, has := softdelete.SoftDelTypeNoFVMap[ctx.softDeleteType]
			if has {
				ctx.wb.fieldValue(value.Name, value.ToValue())
			}
		}
		break
	case sqltype.Select:
		if ctx.softDeleteType != softdelete.None {
			// select
			value, has := softdelete.SoftDelTypeYesFVMap[ctx.softDeleteType]
			if has {
				ctx.modelSelectFieldNames = append(ctx.modelSelectFieldNames, value.Name)
			}

			// where
			value, has = softdelete.SoftDelTypeNoFVMap[ctx.softDeleteType]
			if has {
				ctx.wb.fieldValue(value.Name, value.ToValue())
			}
		}
		break

	default:
		break
	}
	return
}

func (ctx ormContext) Copy() ormContext {
	return ormContext{
		ormConf: ctx.ormConf,
		log:     ctx.log,
	}
}

// select 生成
func (ctx *ormContext) selectArgsArr2SqlStr(args []string) {
	query := ctx.query
	if ctx.started {
		for _, name := range args {
			query.WriteString(", " + name)
		}
	} else {
		query.WriteString("SELECT ")
		for i := range args {
			if i == 0 {
				query.WriteString(args[i])
			} else {
				query.WriteString(", " + args[i])
			}
		}
		if len(args) > 0 {
			ctx.started = true
		}
	}
}

// create 生成
func (ctx *ormContext) tableInsertGen() string {
	args := ctx.columns
	var sb strings.Builder

	sb.WriteString("INSERT INTO ")
	sb.WriteString(ctx.tableName + " ")

	sb.WriteString(" ( ")
	for i, v := range args {
		if i == 0 {
			sb.WriteString(v)
		} else {
			sb.WriteString(" , " + v)
		}
	}
	sb.WriteString(" ) ")
	sb.WriteString(" VALUES ")
	sb.WriteString("( ")
	for i := range args {
		if i == 0 {
			sb.WriteString(" ? ")
		} else {
			sb.WriteString(", ? ")
		}
	}
	sb.WriteString(" ) ")
	return sb.String()
}

// 单表sql生成，insert
func (p *PgDialect) tGenInsert() string {
	args := p.ctx.columns
	var sb strings.Builder

	sb.WriteString("INSERT INTO ")
	sb.WriteString(p.ctx.tableName + " ")

	sb.WriteString(" ( ")
	for i, v := range args {
		if i == 0 {
			sb.WriteString(v)
		} else {
			sb.WriteString(" , " + v)
		}
	}
	sb.WriteString(" ) ")
	sb.WriteString(" VALUES ")
	sb.WriteString("( ")
	for i := range args {
		if i == 0 {
			sb.WriteString(" ? ")
		} else {
			sb.WriteString(", ? ")
		}
	}
	sb.WriteString(" ) ")
	return sb.String()
}

func (ctx *ormContext) createSqlGenera(args []string) string {
	var sb strings.Builder
	sb.WriteString(" ( ")
	for i, v := range args {
		if i == 0 {
			sb.WriteString(v)
		} else {
			sb.WriteString(" , " + v)
		}
	}
	sb.WriteString(" ) ")
	sb.WriteString(" VALUES ")
	sb.WriteString("( ")
	for i := range args {
		if i == 0 {
			sb.WriteString(" ? ")
		} else {
			sb.WriteString(", ? ")
		}
	}
	sb.WriteString(" ) ")
	return sb.String()
}

// upd 生成
func (ctx *ormContext) tableUpdateArgs2SqlStr(args []string) string {
	var sb strings.Builder
	l := len(args)
	for i, v := range args {
		if i != l-1 {
			sb.WriteString(v + " = ? ,")
		} else {
			sb.WriteString(v + " = ? ")
		}
	}
	return sb.String()
}

func (ctx *ormContext) initPrimaryKeyByWhere(wb *WhereBuilder) {
	ctx.primaryKeyValues = ctx.initPrimaryKeyValues(wb.primaryKeyValue)
	if ctx.hasErr() {
		return
	}
	builderAnd := W()
	for _, value := range ctx.primaryKeyValues {
		builder := W()
		for i, name := range ctx.primaryKeyNames {
			builder.Eq(name, value[i].Value)
		}
		builderAnd.Or(builder)
	}
	wb.And(builderAnd)
	ctx.filterPrimaryKeyValues = ctx.initPrimaryKeyValues(wb.filterPrimaryKeyValue)
	if ctx.hasErr() {
		return
	}
	builderAnd = W()
	for _, value := range ctx.filterPrimaryKeyValues {
		builder := W()
		for i, name := range ctx.primaryKeyNames {
			builder.Neq(name, value[i].Value)
		}
		builderAnd.Or(builder)
	}
	wb.And(builderAnd)
}
func (ctx *ormContext) initPrimaryKeyValues(v []any) (idValuess [][]field.Value) {
	if ctx.hasErr() {
		return
	}
	if len(v) == 0 {
		return
	}

	idLen := len(v)
	if idLen == 0 {
		ctx.err = errors.New("ByPrimaryKey arg len num 0")
		return
	}
	pkLen := len(ctx.primaryKeyNames)

	if pkLen == 1 { //单主键
		for _, i := range v {
			value := reflect.ValueOf(i)
			_, value, err := basePtrDeepValue(value)
			if err != nil {
				ctx.err = err
				return
			}

			if ctx.checkParam {
				if !isValuerType(value.Type()) {
					ctx.err = errors.New("ByPrimaryKey typ err,not single")
					return
				}
			}

			idValues := make([]field.Value, 1)
			idValues[0] = field.Value{
				Type:  field.Val,
				Value: value.Interface(),
			}
			idValuess = append(idValuess, idValues)
		}

	} else {
		for _, i := range v {
			value := reflect.ValueOf(i)
			_, value, err := basePtrDeepValue(value)
			if err != nil {
				ctx.err = err
				return
			}

			if ctx.checkParam {
				if !isCompType(value.Type()) {
					ctx.err = errors.New("ByPrimaryKey typ err,not comp")
					return
				}
			}

			columns, values, err := getCompValueCV(value)
			if err != nil {
				ctx.err = err
				return
			}

			if ctx.checkParam {
				if len(columns) != pkLen {
					ctx.err = errors.New("复合主键，filed数量 len err")
					return
				}
			}

			idValues := make([]field.Value, 0)
			idValues = append(idValues, values...)
			idValuess = append(idValuess, idValues)
		}
	}
	return
}

func (ctx *ormContext) hasErr() bool {
	return ctx.err != nil
}
