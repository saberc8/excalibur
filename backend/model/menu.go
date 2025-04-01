package model

import (
	"aquila/global"
)

const TableNameMenuEntity = "menu"

type MenuEntity struct {
	global.GModel
	Name      string `gorm:"column:name;type:varchar(50);comment:菜单名称" json:"name"`
	ParentId  int64  `gorm:"column:parent_id;type:int;comment:父级菜单ID" json:"parentId"`
	OrderNum  int64  `gorm:"column:order_num;type:int;comment:排序" json:"orderNum"`
	Path      string `gorm:"column:path;type:varchar(100);comment:路径" json:"path"`
	Component string `gorm:"column:component;type:varchar(100);comment:组件" json:"component"`
	Query     string `gorm:"column:query;type:varchar(100);comment:参数" json:"query"`
	IsFrame   int64  `gorm:"column:is_frame;type:smallint;comment:是否外链" json:"isFrame"`
	MenuType  string `gorm:"column:menu_type;type:varchar(100);comment:菜单类型" json:"menuType"` // 菜单类型: C目录 M菜单 B按钮
	IsCatch   int64  `gorm:"column:is_catch;type:smallint;comment:是否缓存" json:"isCatch"`
	IsHidden  int64  `gorm:"column:is_hidden;type:smallint;comment:是否可见" json:"isHidden"`
	Perms     string `gorm:"column:perms;type:varchar(100);comment:权限标识" json:"perms"`
	Icon      string `gorm:"column:icon;type:varchar(100);comment:图标" json:"icon"`
	Status    int64  `gorm:"column:status;type:smallint;comment:状态" json:"status"` // 0 正常 1 禁用
	Remark    string `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
}

// TableName UserEntity's table name
func (*MenuEntity) TableName() string {
	return TableNameMenuEntity
}
