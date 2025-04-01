package model

import "aquila/global"

const TableNameProjectEntity = "project"

type ProjectEntity struct {
	global.GModel
	Name        string `gorm:"column:name;type:varchar(100);not null;comment:项目名称" json:"name"`
	SpaceId     int64  `gorm:"column:space_id;type:bigint;not null;comment:所属空间ID" json:"spaceId"`
	Description string `gorm:"column:description;type:text;comment:项目描述" json:"description"`
	OwnerId     int64  `gorm:"column:owner_id;type:bigint;not null;comment:项目负责人ID" json:"ownerId"`
	Status      int    `gorm:"column:status;type:smallint;default:0;comment:状态(0:进行中,1:已完成,2:已关闭)" json:"status"`
}

func (*ProjectEntity) TableName() string {
	return TableNameProjectEntity
}
