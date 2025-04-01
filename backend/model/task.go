package model

import "aquila/global"

const TableNameTaskEntity = "task"

type TaskEntity struct {
	global.GModel
	Title      string `gorm:"column:title;type:varchar(200);not null;comment:任务标题" json:"title"`
	Content    string `gorm:"column:content;type:text;comment:任务内容" json:"content"`
	ProjectId  int64  `gorm:"column:project_id;type:bigint;not null;comment:所属项目ID" json:"projectId"`
	AssigneeId int64  `gorm:"column:assignee_id;type:bigint;comment:执行人ID" json:"assigneeId"`
	CreatorId  int64  `gorm:"column:creator_id;type:bigint;not null;comment:创建人ID" json:"creatorId"`
	Priority   int    `gorm:"column:priority;type:smallint;default:0;comment:优先级(0:普通,1:重要,2:紧急)" json:"priority"`
	Status     int    `gorm:"column:status;type:smallint;default:0;comment:状态(0:待处理,1:进行中,2:已完成)" json:"status"`
	StartDate  string `gorm:"column:start_date;type:date;comment:开始日期" json:"startDate"`
	EndDate    string `gorm:"column:end_date;type:date;comment:结束日期" json:"endDate"`
}

func (*TaskEntity) TableName() string {
	return TableNameTaskEntity
}
