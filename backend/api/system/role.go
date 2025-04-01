package system

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Role struct{}

func (r Role) CreateRoleApi(ctx *gin.Context) {
	var req RoleDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败："+err.Error())
		return
	}
	var role model.RoleEntity
	err := global.AquilaDb.Where("name = ?", req.Name).First(&role).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			utils.Fail(ctx, "查询角色失败："+err.Error())
			return
		}

		role = model.RoleEntity{
			Name:   req.Name,
			Remark: req.Remark,
			Status: req.Status,
		}
		err = global.AquilaDb.Create(&role).Error
		if err != nil {
			global.AquilaLog.Error("创建角色失败：" + err.Error())
			utils.Fail(ctx, "创建角色失败："+err.Error())
			return
		}
		utils.Success(ctx, role)
		return
	}
	utils.Fail(ctx, "角色名已存在")
}

func (r Role) GetRoleApi(ctx *gin.Context) {
	var req RoleDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
}

// updateRoleApi 更新角色
func (r Role) UpdateRoleApi(ctx *gin.Context) {
	var req RoleDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}
	var role model.RoleEntity
	err := global.AquilaDb.Where("id = ?", req.Id).First(&role).Error
	if err != nil {
		utils.Fail(ctx, "角色不存在")
		return
	}
	role.Name = req.Name
	role.Remark = req.Remark
	role.Status = req.Status
	err = global.AquilaDb.Save(&role).Error
	if err != nil {
		utils.Fail(ctx, "角色更新失败")
		return
	}
	utils.Success(ctx, nil)
}

func (r Role) GetRolePageApi(ctx *gin.Context) {
	var req RolePageDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "---step1---"+err.Error())
		return
	}

	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	var roles []model.RoleEntity
	var total int64

	global.AquilaDb.Model(&model.RoleEntity{}).Count(&total)
	global.AquilaDb.Model(&model.RoleEntity{}).
		Order("id ASC"). // 添加这一行，按id升序排序
		Scopes(utils.Paginate(pageNumInt, pageSizeInt)).
		Find(&roles)

	req.Total = total
	req.Data = roles

	utils.Success(ctx, req)
}

// BindMenuApi 绑定菜单
func (r Role) BindMenuApi(ctx *gin.Context) {
	var req RoleMenuDto
	err := ctx.ShouldBind(&req)
	if err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	var role model.RoleEntity
	err = global.AquilaDb.Where("id = ?", req.RoleId).First(&role).Error
	if err != nil {
		utils.Fail(ctx, "角色不存在")
		return
	}

	err = global.AquilaDb.Transaction(func(tx *gorm.DB) error {
		// 先删除该角色的所有菜单关联
		if err := tx.Where("role_id = ?", req.RoleId).Delete(&model.RoleMenuEntity{}).Error; err != nil {
			return err
		}

		// 如果没有要绑定的菜单，直接返回
		if len(req.MenuIds) == 0 {
			return nil
		}

		// 构建批量插入的数据
		roleMenus := make([]model.RoleMenuEntity, 0, len(req.MenuIds))
		for _, menuId := range req.MenuIds {
			roleMenus = append(roleMenus, model.RoleMenuEntity{
				RoleId: req.RoleId,
				MenuId: menuId,
			})
		}

		// 批量插入新的关联关系
		return tx.Create(&roleMenus).Error
	})

	if err != nil {
		utils.Fail(ctx, "绑定失败"+err.Error())
		return
	}
	utils.Success(ctx, nil)
}

// DeleteRoleApi 删除角色
func (r Role) DeleteRoleApi(ctx *gin.Context) {
	var req RoleDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	err := global.AquilaDb.Delete(&model.RoleEntity{}, req.Id).Error
	if err != nil {
		utils.Fail(ctx, "删除角色失败")
		return
	}
	utils.Success(ctx, nil)
}

// UnbindMenuApi 解绑菜单
func (r Role) UnbindMenuApi(ctx *gin.Context) {
	var req RoleMenuDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败"+err.Error())
		return
	}

	err := global.AquilaDb.Where("role_id = ? AND menu_id IN (?)",
		req.RoleId, req.MenuIds).Delete(&model.RoleMenuEntity{}).Error
	if err != nil {
		utils.Fail(ctx, "解绑菜单失败")
		return
	}
	utils.Success(ctx, nil)
}

// GetRoleMenusApi 获取角色的菜单列表
func (r Role) GetRoleMenusApi(ctx *gin.Context) {
	roleId := ctx.Query("id") // 修改这里，从 "roleId" 改为 "id"
	if roleId == "" {
		utils.Fail(ctx, "角色ID不能为空")
		return
	}

	var menus []model.MenuEntity
	err := global.AquilaDb.Joins("JOIN role_menu ON role_menu.menu_id = menu.id").
		Where("role_menu.role_id = ?", roleId).
		Find(&menus).Error

	if err != nil {
		global.AquilaLog.Error("查询角色菜单失败：" + err.Error())
		utils.Fail(ctx, "查询失败")
		return
	}
	utils.Success(ctx, menus)
}

// GetRoleUsersApi 获取拥有该角色的用户列表
func (r Role) GetRoleUsersApi(ctx *gin.Context) {
	roleId := ctx.Query("roleId")
	var users []model.UserEntity

	err := global.AquilaDb.Joins("JOIN user_role ON user_role.user_id = user.id").
		Where("user_role.role_id = ?", roleId).
		Find(&users).Error

	if err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}
	utils.Success(ctx, users)
}
