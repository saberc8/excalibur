package model

import (
	"aquila/global"
)

const TableNameUserRoleEntity = "user_role"

type UserRoleEntity struct {
	global.GModel
	UserId int64 `gorm:"column:user_id;type:int;comment:用户ID" json:"userId"`
	RoleId int64 `gorm:"column:role_id;type:int;comment:角色ID" json:"roleId"`
}

// TableName UserEntity's table name
func (*UserRoleEntity) TableName() string {
	return TableNameUserRoleEntity
}
