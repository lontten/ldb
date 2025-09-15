package ldb

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/lontten/lcore/v2/lcutils"
	"github.com/lontten/lcore/v2/types"
)

func QueryBuild(db Engine) *SqlBuilder {
	return &SqlBuilder{
		db:              db.init(),
		selectQuery:     &strings.Builder{},
		otherSqlBuilder: &strings.Builder{},
	}
}

type SqlBuilder struct {
	db Engine
	// 最终执行sql
	query string
	// 最终参数列表
	args []any

	// select部分
	selectStatus int8

	selectTokens []string
	orderTokens  []string
	selectQuery  *strings.Builder
	selectArgs   []any

	// 分页
	countField   string // 分页时，用于查询总数的字段，默认为 *
	fakeTotalNum int64  // 分页-假数据总数，分页时，跳过查询，直接使用这个总数;默认-1，表示未设置
	noGetList    bool   // 分页-不返回数据，分页时，只查询数量，不返回数据列表;默认false，表示需要获取数据

	// 其他部分
	otherSqlBuilder *strings.Builder
	otherSqlArgs    []any
	whereStatus     int8

	// page
	other any
}

const (
	selectNoSet = iota
	selectSet

	selectDone
)

const (
	whereNoSet = iota
	whereSet
	whereDone
)

func (b *SqlBuilder) initSelectSql() {
	b.selectQuery.WriteString("SELECT ")
	b.selectQuery.WriteString(strings.Join(b.selectTokens, ","))
	b.query = b.selectQuery.String() + " " + b.otherSqlBuilder.String()
	if len(b.orderTokens) > 0 {
		b.query = b.query + " ORDER BY " + strings.Join(b.orderTokens, ",")
	}
	b.args = append(b.selectArgs, b.otherSqlArgs...)
}

// 显示sql
func (b *SqlBuilder) ShowSql(conditions ...bool) *SqlBuilder {
	for _, c := range conditions {
		if !c {
			return b
		}
	}
	b.db.getCtx().showSql = true
	return b
}

// 不执行
func (b *SqlBuilder) NoRun(conditions ...bool) *SqlBuilder {
	for _, c := range conditions {
		if !c {
			return b
		}
	}
	b.db.getCtx().noRun = true
	return b
}

// 添加一个 arg，多个断言
func (b *SqlBuilder) AppendArg(arg any, conditions ...bool) *SqlBuilder {
	if b.db.getCtx().hasErr() {
		return b
	}
	for _, c := range conditions {
		if !c {
			return b
		}
	}
	if b.selectStatus == selectNoSet {
		b.selectArgs = append(b.selectArgs, arg)
	} else {
		b.otherSqlArgs = append(b.otherSqlArgs, arg)
	}
	return b
}

// 添加sql语句
func (b *SqlBuilder) AppendSql(sql string) *SqlBuilder {
	b.otherSqlBuilder.WriteString(sql)
	return b
}

// 添加 多个参数
func (b *SqlBuilder) AppendArgs(args ...any) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	if b.selectStatus == selectDone {
		b.otherSqlArgs = append(b.otherSqlArgs, args...)
	} else {
		b.selectArgs = append(b.selectArgs, args...)
	}
	return b
}

// 添加一个 select 字段，多个断言
func (b *SqlBuilder) Select(arg string, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}

	b.selectStatus = selectSet
	b.selectTokens = append(b.selectTokens, arg)

	return b
}

// 添加 多个 select 字段，从 model中
func (b *SqlBuilder) SelectModel(v any) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	if v == nil {
		return b
	}

	ctx.initScanDestOne(v)
	columns := getStructCAllList(ctx.destBaseType)

	b.selectStatus = selectSet
	b.selectTokens = append(b.selectTokens, columns...)
	return b
}

// from 表名
// 状态从 selectNoSet 变成 selectSet
func (b *SqlBuilder) From(name string) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	b.selectStatus = selectDone
	b.otherSqlBuilder.WriteString(" FROM " + name)
	return b
}

// join 联表
func (b *SqlBuilder) Join(name string, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.otherSqlBuilder.WriteString(" JOIN " + name)
	return b
}

func (b *SqlBuilder) Arg(arg any, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.AppendArgs(arg)
	return b
}

func (b *SqlBuilder) Args(args ...any) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	b.AppendArgs(args...)
	return b
}

func (b *SqlBuilder) LeftJoin(name string, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.otherSqlBuilder.WriteString("\n")
	b.otherSqlBuilder.WriteString("LEFT JOIN " + name)
	b.otherSqlBuilder.WriteString("\n")

	return b
}

func (b *SqlBuilder) RightJoin(name string, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.otherSqlBuilder.WriteString(" RIGHT JOIN " + name)
	return b
}

func (b *SqlBuilder) OrderBy(name string, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.orderTokens = append(b.orderTokens, name+" ASC")
	return b
}

func (b *SqlBuilder) OrderDescBy(name string, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.orderTokens = append(b.orderTokens, name+" DESC")
	return b
}
func (b *SqlBuilder) Native(sql string, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.otherSqlBuilder.WriteString(" ")
	b.otherSqlBuilder.WriteString(sql)
	b.otherSqlBuilder.WriteString(" ")
	return b
}

func (b *SqlBuilder) Limit(num int64, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.otherSqlBuilder.WriteString(" LIMIT " + strconv.FormatInt(num, 10))
	return b
}

func (b *SqlBuilder) Offset(num int64, condition ...bool) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}

	for _, c := range condition {
		if !c {
			return b
		}
	}
	b.otherSqlBuilder.WriteString(" OFFSET " + strconv.FormatInt(num, 10))
	return b
}

func (b *SqlBuilder) WhereBuilder(w *WhereBuilder) *SqlBuilder {
	ctx := b.db.getCtx()
	if ctx.hasErr() {
		return b
	}
	if w == nil {
		return b
	}
	sqlStr, args, err := w.toSql(b.db.getDialect().parse)
	if err != nil {
		b.db.getCtx().err = err
		return b
	}
	if sqlStr == "" {
		return b
	}
	sqlStr = "(" + sqlStr + ")"

	if b.selectStatus != selectDone {
		ctx.err = errors.New("未完成 select 设置")
	}
	switch b.whereStatus {
	case whereNoSet:
		b.whereStatus = whereSet
		b.otherSqlBuilder.WriteString(" WHERE ")
		b.otherSqlBuilder.WriteString(sqlStr)
	case whereSet:
		b.otherSqlBuilder.WriteString(" AND ")
		b.otherSqlBuilder.WriteString(sqlStr)
	case whereDone:
		b.db.getCtx().err = errors.New("where has been done")
	}

	b.AppendArgs(args...)
	return b
}

func (b *SqlBuilder) LinkWhere() *SqlBuilder {
	b.selectStatus = selectDone
	b.whereStatus = whereSet
	return b
}

func (b *SqlBuilder) Where(whereStr string, condition ...bool) *SqlBuilder {
	for _, c := range condition {
		if !c {
			return b
		}
	}
	b._whereArg(whereStr)
	return b
}

func (b *SqlBuilder) _whereArg(whereStr string, args ...any) *SqlBuilder {
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return b
	}
	b.selectStatus = selectDone

	b.AppendArgs(args...)
	switch b.whereStatus {
	case whereNoSet:
		b.whereStatus = whereSet
		b.otherSqlBuilder.WriteString(" WHERE ")
		b.otherSqlBuilder.WriteString(whereStr)
	case whereSet:
		b.otherSqlBuilder.WriteString(" AND ")
		b.otherSqlBuilder.WriteString(whereStr)
	case whereDone:
		ctx.err = errors.New("where has been done")
	}

	return b
}
func (b *SqlBuilder) BoolWhere(condition bool, whereStr string, args ...any) *SqlBuilder {
	if !condition {
		return b
	}
	b._whereArg(whereStr, args...)
	return b
}

func (b *SqlBuilder) WhereIn(whereStr string, args ...any) *SqlBuilder {
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return b
	}

	if args == nil {
		return b
	}
	length := len(args)
	if length == 0 {
		return b
	}

	if b.selectStatus != selectDone {
		ctx.err = errors.New("Where 设置异常：" + whereStr)
		return b
	}

	b.AppendArgs(args...)

	var inArgStr = " (" + gen(length) + ")"
	whereStr = whereStr + " IN" + inArgStr

	switch b.whereStatus {
	case whereNoSet:
		b.whereStatus = whereSet
		b.otherSqlBuilder.WriteString(" WHERE ")

		b.otherSqlBuilder.WriteString(whereStr)

	case whereSet:
		b.otherSqlBuilder.WriteString(" AND ")

		b.otherSqlBuilder.WriteString(whereStr)

	case whereDone:
		ctx.err = errors.New("where has been done")
	}

	return b
}

// WhereSqlIn
// in ? 当参数列表长度为0时，跳过这个where
func (b *SqlBuilder) WhereSqlIn(whereStr string, args ...any) *SqlBuilder {
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return b
	}

	if args == nil {
		return b
	}
	length := len(args)
	if length == 0 {
		return b
	}

	if b.selectStatus != selectDone {
		ctx.err = errors.New("Where 设置异常：" + whereStr)
		return b
	}

	b.AppendArgs(args...)

	var inArgStr = " (" + gen(length) + ")"
	whereStr = strings.Replace(whereStr, "?", inArgStr, -1)

	switch b.whereStatus {
	case whereNoSet:
		b.whereStatus = whereSet
		b.otherSqlBuilder.WriteString(" WHERE ")

		b.otherSqlBuilder.WriteString(whereStr)

	case whereSet:
		b.otherSqlBuilder.WriteString(" AND ")

		b.otherSqlBuilder.WriteString(whereStr)

	case whereDone:
		ctx.err = errors.New("where has been done")
	}

	return b
}

func (b *SqlBuilder) Between(whereStr string, begin, end any, condition ...bool) *SqlBuilder {
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return b
	}
	if b.selectStatus != selectDone {
		ctx.err = errors.New("Where 设置异常：" + whereStr)
		return b
	}

	for _, c := range condition {
		if !c {
			return b
		}
	}
	// any类型，无法直接判断 nil
	begin = getFieldInterZero(reflect.ValueOf(begin))
	end = getFieldInterZero(reflect.ValueOf(end))

	if begin != nil {
		if end != nil {
			b._whereArg(whereStr+" BETWEEN ? AND ?", begin, end)
			return b
		}
		b._whereArg(whereStr+" >= ?", begin)
		return b
	}
	if end != nil {
		b._whereArg(whereStr+" <= ?", end)
		return b
	}
	return b
}

func (b *SqlBuilder) Like(key *string, fields ...string) *SqlBuilder {
	b._like(key, 1, fields...)
	return b
}
func (b *SqlBuilder) LikeLeft(key *string, fields ...string) *SqlBuilder {
	b._like(key, 2, fields...)
	return b
}
func (b *SqlBuilder) LikeRight(key *string, fields ...string) *SqlBuilder {
	b._like(key, 3, fields...)
	return b
}

// likeType
// 1 表示 like '%key%';
// 2 表示 like '%key';
// 3 表示 like 'key%';
func (b *SqlBuilder) _like(key *string, likeType int, fields ...string) *SqlBuilder {
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return b
	}
	if b.selectStatus != selectDone {
		ctx.err = errors.New("Where 设置异常：like ")
		return b
	}

	if lcutils.NilToZero(key) == "" {
		return b
	}
	if len(fields) == 0 {
		return b
	}
	var args []any
	var k = ""
	if likeType == 1 {
		k = "%" + *key + "%"
	} else if likeType == 2 {
		k = "%" + *key
	} else if likeType == 3 {
		k = *key + "%"
	}

	var tokens []string
	for _, field := range fields {
		tokens = append(tokens, field+" LIKE ? ")
		args = append(args, k)
	}
	var whereStr = "(" + strings.Join(tokens, " OR ") + ")"
	b._whereArg(whereStr, args...)
	return b
}

// BetweenDateTimeOfDate
// 用 Date类型，去查询 DateTime 字段
func (b *SqlBuilder) BetweenDateTimeOfDate(whereStr string, dateBegin, dateEnd *types.LocalDate, condition ...bool) *SqlBuilder {
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return b
	}
	if b.selectStatus != selectDone {
		ctx.err = errors.New("Where 设置异常：" + whereStr)
		return b
	}

	for _, c := range condition {
		if !c {
			return b
		}
	}

	var dateTimeBegin *types.LocalDateTime = nil
	if dateBegin != nil {
		dateTimeBegin = dateBegin.ToDateTimeP()
	}

	var dateTimeEnd *types.LocalDateTime = nil
	if dateEnd != nil {
		dateTimeEnd = dateEnd.Add(types.Duration().Day(1)).ToDateTimeP()
	}

	if dateTimeBegin != nil {
		b._whereArg(whereStr+" >= ?", dateTimeBegin)
	}
	if dateTimeEnd != nil {
		b._whereArg(whereStr+" < ?", dateTimeEnd)
	}

	return b
}

func (b *SqlBuilder) ScanOne(dest any) (rowsNum int64, err error) {
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return 0, ctx.err
	}
	b.selectStatus = selectDone
	b.whereStatus = whereDone

	ctx.initScanDestOne(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}

	b.initSelectSql()
	query := b.query
	args := b.args
	if ctx.showSql {
		fmt.Println(query, args)
	}
	if ctx.noRun {
		return 0, nil
	}

	rows, err := db.query(query, args...)
	if err != nil {
		return 0, err
	}
	return ctx.ScanLnT(rows)
}

func (b *SqlBuilder) ScanList(dest any) (rowsNum int64, err error) {
	db := b.db
	ctx := db.getCtx()
	if ctx.hasErr() {
		return 0, ctx.err
	}
	b.selectStatus = selectDone
	b.whereStatus = whereDone

	ctx.initScanDestList(dest)
	if ctx.err != nil {
		return 0, ctx.err
	}
	b.initSelectSql()

	query := b.query
	args := b.args

	if ctx.showSql {
		fmt.Println(query, args)
	}
	if ctx.noRun {
		return 0, nil
	}
	rows, err := db.query(query, args...)
	if err != nil {
		return 0, err
	}
	return ctx.ScanT(rows)
}

func (b *SqlBuilder) Exec() (sql.Result, error) {
	db := b.db
	ctx := db.getCtx()
	b.selectStatus = selectDone
	b.whereStatus = whereDone
	if ctx.hasErr() {
		return nil, ctx.err
	}
	b.initSelectSql()
	if ctx.showSql {
		fmt.Println(b.query, b.args)
	}
	if ctx.noRun {
		return nil, nil
	}
	return db.exec(b.query, b.args...)
}
