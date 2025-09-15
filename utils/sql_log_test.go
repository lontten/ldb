package utils

import (
	"testing"

	"github.com/lontten/lcore/v2/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestDbArg2String(t *testing.T) {
	as := assert.New(t)

	var a1 = types.LocalDateOfYmd(2023, 1, 1)
	as.Equal("2023-01-01", dbArg2String(true, a1))

	var a1p = &a1
	as.Equal("2023-01-01", dbArg2String(true, a1p))

	var a2 = 1
	as.Equal("1", dbArg2String(true, a2))

	var a3 = 1.1
	as.Equal("1.1", dbArg2String(true, a3))

	var a4 = true
	as.Equal("true", dbArg2String(true, a4))

	var a5, _ = decimal.NewFromString("1.1")
	as.Equal("1.1", dbArg2String(true, a5))

}
