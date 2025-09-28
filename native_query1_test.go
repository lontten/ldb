package ldb

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lontten/lcore/v2/types"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

type Ka struct {
	Id   int
	Name string
	Day1 types.LocalDate
	Day2 types.LocalDate
}

func (k Ka) TableConf() *TableConf {
	return new(TableConf).
		Table("t_ka").
		AutoPrimaryKey("id")
}

func TestQuery1_pg(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, "new sqlmock error")
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectQuery("select q1").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "day1", "day2"}).
			AddRow(12, "nil", nil, "2022-02-02"),
		)

	list, err := QueryList[Ka](engine, "select q1")
	as.Nil(err)
	as.NotNil(list)
	as.Equal(1, len(list), "list length error")
	ka := list[0]
	as.Equal(12, ka.Id, "id error")
	as.Equal("nil", ka.Name, "name error")
	as.True(ka.Day1.IsZero(), "day1 error")
	as.Equal("2022-02-02", ka.Day2.String(), "day2 error")
}

type UserNil2 struct {
	Id    int
	Name  string
	Money decimal.Decimal
	Day1  types.LocalDate
	Day2  time.Time
}

func (u UserNil2) TableConf() *TableConf {
	return new(TableConf).
		Table("t_user").
		AutoPrimaryKey("id")
}

func TestQuery2_pg(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, "new sqlmock error")
	engine := MustConnectMock(db, &PgConf{})

	mock.ExpectQuery("select q1").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "money", "day1", "day2"}).
			AddRow(12, nil, nil, nil, nil),
		)

	list, err := QueryList[UserNil2](engine, "select q1")
	as.Nil(err)
	as.NotNil(list)
	as.Equal(1, len(list), "list length error")
	ka := list[0]
	as.Equal(12, ka.Id, "id error")
	as.Equal("", ka.Name, "name error")
	as.True(ka.Day1.IsZero(), "day1 error")
	as.True(ka.Day2.IsZero(), "day2 error")
}
