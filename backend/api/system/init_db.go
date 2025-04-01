package system

import (
	"aquila/global"
	"aquila/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type InitDb struct{}

type InitDto struct {
	Force bool `json:"force"` // 是否强制初始化
}

const (
	// PostgreSQL语法
	userSQLPG = `INSERT INTO "user" ("id", "created_at", "updated_at", "deleted_at", "username", "password", "user_type", "mobile", "sort", "status", "last_login_ip", "last_login_nation", "last_login_province", "last_login_city", "last_login_date", "salt", "email", "avatar", "nickname") VALUES (1, '2024-09-04 21:22:07.312112+08', '2024-09-12 13:58:03.917271+08', NULL, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 0, '', 1, 0, '', '', '', '', '2024-09-04 21:22:07.31203', '', '', '', '111');`

	roleSQLPG = `INSERT INTO "role" ("id", "created_at", "updated_at", "deleted_at", "name", "remark", "status") 
		VALUES (1, '2024-09-10 23:20:44.683443+08', '2024-09-10 23:20:44.683443+08', NULL, 'admin', '', 0),
               (2, '2024-09-10 23:21:45.011993+08', '2024-09-10 23:21:45.011993+08', NULL, 'test1', '', 0),
               (3, '2024-09-10 23:21:48.373101+08', '2024-09-10 23:21:48.373101+08', NULL, 'test2', '', 0),
               (4, '2024-09-11 10:52:50.96456+08', '2024-09-11 10:52:50.96456+08', NULL, 'test3', '', 0),
               (5, '2024-09-12 14:37:09.868964+08', '2024-09-12 14:38:58.992348+08', NULL, 'role1', '', 0);`

	menuSQLPG = `INSERT INTO "menu" ("id", "created_at", "updated_at", "deleted_at", "name", "parent_id", "order_num", "path", "component", "query", "is_frame", "menu_type", "is_catch", "is_hidden", "perms", "icon", "status", "remark") 
		VALUES (1, '2024-09-11 10:03:48.465984+08', '2024-09-11 10:03:48.465984+08', NULL, '系统管理', 0, 0, '/system', 'Layout', '', 1, 'C', 0, 0, '', 'system', 0, ''),
               (2, '2024-09-11 10:12:11.175499+08', '2024-09-11 10:12:11.175499+08', NULL, '用户管理', 1, 1, 'user', 'system/user/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (3, '2024-09-11 10:23:24.565145+08', '2024-09-11 10:23:24.565145+08', NULL, '菜单管理', 1, 3, 'menu', 'system/menu/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (4, '2024-09-11 10:23:41.717165+08', '2024-09-11 10:23:41.717165+08', NULL, '角色管理', 1, 2, 'role', 'system/role/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (5, '2024-09-12 16:27:44.598973+08', '2024-09-12 16:27:44.598973+08', NULL, '1', 0, 0, '', '', '', 0, '', 0, 0, '', '', 0, '');`

	roleMenuSQLPG = `INSERT INTO "role_menu" ("id", "created_at", "updated_at", "deleted_at", "role_id", "menu_id") 
		VALUES (22, '2024-09-11 12:03:39.364184+08', '2024-09-11 12:03:39.364184+08', NULL, 1, 1),
               (23, '2024-09-11 12:03:39.364184+08', '2024-09-11 12:03:39.364184+08', NULL, 1, 2),
               (24, '2024-09-11 12:03:39.364184+08', '2024-09-11 12:03:39.364184+08', NULL, 1, 3),
               (25, '2024-09-11 12:03:39.364184+08', '2024-09-11 12:03:39.364184+08', NULL, 1, 4);`

	userRoleSQLPG = `INSERT INTO "user_role" ("id", "created_at", "updated_at", "deleted_at", "user_id", "role_id") 
		VALUES (5, '2024-09-11 13:07:37.054493+08', '2024-09-11 13:07:37.054493+08', NULL, 1, 1),
               (6, '2024-09-11 13:07:37.054493+08', '2024-09-11 13:07:37.054493+08', NULL, 1, 2),
               (7, '2024-09-11 13:07:37.054493+08', '2024-09-11 13:07:37.054493+08', NULL, 1, 3),
               (8, '2024-09-11 13:07:37.054493+08', '2024-09-11 13:07:37.054493+08', NULL, 1, 4);`

	// MySQL语法
	userSQLMySQL = `INSERT INTO ` + "`user`" + ` (` + "`id`" + `, ` + "`created_at`" + `, ` + "`updated_at`" + `, ` + "`deleted_at`" + `, ` + "`username`" + `, ` + "`password`" + `, ` + "`user_type`" + `, ` + "`mobile`" + `, ` + "`sort`" + `, ` + "`status`" + `, ` + "`last_login_ip`" + `, ` + "`last_login_nation`" + `, ` + "`last_login_province`" + `, ` + "`last_login_city`" + `, ` + "`last_login_date`" + `, ` + "`salt`" + `, ` + "`email`" + `, ` + "`avatar`" + `, ` + "`nickname`" + `) VALUES (1, '2024-09-04 21:22:07', '2024-09-12 13:58:03', NULL, 'admin', 'e10adc3949ba59abbe56e057f20f883e', 0, '', 1, 0, '', '', '', '', '2024-09-04 21:22:07', '', '', '', '111');`

	roleSQLMySQL = `INSERT INTO ` + "`role`" + ` (` + "`id`" + `, ` + "`created_at`" + `, ` + "`updated_at`" + `, ` + "`deleted_at`" + `, ` + "`name`" + `, ` + "`remark`" + `, ` + "`status`" + `) 
		VALUES (1, '2024-09-10 23:20:44', '2024-09-10 23:20:44', NULL, 'admin', '', 0),
               (2, '2024-09-10 23:21:45', '2024-09-10 23:21:45', NULL, 'test1', '', 0),
               (3, '2024-09-10 23:21:48', '2024-09-10 23:21:48', NULL, 'test2', '', 0),
               (4, '2024-09-11 10:52:50', '2024-09-11 10:52:50', NULL, 'test3', '', 0),
               (5, '2024-09-12 14:37:09', '2024-09-12 14:38:58', NULL, 'role1', '', 0);`

	menuSQLMySQL = `INSERT INTO ` + "`menu`" + ` (` + "`id`" + `, ` + "`created_at`" + `, ` + "`updated_at`" + `, ` + "`deleted_at`" + `, ` + "`name`" + `, ` + "`parent_id`" + `, ` + "`order_num`" + `, ` + "`path`" + `, ` + "`component`" + `, ` + "`query`" + `, ` + "`is_frame`" + `, ` + "`menu_type`" + `, ` + "`is_catch`" + `, ` + "`is_hidden`" + `, ` + "`perms`" + `, ` + "`icon`" + `, ` + "`status`" + `, ` + "`remark`" + `) 
		VALUES (1, '2024-09-11 10:03:48', '2024-09-11 10:03:48', NULL, '系统管理', 0, 0, '/system', 'Layout', '', 1, 'C', 0, 0, '', 'system', 0, ''),
               (2, '2024-09-11 10:12:11', '2024-09-11 10:12:11', NULL, '用户管理', 1, 1, 'user', 'system/user/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (3, '2024-09-11 10:23:24', '2024-09-11 10:23:24', NULL, '菜单管理', 1, 3, 'menu', 'system/menu/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (4, '2024-09-11 10:23:41', '2024-09-11 10:23:41', NULL, '角色管理', 1, 2, 'role', 'system/role/index', '', 1, 'M', 0, 0, 'menu:page', '', 0, ''),
               (5, '2024-09-12 16:27:44', '2024-09-12 16:27:44', NULL, '1', 0, 0, '', '', '', 0, '', 0, 0, '', '', 0, '');`

	roleMenuSQLMySQL = `INSERT INTO ` + "`role_menu`" + ` (` + "`id`" + `, ` + "`created_at`" + `, ` + "`updated_at`" + `, ` + "`deleted_at`" + `, ` + "`role_id`" + `, ` + "`menu_id`" + `) 
		VALUES (22, '2024-09-11 12:03:39', '2024-09-11 12:03:39', NULL, 1, 1),
               (23, '2024-09-11 12:03:39', '2024-09-11 12:03:39', NULL, 1, 2),
               (24, '2024-09-11 12:03:39', '2024-09-11 12:03:39', NULL, 1, 3),
               (25, '2024-09-11 12:03:39', '2024-09-11 12:03:39', NULL, 1, 4);`

	userRoleSQLMySQL = `INSERT INTO ` + "`user_role`" + ` (` + "`id`" + `, ` + "`created_at`" + `, ` + "`updated_at`" + `, ` + "`deleted_at`" + `, ` + "`user_id`" + `, ` + "`role_id`" + `) 
		VALUES (5, '2024-09-11 13:07:37', '2024-09-11 13:07:37', NULL, 1, 1),
               (6, '2024-09-11 13:07:37', '2024-09-11 13:07:37', NULL, 1, 2),
               (7, '2024-09-11 13:07:37', '2024-09-11 13:07:37', NULL, 1, 3),
               (8, '2024-09-11 13:07:37', '2024-09-11 13:07:37', NULL, 1, 4);`
)

func (i InitDb) InitializeDBApi(ctx *gin.Context) {
	var req InitDto
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Fail(ctx, "参数绑定失败:"+err.Error())
		return
	}

	// 如果不是强制初始化,检查是否已经初始化过
	if !req.Force {
		var count int64
		global.AquilaDb.Table("user").Count(&count)
		if count > 0 {
			utils.Fail(ctx, "数据库已初始化")
			return
		}
	}

	// 开启事务
	tx := global.AquilaDb.Begin()

	// 确定当前使用的数据库类型
	dbType := global.AquilaConfig.App.DbType

	// 根据数据库类型选择不同的清空表语句
	var truncateStatements []string
	tables := []string{"user_role", "role_menu", "menu", "role", "user"}

	if req.Force {
		if dbType == "mysql" {
			for _, table := range tables {
				truncateStatements = append(truncateStatements, fmt.Sprintf("TRUNCATE TABLE `%s`", table))
			}
		} else if dbType == "pgsql" {
			for _, table := range tables {
				truncateStatements = append(truncateStatements, fmt.Sprintf("TRUNCATE TABLE \"%s\" RESTART IDENTITY CASCADE", table))
			}
		} else {
			tx.Rollback()
			utils.Fail(ctx, "不支持的数据库类型:"+dbType)
			return
		}

		// 执行清空表操作
		for _, sql := range truncateStatements {
			if err := tx.Exec(sql).Error; err != nil {
				tx.Rollback()
				utils.Fail(ctx, "清空表数据失败:"+err.Error())
				return
			}
		}
	}

	// 根据数据库类型选择对应的SQL语句
	var sqlStatements []string

	if dbType == "mysql" {
		sqlStatements = []string{
			userSQLMySQL,
			roleSQLMySQL,
			menuSQLMySQL,
			roleMenuSQLMySQL,
			userRoleSQLMySQL,
		}
	} else if dbType == "pgsql" {
		sqlStatements = []string{
			userSQLPG,
			roleSQLPG,
			menuSQLPG,
			roleMenuSQLPG,
			userRoleSQLPG,
		}
	} else {
		tx.Rollback()
		utils.Fail(ctx, "不支持的数据库类型:"+dbType)
		return
	}

	// 执行SQL语句
	for _, sql := range sqlStatements {
		if err := tx.Exec(sql).Error; err != nil {
			tx.Rollback()
			utils.Fail(ctx, "执行SQL失败:"+err.Error())
			return
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		utils.Fail(ctx, "提交事务失败:"+err.Error())
		return
	}

	utils.Success(ctx, nil)
}
