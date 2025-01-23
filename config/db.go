package config

import (
	"gin-backend-api/global"
	"gin-backend-api/models"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() {
	// 初始化数据库连接
	dsn := AppConfig.DataBase.DSN
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error configuring database connection pool: %v", err)
	}

	// 配置连接池参数
	sqlDB.SetMaxIdleConns(AppConfig.DataBase.MaxIdleConns)
	sqlDB.SetMaxOpenConns(AppConfig.DataBase.MaxOpenCons)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 初始化表结构
	if err := autoMigrateTables(db); err != nil {
		log.Fatalf("Error migrating tables: %v", err)
	}

	global.Db = db
	log.Println("MySQL initialized successfully")

}

// 自动迁移表结构
func autoMigrateTables(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Sys_User{},
		&models.SysRole{},
		&models.SysPermission{},
	)
}
