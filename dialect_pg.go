package ldb

import (
	"errors"
	"github.com/lontten/ldb/insert-type"
	"github.com/lontten/ldb/return-type"
	"github.com/lontten/ldb/utils"
	"strconv"
	"strings"
)

type PgDialect struct {
	ctx *ormContext
}

// ===----------------------------------------------------------------------===//
// 获取上下文
// ===----------------------------------------------------------------------===//
func (d *PgDialect) getCtx() *ormContext {
	return d.ctx
}
func (d *PgDialect) initContext() Dialecter {
	return &PgDialect{
		ctx: &ormContext{
			ormConf:    d.ctx.ormConf,
			query:      &strings.Builder{},
			wb:         W(),
			insertType: insert_type.Err,
		},
	}
}
func (d *PgDialect) hasErr() bool {
	return d.ctx.err != nil
}
func (d *PgDialect) getErr() error {
	return d.ctx.err
}

// ===----------------------------------------------------------------------===//
// sql 方言化
// ===----------------------------------------------------------------------===//

func (d *PgDialect) prepare(query string) string {
	query = toPgSql(query)
	return query
}
func (d *PgDialect) exec(query string, args ...any) (string, []any) {
	query = toPgSql(query)
	return query, args
}

func (d *PgDialect) query(query string, args ...any) (string, []any) {
	query = toPgSql(query)
	return query, args
}

func (d *PgDialect) queryBatch(query string) string {
	query = toPgSql(query)

	//return m.ldb.Prepare(query)
	return query
}

func (d *PgDialect) getSql() string {
	s := d.ctx.query.String()
	return toPgSql(s)
}

// insert 生成
func (d *PgDialect) tableInsertGen() {
	ctx := d.ctx
	if ctx.hasErr() {
		return
	}
	if ctx.insertType == insert_type.Replace {
		ctx.err = errors.New("pg不支持的插入类型 insert-type.Replace")
		return
	}
	extra := ctx.extra
	set := extra.set

	columns := ctx.columns
	var query = d.ctx.query

	query.WriteString("INSERT INTO ")

	query.WriteString(ctx.tableName + " ")

	query.WriteString("(")
	query.WriteString(strings.Join(columns, ","))
	query.WriteString(") ")
	query.WriteString("VALUES")
	query.WriteString("(")
	ctx.genInsertValuesSqlBycolumnValues()
	query.WriteString(") ")

	if ctx.insertType == insert_type.Ignore || ctx.insertType == insert_type.Update {
		query.WriteString("ON CONFLICT (")
		query.WriteString(strings.Join(extra.duplicateKeyNames, ","))
		query.WriteString(") DO ")
	}

	switch ctx.insertType {
	case insert_type.Ignore:
		query.WriteString("NOTHING ")
		break
	case insert_type.Update:
		query.WriteString("UPDATE SET ")

		// 当未设置更新字段时，默认为所有字段
		if len(set.columns) == 0 && len(set.fieldNames) == 0 {
			list := append(ctx.columns, extra.columns...)

			for _, name := range list {
				find := utils.Find(extra.duplicateKeyNames, name)
				if find < 0 { // 排除 主键 字段
					set.fieldNames = append(set.fieldNames, name)
				}
			}
		}

		for i, name := range set.fieldNames {
			query.WriteString(name + "= EXCLUDED." + name)
			if i < len(set.fieldNames)-1 {
				query.WriteString(",")
			}
		}

		for i, column := range set.columns {
			if i > 0 {
				query.WriteString(", ")
			}
			query.WriteString(column + "= ? ")
			ctx.args = append(ctx.args, set.columnValues[i].Value)
		}
		break
	default:
		break
	}

	// 当scan为指针类型时，返回字段。
	if ctx.returnAutoPrimaryKey != pkNoReturn {
		switch expr := ctx.returnType; expr {
		case return_type.None:
			break
		case return_type.Auto:
			var list []string
			for _, s := range ctx.otherAutoFieldNames {
				list = append(list, s)
			}
			list = append(list, *ctx.autoPrimaryKeyFieldName)
			query.WriteString(" RETURNING " + strings.Join(list, ","))
		case return_type.ZeroField:
			query.WriteString(" RETURNING " + strings.Join(ctx.modelZeroFieldNames, ","))
		case return_type.AllField:
			query.WriteString(" RETURNING " + strings.Join(ctx.modelAllFieldNames, ","))
		}
	}
	query.WriteString(";")
}

// del 生成
func (d *PgDialect) tableDelGen() {

}

// update 生成
func (d *PgDialect) tableUpdateGen() {

}

// select 生成
func (d *PgDialect) tableSelectGen() {

}

func (d *PgDialect) execBatch(query string, args [][]any) (string, [][]any) {
	query = toPgSql(query)
	//var num int64 = 0
	//stmt, err := m.ldb.Prepare(query)
	//defer stmt.Close()
	//if err != nil {
	//	return 0, err
	//}
	//for _, arg := range args {
	//	exec, err := stmt.Exec(arg...)
	//
	//	m.log.Println(query, arg)
	//	if err != nil {
	//		return num, err
	//	}
	//	rowsAffected, err := exec.RowsAffected()
	//	if err != nil {
	//		return num, err
	//	}
	//	num += rowsAffected
	//}
	return query, args
}

// ===----------------------------------------------------------------------===//
// 工具
// ===----------------------------------------------------------------------===//

func toPgSql(sql string) string {
	var i = 1
	for {
		t := strings.Replace(sql, "?", "$"+strconv.Itoa(i), 1)
		if t == sql {
			break
		}
		i++
		sql = t
	}
	return sql
}

// ===----------------------------------------------------------------------===//
// 中间服务
// ===----------------------------------------------------------------------===//
func (d *PgDialect) toSqlInsert() (string, []any) {
	tableName := d.ctx.tableName
	return tableName, nil
}

func (d *PgDialect) parse(c Clause) (string, error) {
	sb := strings.Builder{}
	switch c.Type {
	case Eq:
		sb.WriteString(c.query + " = ?")
	case Neq:
		sb.WriteString(c.query + " <> ?")
	case Less:
		sb.WriteString(c.query + " < ?")
	case LessEq:
		sb.WriteString(c.query + " <= ?")
	case Greater:
		sb.WriteString(c.query + " > ?")
	case GreaterEq:
		sb.WriteString(c.query + " >= ?")
	case Like:
		sb.WriteString(c.query + " LIKE ?")
	case NotLike:
		sb.WriteString(c.query + " NOT LIKE ?")
	case In:
		sb.WriteString(c.query + " IN (")
		sb.WriteString(gen(c.argsNum))
		sb.WriteString(")")
	case NotIn:
		sb.WriteString(c.query + " NOT IN (")
		sb.WriteString(gen(c.argsNum))
		sb.WriteString(")")
	case Between:
		sb.WriteString(c.query + " BETWEEN ? AND ?")
	case NotBetween:
		sb.WriteString(c.query + " NOT BETWEEN ? AND ?")
	case IsNull:
		sb.WriteString(c.query + " IS NULL")
	case IsNotNull:
		sb.WriteString(c.query + " IS NOT NULL")
	case IsFalse:
		sb.WriteString(c.query + " IS FALSE")
	default:
		return "", errors.New("unknown where token type")
	}

	return sb.String(), nil
}
