package model

import (
	"aquila/global"
)

const TableNameRoleMenuEntity = "role_menu"

type RoleMenuEntity struct {
	global.GModel
	RoleId int64 `gorm:"column:role_id;type:int;comment:角色ID" json:"roleId"`
	MenuId int64 `gorm:"column:menu_id;type:int;comment:菜单ID" json:"menuId"`
}

// TableName UserEntity's table name
func (*RoleMenuEntity) TableName() string {
	return TableNameRoleMenuEntity
}
