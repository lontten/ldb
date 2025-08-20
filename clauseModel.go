package ldb

// 句子类型，用于whereBuilder
type clauseType int

const (
	Eq clauseType = iota
	Neq
	Less
	LessEq
	Greater
	GreaterEq
	Like
	NotLike
	In
	NotIn
	Between
	NotBetween
	IsNull
	IsNotNull
	IsFalse

	PrimaryKeys       // 主键
	FilterPrimaryKeys // 过滤主键

	// Contains 包含
	// pg 独有
	// [1] @< [1,2]
	Contains
)
