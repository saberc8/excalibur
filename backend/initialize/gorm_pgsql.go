package initialize

import (
	"aquila/global"
	"aquila/model"
	"fmt"
	"log"

	"go.uber.org/zap"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GormPgSql 初始化 Postgresql 数据库
func GormPgSql() *gorm.DB {
	p := global.AquilaConfig.PGSQL

	if p.Dbname == "" {
		return nil
	}

	pgsqlConfig := postgres.Config{
		DSN:                  p.DSN(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	fmt.Printf("====3-gorm====: p.DSN()=%v\n", p.DSN())
	db, err := gorm.Open(postgres.New(pgsqlConfig), Gorm.Config(p.Prefix, p.Singular))

	if err != nil {
		fmt.Printf("====3-gorm====: gorm link PostgreSQL failed, err: %v\n", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)

		fmt.Println("====3-gorm====: gorm link PostgreSQL success")
		// 打印数据库连接信息
		global.AquilaLog.Info("PostgreSQL启动成功", zap.String("host", p.Host), zap.String("port", p.Port), zap.String("dbname", p.Dbname))

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
