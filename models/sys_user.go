package models

import (
	"gin-backend-api/global"
)

type Sys_User struct {
	global.BASE_MODEL
	Username  string    `json:"username" gorm:"index;comment:用户登录名"`                                                   // 用户登录名
	Password  string    `json:"-" gorm:"comment:用户登录密码"`                                                               // 用户登录密码
	NickName  string    `json:"nick_name" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	HeaderImg string    `json:"header_img" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	Roles     []SysRole `json:"roles" gorm:"many2many:sys_user_roles;comment:用户角色"`                                    // 用户角色
	Phone     string    `json:"phone" gorm:"comment:用户手机号"`                                                            // 用户手机号
	Email     string    `json:"email" gorm:"comment:用户邮箱"`                                                             // 用户邮箱
	Enable    int       `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`                                       // 用户状态
}

func (Sys_User) TableName() string {
	return "sys_users"
}
