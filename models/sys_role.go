package models

import (
	"time"

	"gorm.io/gorm"
)

type SysRole struct {
	CreatedAt   time.Time       // 创建时间
	UpdatedAt   time.Time       // 更新时间
	DeletedAt   gorm.DeletedAt  `gorm:"index" json:"-"`                                          // 删除时间
	RoleId      uint            `json:"role_id" gorm:"not null;unique;primary_key;comment:角色ID"` // 角色ID
	RoleName    string          `json:"authorityName" gorm:"comment:角色名"`                        // 角色名
	ParentId    *uint           `json:"parentId" gorm:"comment:父角色ID"`                           // 父角色ID
	Permissions []SysPermission `gorm:"many2many:sys_role_permissions;comment:角色权限"`             // 角色权限
}

func (SysRole) TableName() string {
	return "sys_roles"
}
