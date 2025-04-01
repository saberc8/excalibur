package system

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"

	"github.com/gin-gonic/gin"
)

type Menu struct{}

// CreateMenuApi 创建菜单
func (m *Menu) CreateMenuApi(ctx *gin.Context) {
	var req MenuDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败:"+err.Error())
		return
	}

	// 检查同名菜单
	var count int64
	if err := global.AquilaDb.Model(&model.MenuEntity{}).Where("name = ?", req.Name).Count(&count).Error; err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	if count > 0 {
		utils.Fail(ctx, "菜单名称已存在")
		return
	}

	menu := model.MenuEntity{
		Name:      req.Name,
		ParentId:  req.ParentId,
		OrderNum:  req.OrderNum,
		Path:      req.Path,
		Component: req.Component,
		Query:     req.Query,
		IsFrame:   req.IsFrame,
		MenuType:  req.MenuType,
		IsCatch:   req.IsCatch,
		IsHidden:  req.IsHidden,
		Perms:     req.Perms,
		Icon:      req.Icon,
		Status:    req.Status,
		Remark:    req.Remark,
	}

	if err := global.AquilaDb.Create(&menu).Error; err != nil {
		utils.Fail(ctx, "创建失败")
		return
	}

	utils.Success(ctx, menu)
}

// GetMenuApi 获取单个菜单
func (m *Menu) GetMenuApi(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		utils.Fail(ctx, "参数错误")
		return
	}

	var menu model.MenuEntity
	if err := global.AquilaDb.First(&menu, id).Error; err != nil {
		utils.Fail(ctx, "菜单不存在")
		return
	}

	utils.Success(ctx, menu)
}

// GetMenuPageApi 获取菜单列表
func (m *Menu) GetMenuPageApi(ctx *gin.Context) {
	var req MenuPageDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败:"+err.Error())
		return
	}

	var menus []model.MenuEntity
	var total int64
	db := global.AquilaDb.Model(&model.MenuEntity{})

	// 条件查询
	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	if req.Status != nil {
		db = db.Where("status = ?", *req.Status)
	}

	// 获取总数
	db.Count(&total)

	// 分页查询
	if err := db.Scopes(utils.Paginate(req.PageNum, req.PageSize)).
		Order("order_num ASC, created_at DESC, id ASC"). // 修改这里：先按order_num升序，再按创建时间降序，最后按id升序
		Find(&menus).Error; err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}

	utils.Success(ctx, gin.H{
		"total": total,
		"data":  menus,
	})
}

// UpdateMenuApi 更新菜单
func (m *Menu) UpdateMenuApi(ctx *gin.Context) {
	var req MenuDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败:"+err.Error())
		return
	}

	var menu model.MenuEntity
	if err := global.AquilaDb.First(&menu, req.Id).Error; err != nil {
		utils.Fail(ctx, "菜单不存在")
		return
	}

	menu.Name = req.Name
	menu.ParentId = req.ParentId
	menu.OrderNum = req.OrderNum
	menu.Path = req.Path
	menu.Component = req.Component
	menu.Query = req.Query
	menu.IsFrame = req.IsFrame
	menu.MenuType = req.MenuType
	menu.IsCatch = req.IsCatch
	menu.IsHidden = req.IsHidden
	menu.Perms = req.Perms
	menu.Icon = req.Icon
	menu.Status = req.Status
	menu.Remark = req.Remark

	if err := global.AquilaDb.Save(&menu).Error; err != nil {
		utils.Fail(ctx, "更新失败")
		return
	}

	utils.Success(ctx, nil)
}

// DeleteMenuApi 删除菜单
func (m *Menu) DeleteMenuApi(ctx *gin.Context) {
	var req MenuDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败:"+err.Error())
		return
	}

	// 检查是否有子菜单
	var count int64
	if err := global.AquilaDb.Model(&model.MenuEntity{}).Where("parent_id = ?", req.Id).Count(&count).Error; err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	if count > 0 {
		utils.Fail(ctx, "存在子菜单,不能删除")
		return
	}

	// 删除菜单
	if err := global.AquilaDb.Delete(&model.MenuEntity{}, req.Id).Error; err != nil {
		utils.Fail(ctx, "删除失败")
		return
	}

	utils.Success(ctx, nil)
}

// GetMenuRolesApi 获取拥有该菜单访问权限的角色列表
func (m *Menu) GetMenuRolesApi(ctx *gin.Context) {
	menuId := ctx.Query("menuId")
	if menuId == "" {
		utils.Fail(ctx, "参数错误")
		return
	}

	var roles []model.RoleEntity
	if err := global.AquilaDb.Joins("JOIN role_menu ON role_menu.role_id = role.id").
		Where("role_menu.menu_id = ?", menuId).
		Find(&roles).Error; err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}

	utils.Success(ctx, roles)
}

// GetMenuTreeApi 获取菜单树形结构
func (m *Menu) GetMenuTreeApi(ctx *gin.Context) {
	var menus []model.MenuEntity
	if err := global.AquilaDb.Order("order_num ASC, created_at DESC, id ASC").Find(&menus).Error; err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}

	tree := getMenuTree(0, menus)
	utils.Success(ctx, tree)
}

// GetRoleMenuTreeApi 获取角色的菜单树
func (m *Menu) GetRoleMenuTreeApi(ctx *gin.Context) {
	var req RoleMenuTreeReqDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, err.Error())
		return
	}

	// 1. 查询角色关联的菜单ID
	var roleMenus []model.RoleMenuEntity
	err := global.AquilaDb.Where("role_id = ?", req.RoleId).Find(&roleMenus).Error
	if err != nil {
		utils.Fail(ctx, "查询角色菜单关系失败")
		return
	}

	// 2. 查询所有菜单
	var allMenus []model.MenuEntity
	err = global.AquilaDb.Find(&allMenus).Error
	if err != nil {
		utils.Fail(ctx, "查询菜单失败")
		return
	}

	// 4. 标记角色拥有的菜单
	menuIds := make([]int64, 0)
	for _, rm := range roleMenus {
		menuIds = append(menuIds, rm.MenuId)
	}

	resp := RoleMenuTreeRespDto{
		MenuIds: menuIds,
	}

	utils.Success(ctx, resp)
}
