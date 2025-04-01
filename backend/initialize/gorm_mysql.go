package initialize

import (
	"aquila/global"
	"aquila/model"
	"fmt"
	"go.uber.org/zap"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {
	m := global.AquilaConfig.MySQL
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.DSN(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}

	db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config(m.Prefix, m.Singular))
	if err != nil {
		global.AquilaLog.Error("MySQL启动异常", zap.Any("err", err))
		return nil
	} else {

		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		// 打印数据库连接信息
		fmt.Println("====3-gorm====: gorm link MySQL success")
		global.AquilaLog.Info("MySQL启动成功", zap.String("host", m.Host), zap.String("port", m.Port), zap.String("dbname", m.Dbname))
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)

		// 创建表
		err := db.AutoMigrate(
			&model.UserEntity{},
			&model.RoleEntity{},
			&model.MenuEntity{},
			&model.RoleMenuEntity{},
			&model.UserRoleEntity{},
		)
		if err != nil {
			log.Fatalf("failed to auto migrate models: %v", err)
		}
		return db
	}

}
