package dbinit

import (
	"github.com/lontten/lcore/v2/types"
	"github.com/lontten/ldb/v2"
	"github.com/lontten/ldb/v2/softdelete"
	"gorm.io/gorm"
)

type TestModel struct {
	softdelete.DeleteTimeNil
	Id   *int    `db:"id"`
	Name *string `db:"name"`
}

func (u TestModel) TableConf() *ldb.TableConfContext {
	return new(ldb.TableConfContext).Table("t_test").
		PrimaryKeys("id").
		AutoColumn("id")
}

type LN_MODEL_DEL struct {
	CreatedAt *types.LocalDateTime `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;"` //创建时间
	UpdatedAt *types.LocalDateTime `json:"updatedAt" form:"updatedAt" gorm:"column:updated_at;comment:更新时间;"`
	DeletedAt gorm.DeletedAt       `gorm:"index" json:"-"` // 删除时间
}

type TestModelDel struct {
	softdelete.DeleteTimeNil
	LN_MODEL_DEL

	Id   *int    `db:"id"`
	Name *string `db:"name"`
}

func (u TestModelDel) TableConf() *ldb.TableConfContext {
	return ldb.TableConf("t_test").
		PrimaryKeys("id").
		AutoColumn("id")
}
