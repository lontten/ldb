package ldb

//
//import (
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/lontten/lcore/types"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestSelectByPrimaryKey(t *testing.T) {
//	as := assert.New(t)
//	ldb, mock, err := sqlmock.New()
//	as.Nil(err, "new sqlmock error")
//	engine := MustConnectMock(ldb, &PgConf{})
//
//	mock.ExpectQuery("SELECT *").
//		WithArgs(1).
//		WillReturnError(nil).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "test"))
//
//	user := User{}
//	num, err := engine.Select(user).ByPrimaryKey(1).ScanOne(&user)
//	as.Nil(err)
//	as.Equal(int64(1), num)
//	as.Equal(int64(1), user.Id)
//	as.Equal("test", user.Name)
//
//	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
//}
//
//func TestSelectByPrimaryKeys(t *testing.T) {
//	as := assert.New(t)
//	ldb, mock, err := sqlmock.New()
//	as.Nil(err, "new sqlmock error")
//	engine := MustConnectMock(ldb, &PgConf{})
//
//	mock.ExpectQuery("SELECT *").
//		WithArgs(1, 2).
//		WillReturnError(nil).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
//			AddRow(1, "test").
//			AddRow(2, "test2"),
//		)
//
//	users := make([]User, 0)
//	num, err := engine.Select(User{}).ByPrimaryKey(1, 2).ScanList(&users)
//	as.Nil(err)
//	as.Equal(int64(2), num)
//	as.Equal(2, len(users))
//	as.Equal(int64(1), users[0].Id)
//	as.Equal("test", users[0].Name)
//	as.Equal(int64(2), users[1].Id)
//	as.Equal("test2", users[1].Name)
//
//	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
//}
//
//func TestSelectByModel(t *testing.T) {
//	as := assert.New(t)
//	ldb, mock, err := sqlmock.New()
//	as.Nil(err, "new sqlmock error")
//	engine := MustConnectMock(ldb, &PgConf{})
//
//	mock.ExpectQuery("SELECT *").
//		WithArgs("123").
//		WillReturnError(nil).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
//			AddRow(1, "test"),
//		)
//
//	mock.ExpectQuery("SELECT *").
//		WithArgs(233, "kk").
//		WillReturnError(nil).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
//			AddRow(1, "test").
//			AddRow(2, "test2"),
//		)
//
//	user := User{}
//	num, err := engine.Select(User{}).ByModel(Whe{
//		Name: types.NewString("123"),
//		Age:  nil,
//		Uid:  nil,
//	}).ScanOne(&user)
//	as.Nil(err)
//	as.Equal(int64(1), num)
//	as.Equal(int64(1), user.Id)
//	as.Equal("test", user.Name)
//
//	users := make([]User, 0)
//	num, err = engine.Select(User{}).ByModel(Whe{
//		Name: types.NewString("kk"),
//		Age:  types.NewInt(233),
//		Uid:  nil,
//	}).ScanList(&users)
//	as.Nil(err)
//	as.Equal(int64(2), num)
//	as.Equal(2, len(users))
//	as.Equal(int64(1), users[0].Id)
//	as.Equal("test", users[0].Name)
//	as.Equal(int64(2), users[1].Id)
//	as.Equal("test2", users[1].Name)
//
//	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
//}
//
//func TestSelectByWhere(t *testing.T) {
//	as := assert.New(t)
//	ldb, mock, err := sqlmock.New()
//	as.Nil(err, "new sqlmock error")
//	engine := MustConnectMock(ldb, &PgConf{})
//
//	//---------------------------scan one------------------------
//
//	mock.ExpectQuery("SELECT *").
//		WithArgs("kk").
//		WillReturnError(nil).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
//			AddRow(1, "test"),
//		)
//	user := User{}
//	num, err := engine.Select(User{}).ByWhere(new(WhereBuilder).
//		Eq("name", "kk"),
//	).ScanOne(&user)
//	as.Nil(err)
//	as.Equal(int64(1), num)
//	as.Equal(int64(1), user.Id)
//	as.Equal("test", user.Name)
//
//	//---------------------------scan list------------------------
//	mock.ExpectQuery("SELECT *").
//		WithArgs("kk", 233).
//		WillReturnError(nil).
//		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
//			AddRow(1, "test").
//			AddRow(2, "test2"),
//		)
//	users := make([]User, 0)
//	num, err = engine.Select(User{}).ByWhere(new(WhereBuilder).
//		Eq("name", "kk").
//		Eq("age", 233),
//	).ScanList(&users)
//	as.Nil(err)
//	as.Equal(int64(2), num)
//	as.Equal(2, len(users))
//	as.Equal(int64(1), users[0].Id)
//	as.Equal("test", users[0].Name)
//	as.Equal(int64(2), users[1].Id)
//	as.Equal("test2", users[1].Name)
//
//	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
//}
