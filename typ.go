package ldb

type ArgArray []any

func ArrayOf(v ...any) ArgArray {
	return v
}
