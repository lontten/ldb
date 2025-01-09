package ldb

import (
	"fmt"
	"github.com/lontten/ldb/softdelete"
	"reflect"
	"testing"
)

type Ka struct {
	Ka1 *string
	Ka2 *string

	softdelete.DeleteGormMilli `db:"abc"`
}

type Kb struct {
	Ka
}
type Kc struct {
	Kb
}

func Test_getStructColName2fieldNameMap(t *testing.T) {
	var k Kc
	v := reflect.ValueOf(k)
	cv := getStructC(v.Type())
	for a, b := range cv {
		fmt.Println(a, b)
	}
}

func Test_getStructColName2fieldNameAllMap(t *testing.T) {
	var k Kc
	v := reflect.ValueOf(k)
	cv := getStructC(v.Type())
	for a, b := range cv {
		fmt.Println(a, b)
	}
}
