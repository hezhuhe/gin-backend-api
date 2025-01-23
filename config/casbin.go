package config

import (
	"gin-backend-api/global"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func InitCasbin() {
	// 定义 Casbin 的模型
	text := `
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
	`

	// 创建 Casbin 模型
	m, err := model.NewModelFromString(text)
	if err != nil {
		log.Fatal("Failed to create Casbin model:", err)
	}

	// 使用已初始化的数据库连接
	a, err := gormadapter.NewAdapterByDB(global.Db)
	if err != nil {
		log.Fatal("Failed to create Casbin adapter:", err)
	}

	// 自动迁移 Casbin 表结构
	err = global.Db.AutoMigrate(&gormadapter.CasbinRule{})
	if err != nil {
		log.Fatal("Failed to migrate Casbin table:", err)
	}

	// 创建 Casbin Enforcer
	enforcer, err := casbin.NewEnforcer(m, a)
	if err != nil {
		log.Fatal("Failed to create Casbin enforcer:", err)
	}

	// 将 Casbin Enforcer 存储到全局变量
	global.Enforcer = enforcer
	log.Println("Casbin initialized successfully")

	log.Println("Database initialization completed!")
}
