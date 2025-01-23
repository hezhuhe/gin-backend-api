package models

import "gin-backend-api/global"

type SysPermission struct {
	global.BASE_MODEL
	Name     string    `json:"name" gorm:"index;comment:权限名"`               // 权限名
	Describe string    `json:"describe" gorm:"comment:权限描述"`                // 权限描述
	Path     string    `json:"paht" gorm:"comment:api路经"`                   // api路经
	Method   string    `json:"method" gorm:"comment:http方法"`                // http方法
	Roles    []SysRole `gorm:"many2many:sys_role_permissions;comment:关联角色"` // 关联角色
}

func (SysPermission) TableName() string {
	return "sys_permissions"
}
