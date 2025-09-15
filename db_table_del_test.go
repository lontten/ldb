package ldb

import "github.com/lontten/lcore/v2/types"

// import (
//
//	"github.com/DATA-DOG/go-sqlmock"
//	"github.com/lontten/lcore/v2/types"
//	"github.com/stretchr/testify/assert"
//	"testing"
//
// )

//	func TestDeleteByPrimaryKey(t *testing.T) {
//		as := assert.New(t)
//		ldb, mock, err := sqlmock.New()
//		as.Nil(err, "new sqlmock error")
//		engine := MustConnectMock(ldb, &PgConf{})
//
//		mock.ExpectExec("DELETE FROM *").
//			WithArgs(1).
//			WillReturnError(nil).
//			WillReturnResult(sqlmock.NewResult(0, 1))
//
//		num, err := engine.Delete(User{}).ByPrimaryKey(1).Result()
//		as.Nil(err)
//		as.Equal(int64(1), num)
//
//		as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
//	}
//
//	func TestDeleteByPrimaryKeys(t *testing.T) {
//		as := assert.New(t)
//		ldb, mock, err := sqlmock.New()
//		as.Nil(err, "new sqlmock error")
//		engine := MustConnectMock(ldb, &PgConf{})
//
//		mock.ExpectExec("DELETE FROM *").
//			WithArgs(1, 2, 3).
//			WillReturnError(nil).
//			WillReturnResult(sqlmock.NewResult(0, 3))
//
//		num, err := engine.Delete(User{}).ByPrimaryKey(1, 2, 3).Result()
//		as.Nil(err)
//		as.Equal(int64(3), num, "num error")
//
//		as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
//	}
type Whe struct {
	Name *string
	Age  *int
	Uid  *types.UUID
}

//
//func TestDeleteByModel(t *testing.T) {
//	as := assert.New(t)
//	ldb, mock, err := sqlmock.New()
//	as.Nil(err, "new sqlmock error")
//	engine := MustConnectMock(ldb, &PgConf{})
//
//	mock.ExpectExec("DELETE FROM *").
//		WithArgs("kk").
//		WillReturnError(nil).
//		WillReturnResult(sqlmock.NewResult(0, 1))
//
//	mock.ExpectExec("DELETE FROM *").
//		WithArgs(233, "kk").
//		WillReturnError(nil).
//		WillReturnResult(sqlmock.NewResult(0, 1))
//
//	num, err := engine.Delete(User{}).ByModel(Whe{
//		Name: types.NewString("kk"),
//		Age:  nil,
//		Uid:  nil,
//	}).Result()
//	as.Nil(err)
//	as.Equal(int64(1), num)
//
//	num, err = engine.Delete(User{}).ByModel(Whe{
//		Name: types.NewString("kk"),
//		Age:  types.NewInt(233),
//		Uid:  nil,
//	}).Result()
//	as.Nil(err)
//	as.Equal(int64(1), num)
//
//	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
//}
//
//func TestDeleteByWhere(t *testing.T) {
//	as := assert.New(t)
//	ldb, mock, err := sqlmock.New()
//	as.Nil(err, "new sqlmock error")
//	engine := MustConnectMock(ldb, &PgConf{})
//
//	mock.ExpectExec("DELETE FROM *").
//		WithArgs("kk", "kk").
//		WillReturnError(nil).
//		WillReturnResult(sqlmock.NewResult(0, 1))
//
//	num, err := engine.Delete(User{}).ByWhere(new(WhereBuilder).
//		Eq("name", "kk").
//		Like("age", types.NewString("kk")),
//	).Result()
//	as.Nil(err)
//	as.Equal(int64(1), num)
//
//	as.Nil(mock.ExpectationsWereMet(), "we make sure that all expectations were met")
//}
