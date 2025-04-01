package manager

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Space struct{}

// CreateSpaceApi 创建空间
func (s Space) CreateSpaceApi(ctx *gin.Context) {
	var req SpaceDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败: "+err.Error())
		return
	}

	userId, _ := ctx.Get("uid")
	space := model.SpaceEntity{
		Name:        req.Name,
		Description: req.Description,
		Path:        req.Path,
		OwnerId:     int64(userId.(int)), // 修改这里：将 int 转换为 int64
		Status:      0,
	}

	if err := global.AquilaDb.Create(&space).Error; err != nil {
		utils.Fail(ctx, "创建空间失败")
		return
	}

	utils.Success(ctx, nil)
}

// GetSpaceListApi 获取空间列表
func (s Space) GetSpaceListApi(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	var spaces []model.SpaceEntity
	var total int64

	query := global.AquilaDb.Model(&model.SpaceEntity{})

	// 检查用户ID
	userId, exists := ctx.Get("uid")
	if !exists {
		utils.Fail(ctx, "用户ID获取失败")
		return
	}

	// 直接使用用户ID过滤
	query = query.Where("owner_id = ?", userId)

	query.Count(&total)
	if err := query.Order("id desc").Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&spaces).Error; err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}

	utils.Success(ctx, gin.H{
		"total": total,
		"data":  spaces,
	})
}

// GetSpaceApi 获取单个空间
func (s Space) GetSpaceApi(ctx *gin.Context) {
	id := ctx.Query("id")
	var space model.SpaceEntity

	if err := global.AquilaDb.First(&space, id).Error; err != nil {
		utils.Fail(ctx, "空间不存在")
		return
	}

	utils.Success(ctx, space)
}

// UpdateSpaceApi 更新空间
func (s Space) UpdateSpaceApi(ctx *gin.Context) {
	var req SpaceDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败")
		return
	}

	updates := make(map[string]interface{})
	if req.Name != "" {
		updates["name"] = req.Name
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Path != "" {
		updates["path"] = req.Path
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := global.AquilaDb.Model(&model.SpaceEntity{}).Where("id = ?", req.Id).Updates(updates).Error; err != nil {
		utils.Fail(ctx, "更新失败")
		return
	}

	utils.Success(ctx, nil)
}

// DeleteSpaceApi 删除空间
func (s Space) DeleteSpaceApi(ctx *gin.Context) {
	var req SpaceDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败")
		return
	}

	if err := global.AquilaDb.Delete(&model.SpaceEntity{}, req.Id).Error; err != nil {
		utils.Fail(ctx, "删除失败")
		return
	}

	utils.Success(ctx, nil)
}
