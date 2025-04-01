package manager

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Project struct{}

// CreateProjectApi 创建项目
func (p Project) CreateProjectApi(ctx *gin.Context) {
	var req ProjectDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败: "+err.Error())
		return
	}

	userId, _ := ctx.Get("uid")
	project := model.ProjectEntity{
		Name:        req.Name,
		SpaceId:     req.SpaceId,
		Description: req.Description,
		OwnerId:     int64(userId.(int)),
		Status:      0,
	}

	if err := global.AquilaDb.Create(&project).Error; err != nil {
		utils.Fail(ctx, "创建项目失败")
		return
	}

	utils.Success(ctx, nil)
}

// GetProjectListApi 获取项目列表
func (p Project) GetProjectListApi(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	spaceId := ctx.Query("spaceId")

	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	query := global.AquilaDb.Model(&model.ProjectEntity{})
	if spaceId != "" {
		query = query.Where("space_id = ?", spaceId)
	}

	var projects []model.ProjectEntity
	var total int64

	query.Count(&total)
	if err := query.Order("id desc").Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&projects).Error; err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}

	utils.Success(ctx, gin.H{
		"total": total,
		"data":  projects,
	})
}

// GetProjectApi 获取单个项目
func (p Project) GetProjectApi(ctx *gin.Context) {
	id := ctx.Query("id")
	var project model.ProjectEntity

	if err := global.AquilaDb.First(&project, id).Error; err != nil {
		utils.Fail(ctx, "项目不存在")
		return
	}

	utils.Success(ctx, project)
}

// UpdateProjectApi 更新项目
func (p Project) UpdateProjectApi(ctx *gin.Context) {
	var req ProjectDto
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
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	if err := global.AquilaDb.Model(&model.ProjectEntity{}).Where("id = ?", req.Id).Updates(updates).Error; err != nil {
		utils.Fail(ctx, "更新失败")
		return
	}

	utils.Success(ctx, nil)
}

// DeleteProjectApi 删除项目
func (p Project) DeleteProjectApi(ctx *gin.Context) {
	var req ProjectDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败")
		return
	}

	if err := global.AquilaDb.Delete(&model.ProjectEntity{}, req.Id).Error; err != nil {
		utils.Fail(ctx, "删除失败")
		return
	}

	utils.Success(ctx, nil)
}
