package dbinit

import (
	"github.com/lontten/lcore/types"
	"github.com/lontten/ldb"
	"github.com/lontten/ldb/softdelete"
	"gorm.io/gorm"
)

type TestModel struct {
	softdelete.DeleteTimeNil
	Id   *int    `db:"id"`
	Name *string `db:"name"`
}

func (u TestModel) TableConf() *ldb.TableConf {
	return new(ldb.TableConf).Table("t_test").
		PrimaryKeys("id").
		AutoPrimaryKey("id")
}

type LN_MODEL_DEL struct {
	CreatedAt *types.DateTime `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"` //创建时间
	UpdatedAt *types.DateTime `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"` // 删除时间
}

type TestModelDel struct {
	softdelete.DeleteTimeNil
	LN_MODEL_DEL

	Id   *int    `db:"id"`
	Name *string `db:"name"`
}

func (u TestModelDel) TableConf() *ldb.TableConf {
	return new(ldb.TableConf).Table("t_test").
		PrimaryKeys("id").
		AutoPrimaryKey("id")
}
