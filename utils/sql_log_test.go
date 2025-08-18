package utils

import (
	"github.com/lontten/lcore/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDbArg2String(t *testing.T) {
	as := assert.New(t)

	var a1 = types.DateOfYmd(2023, 1, 1)
	as.Equal("2023-01-01", DbArg2String(a1))

	var a1p = &a1
	as.Equal("2023-01-01", DbArg2String(a1p))

	var a2 = 1
	as.Equal("1", DbArg2String(a2))

	var a3 = 1.1
	as.Equal("1.1", DbArg2String(a3))

	var a4 = true
	as.Equal("true", DbArg2String(a4))

	var a5, _ = decimal.NewFromString("1.1")
	as.Equal("1.1", DbArg2String(a5))

}
