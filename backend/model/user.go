package model

import (
	"aquila/global"
	"time"
)

const TableNameUserEntity = "user"

type UserEntity struct {
	global.GModel
	Username          string    `gorm:"column:username;type:varchar(50);comment:账号" json:"username"`                                                        // 账号
	Password          string    `gorm:"column:password;type:varchar(100);not null;comment:密码" json:"password"`                                              // 密码
	Avatar            string    `gorm:"column:avatar;type:varchar(100);comment:头像" json:"avatar"`                                                           // 头像
	Nickname          string    `gorm:"column:nickname;type:varchar(50);comment:昵称" json:"nickname"`                                                        // 昵称
	UserType          int64     `gorm:"column:user_type;type:smallint;comment:账号类型:0普通账号,1是超管" json:"userType"`                                             // 账号类型:0普通账号,1是超管
	Email             string    `gorm:"column:email;type:varchar(50);comment:电邮地址" json:"email"`                                                            // 电邮地址
	Mobile            string    `gorm:"column:mobile;type:varchar(30);comment:手机号码" json:"mobile"`                                                          // 手机号码
	Sort              int64     `gorm:"column:sort;type:int;default:1;comment:排序" json:"sort"`                                                              // 排序
	Status            int64     `gorm:"column:status;type:smallint;comment:状态0是正常,1是禁用" json:"status"`                                                      // 状态0是正常,1是禁用
	LastLoginIP       string    `gorm:"column:last_login_ip;type:varchar(30);comment:最后登录ip地址" json:"lastLoginIp"`                                          // 最后登录ip地址
	LastLoginNation   string    `gorm:"column:last_login_nation;type:varchar(100);comment:最后登录国家" json:"lastLoginNation"`                                   // 最后登录国家
	LastLoginProvince string    `gorm:"column:last_login_province;type:varchar(100);comment:最后登录省份" json:"lastLoginProvince"`                               // 最后登录省份
	LastLoginCity     string    `gorm:"column:last_login_city;type:varchar(100);comment:最后登录城市" json:"lastLoginCity"`                                       // 最后登录城市
	LastLoginDate     time.Time `gorm:"column:last_login_date;type:timestamp(6);not null;default:CURRENT_TIMESTAMP(6);comment:最后登录时间" json:"lastLoginDate"` // 最后登录时间
	Salt              string    `gorm:"column:salt;type:varchar(30);comment:密码盐" json:"salt"`                                                               // 密码盐
}

// TableName UserEntity's table name
func (*UserEntity) TableName() string {
	return TableNameUserEntity
}
