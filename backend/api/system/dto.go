package system

import "aquila/model"

type UserDto struct {
	Id       int    `form:"id" json:"id"`
	Username string `form:"username" json:"username"` // 账号
	Password string `form:"password" json:"password"` // 密码
	Nickname string `form:"nickname" json:"nickname"` // 昵称
	Status   *int   `form:"status" json:"status"`     // 状态，使用指针类型以区分是否提供了该字段
}

type UserPageVo struct {
	Total int64              `form:"total" json:"total"`
	Data  []model.UserEntity `form:"data" json:"data"`
}

type UserRoleDto struct {
	UserId  int64   `json:"userId"`
	RoleIds []int64 `json:"roleIds"` // 修改为 []int64 类型
}

type GetUserMenuDto struct {
	UserId int64 `form:"userId" json:"userId"`
}

type UserMenuTreeDto struct {
	MenuDto
	Children []UserMenuTreeDto `json:"children"`
}

type RoleDto struct {
	Id     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name"`     // 角色名称
	Remark string `form:"remark" json:"remark"` // 备注
	Status int64  `form:"status" json:"status"` // 状态
}

type RolePageDto struct {
	Total int64              `form:"total" json:"total"`
	Data  []model.RoleEntity `form:"data" json:"data"`
}

type MenuDto struct {
	Id        int    `form:"id" json:"id"`
	Name      string `form:"name" json:"name"`           // 菜单名称
	ParentId  int64  `form:"parentId" json:"parentId"`   // 父菜单ID
	OrderNum  int64  `form:"orderNum" json:"orderNum"`   // 排序
	Path      string `form:"path" json:"path"`           // 路由地址
	Component string `form:"component" json:"component"` // 组件路径
	Query     string `form:"query" json:"query"`         // 请求参数
	IsFrame   int64  `form:"isFrame" json:"isFrame"`     // 是否外链
	MenuType  string `form:"menuType" json:"menuType"`   // 菜单类型
	IsCatch   int64  `form:"isCatch" json:"isCatch"`     // 缓存
	IsHidden  int64  `form:"isHidden" json:"isHidden"`   // 是否隐藏
	Perms     string `form:"perms" json:"perms"`         // 权限标识
	Icon      string `form:"icon" json:"icon"`           // 图标
	Status    int64  `form:"status" json:"status"`       // 状态
	Remark    string `form:"remark" json:"remark"`       // 备注
}

type MenuPageDto struct {
	PageNum  int         `form:"pageNum" json:"pageNum"`   // 页码
	PageSize int         `form:"pageSize" json:"pageSize"` // 每页大小
	Name     string      `form:"name" json:"name"`         // 菜单名称查询条件
	Status   *int64      `form:"status" json:"status"`     // 状态查询条件
	Total    int64       `form:"total" json:"total"`       // 总数
	Data     interface{} `form:"data" json:"data"`         // 数据列表
}

type RoleMenuDto struct {
	RoleId  int64   `form:"roleId" json:"roleId"`
	MenuIds []int64 `form:"menuIds" json:"menuIds"` // 修改为 []int64 类型
}

type ChangePasswordDto struct {
	UserId      int64  `form:"userId" json:"userId"`
	OldPassword string `form:"oldPassword" json:"oldPassword"`
	NewPassword string `form:"newPassword" json:"newPassword"`
}

// RoleMenuTreeReqDto 获取角色菜单树请求
type RoleMenuTreeReqDto struct {
	RoleId int64 `form:"roleId" json:"roleId"`
}

// RoleMenuTreeRespDto 角色菜单树响应
type RoleMenuTreeRespDto struct {
	MenuIds []int64 `json:"menuIds"` // 角色拥有的菜单ID
}
