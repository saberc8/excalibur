package manager

// SpaceDto 空间数据传输对象
type SpaceDto struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	OwnerId     int64  `json:"ownerId"`
	Path        string `json:"path" binding:"required"`
	Status      *int   `json:"status"`
}

// ProjectDto 项目数据传输对象
type ProjectDto struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	SpaceId     int64  `json:"spaceId" binding:"required"`
	Description string `json:"description"`
	OwnerId     int64  `json:"ownerId"`
	Status      *int   `json:"status"`
}

// TaskDto 任务数据传输对象
type TaskDto struct {
	Id         int64  `json:"id"`
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content"`
	ProjectId  int64  `json:"projectId" binding:"required"`
	AssigneeId int64  `json:"assigneeId"`
	CreatorId  int64  `json:"creatorId"`
	Priority   *int   `json:"priority"`
	Status     *int   `json:"status"`
	StartDate  string `json:"startDate"`
	EndDate    string `json:"endDate"`
}
