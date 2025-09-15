package utils

import (
	"reflect"
	"testing"

	"github.com/lontten/ldb/v2/softdelete"
)

type TestSoftDel1 struct {
	softdelete.DeleteGormMilli
}
type T1 struct {
}
type TestSoftDel2 struct {
	T1
	TestSoftDel1
}

func TestCheckSoftDelType(t *testing.T) {

	delType := GetSoftDelType(reflect.TypeOf(TestSoftDel2{}))
	t.Log(delType)
}

func TestIsSoftDelFieldType(t *testing.T) {

	delType := IsSoftDelFieldType(reflect.TypeOf(TestSoftDel2{}))
	t.Log(delType)
}
