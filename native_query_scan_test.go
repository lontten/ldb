package ldb

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lontten/lcore/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, "new sqlmock error")
	engine := MustConnectMock(db, &PgConf{})

	//-------------base------------

	mock.ExpectQuery("select 1").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(4),
		)

	n := 0

	num, err := NativeQueryScan(engine, "select 1").ScanOne(&n)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
	as.Equal(4, n, "n error")

	mock.ExpectQuery("select 'kk' ").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow("kk"),
		)

	name := ""
	num, err = NativeQueryScan(engine, "select 'kk' ").ScanOne(&name)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
	as.Equal("kk", name, "name error")

	//-------------------uuid---------------

	v4 := types.NewV4()
	mock.ExpectQuery("select gen_random_uuid() ").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(v4),
		)

	uid := types.UUID{}
	num, err = NativeQueryScan(engine, "select gen_random_uuid() ").ScanOne(&uid)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
	as.Equal(v4, uid, "uuid error")

	//-------------------date---------------
	date := types.NowDate()
	mock.ExpectQuery("select now() ").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(date),
		)

	d := types.Date{}
	num, err = NativeQueryScan(engine, "select now() ").ScanOne(&d)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
	as.Equal(date, d, "date error")
	as.NotEqual(types.Date{}, d, "date error")

	//-------------------user---------------
	user := User{Id: 1, Name: "lontten"}
	mock.ExpectQuery("select id,name from user limit 1").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "lontten"),
		)

	u := User{}
	num, err = NativeQueryScan(engine, "select id,name from user limit 1").ScanOne(&u)
	as.Nil(err)
	as.Equal(int64(1), num, "num error")
	as.Equal(user, u, "user error")

	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
}

func TestQueryT(t *testing.T) {
	as := assert.New(t)
	db, mock, err := sqlmock.New()
	as.Nil(err, "new sqlmock error")
	engine := MustConnectMock(db, &PgConf{})

	//-------------base------------

	mock.ExpectQuery("select 1").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(4),
		)
	n, err := QueryOne[int](engine, "select 1")
	as.Nil(err)
	as.NotNil(n, "n error")
	as.Equal(4, *n, "n error")

	mock.ExpectQuery("select 'kk' ").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow("kk"),
		)
	name, err := QueryOne[string](engine, "select 'kk' ")
	as.Nil(err)
	as.NotNil(name, "s error")
	as.Equal("kk", *name, "name error")

	//-------------------uuid---------------

	v4 := types.NewV4()
	mock.ExpectQuery("select gen_random_uuid() ").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(v4),
		)

	uid, err := QueryOne[types.UUID](engine, "select gen_random_uuid() ")
	as.Nil(err)
	as.NotNil(uid, "uuid error")
	as.Equal(v4, *uid, "uuid error")

	//-------------------date---------------
	date := types.NowDate()
	mock.ExpectQuery("select now() ").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{""}).
			AddRow(date),
		)

	d, err := QueryOne[types.Date](engine, "select now() ")
	as.Nil(err)
	as.NotNil(d, "d error")
	as.Equal(date, *d, "date error")
	as.NotEqual(types.Date{}, *d, "date error")

	//-------------------user---------------
	user := User{Id: 1, Name: "lontten"}
	mock.ExpectQuery("select id,name from user limit 1").
		WillReturnError(nil).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "lontten"),
		)

	u, err := QueryOne[User](engine, "select id,name from user limit 1")
	as.Nil(err)
	as.NotNil(u, "u error")
	as.Equal(user, *u, "user error")

	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
}
func TestQueryT1(t *testing.T) {
	as := assert.New(t)
	date := types.NowDate()
	var p = &date
	value, err := p.Value()
	fmt.Println(value)
	as.Nil(err)
}
