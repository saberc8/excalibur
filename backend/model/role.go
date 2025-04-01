package model

import (
	"aquila/global"
)

const TableNameRoleEntity = "role"

type RoleEntity struct {
	global.GModel
	Name   string `gorm:"column:name;type:varchar(50);comment:名称" json:"name"`
	Remark string `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
	Status int64  `gorm:"column:status;type:smallint;comment:状态" json:"status"` // 0 正常 1 禁用

}

// TableName UserEntity's table name
func (*RoleEntity) TableName() string {
	return TableNameRoleEntity
}
