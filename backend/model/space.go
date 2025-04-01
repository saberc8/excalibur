package model

import "aquila/global"

const TableNameSpaceEntity = "space"

type SpaceEntity struct {
	global.GModel
	Name        string `gorm:"column:name;type:varchar(100);not null;comment:空间名称" json:"name"`
	Description string `gorm:"column:description;type:text;comment:空间描述" json:"description"`
	Path        string `gorm:"column:path;type:varchar(255);not null;unique;comment:空间路径" json:"path"`
	OwnerId     int64  `gorm:"column:owner_id;type:bigint;not null;comment:空间拥有者ID" json:"ownerId"`
	Status      int    `gorm:"column:status;type:smallint;default:0;comment:状态(0:正常,1:禁用)" json:"status"`
}

func (*SpaceEntity) TableName() string {
	return TableNameSpaceEntity
}
