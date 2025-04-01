package manager

import (
	"aquila/global"
	"aquila/model"
	"aquila/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Task struct{}

// CreateTaskApi 创建任务
func (t Task) CreateTaskApi(ctx *gin.Context) {
	var req TaskDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败: "+err.Error())
		return
	}

	userId, _ := ctx.Get("uid")
	task := model.TaskEntity{
		Title:      req.Title,
		Content:    req.Content,
		ProjectId:  req.ProjectId,
		AssigneeId: req.AssigneeId,
		CreatorId:  int64(userId.(int)),
		Priority:   0,
		Status:     0,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	}

	if err := global.AquilaDb.Create(&task).Error; err != nil {
		utils.Fail(ctx, "创建任务失败")
		return
	}

	utils.Success(ctx, nil)
}

// GetTaskListApi 获取任务列表
func (t Task) GetTaskListApi(ctx *gin.Context) {
	pageNum := ctx.DefaultQuery("pageNum", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	projectId := ctx.Query("projectId")
	status := ctx.Query("status")

	query := global.AquilaDb.Model(&model.TaskEntity{})

	if projectId != "" {
		query = query.Where("project_id = ?", projectId)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	var tasks []model.TaskEntity
	var total int64

	pageNumInt, _ := strconv.Atoi(pageNum)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	query.Count(&total)
	if err := query.Order("id desc").Scopes(utils.Paginate(pageNumInt, pageSizeInt)).Find(&tasks).Error; err != nil {
		utils.Fail(ctx, "查询失败")
		return
	}

	utils.Success(ctx, gin.H{
		"total": total,
		"data":  tasks,
	})
}

// GetTaskApi 获取单个任务
func (t Task) GetTaskApi(ctx *gin.Context) {
	id := ctx.Query("id")
	var task model.TaskEntity

	if err := global.AquilaDb.First(&task, id).Error; err != nil {
		utils.Fail(ctx, "任务不存在")
		return
	}

	utils.Success(ctx, task)
}

// UpdateTaskApi 更新任务
func (t Task) UpdateTaskApi(ctx *gin.Context) {
	var req TaskDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败")
		return
	}

	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Content != "" {
		updates["content"] = req.Content
	}
	if req.AssigneeId != 0 {
		updates["assignee_id"] = req.AssigneeId
	}
	if req.Priority != nil {
		updates["priority"] = *req.Priority
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.StartDate != "" {
		updates["start_date"] = req.StartDate
	}
	if req.EndDate != "" {
		updates["end_date"] = req.EndDate
	}

	if err := global.AquilaDb.Model(&model.TaskEntity{}).Where("id = ?", req.Id).Updates(updates).Error; err != nil {
		utils.Fail(ctx, "更新失败")
		return
	}

	utils.Success(ctx, nil)
}

// DeleteTaskApi 删除任务
func (t Task) DeleteTaskApi(ctx *gin.Context) {
	var req TaskDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败")
		return
	}

	if err := global.AquilaDb.Delete(&model.TaskEntity{}, req.Id).Error; err != nil {
		utils.Fail(ctx, "删除失败")
		return
	}

	utils.Success(ctx, nil)
}
