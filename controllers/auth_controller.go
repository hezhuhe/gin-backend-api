package controllers

import (
	"fmt"
	"gin-backend-api/global"
	"gin-backend-api/models"
	"gin-backend-api/utils"
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	var user models.Sys_User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hashedPwd, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

// @Summary 用户登录
// @Description 用户通过账号密码登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Router /api/auth/login [post]
// @Success 200 {object} utils.Response{data=string} "登录成功，返回token"
// @Failure 400 {object} utils.Response "请求参数错误"
// @Failure 401 {object} utils.Response "用户名或密码错误"
// @Failure 500 {object} utils.Response "服务器内部错误"
func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// 绑定输入参数
	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.Error(ctx, http.StatusBadRequest, err.Error()) // 使用统一的错误返回
		return
	}

	var user models.Sys_User
	// 查询用户
	if err := global.Db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		utils.Error(ctx, http.StatusUnauthorized, "错误的用户名或密码") // 错误的用户名或密码
		return
	}

	// 验证密码
	if !utils.CheckPassword(input.Password, user.Password) {
		utils.Error(ctx, http.StatusBadRequest, "错误的密码") // 密码错误
		return
	}

	// 生成JWT
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		utils.Error(ctx, http.StatusInternalServerError, err.Error()) // 生成JWT失败
		return
	}

	// 成功返回Token
	utils.Success(ctx, gin.H{"token": token}, "登录成功")
}

// 初始化管理员及相关数据
func initAdminData(db *gorm.DB, enforcer *casbin.Enforcer) error {
	// 检查管理员用户是否已存在
	var adminUser models.Sys_User
	if err := db.Where("username = ?", "admin").First(&adminUser).Error; err == nil {
		log.Println("Admin user already exists, skipping initialization")
		return nil
	} else if err != gorm.ErrRecordNotFound {
		return fmt.Errorf("error checking admin user existence: %v", err)
	}

	log.Println("Admin user does not exist, initializing data...")

	// 1. 初始化权限
	permissions := getDefaultPermissions()
	if err := db.Create(&permissions).Error; err != nil {
		return fmt.Errorf("error creating permissions: %v", err)
	}

	// 2. 初始化管理员角色
	adminRole := models.SysRole{
		RoleName:    "管理员",
		Permissions: permissions, // 赋予管理员所有权限
	}
	if err := db.Create(&adminRole).Error; err != nil {
		return fmt.Errorf("error creating admin role: %v", err)
	}

	hashedPwd, err := utils.HashPassword("123456")

	if err != nil {
		return fmt.Errorf("error hash password: %v", err)
	}
	// 3. 初始化管理员用户
	adminUser = models.Sys_User{
		Username: "admin",
		Password: hashedPwd, // 注意：实际项目中请加密密码
		NickName: "超级管理员",
		Roles:    []models.SysRole{adminRole}, // 绑定管理员角色
	}
	if err := db.Create(&adminUser).Error; err != nil {
		return fmt.Errorf("error creating admin user: %v", err)
	}

	log.Println("Admin user, role, and permissions initialized successfully!")

	// 4. 初始化 Casbin 策略
	// 为管理员角色配置权限
	for _, permission := range permissions {
		// 在 Casbin 中添加角色 -> 权限 策略
		if _, err := enforcer.AddPolicy("admin", permission.Path, permission.Method); err != nil {
			return fmt.Errorf("error adding Casbin policy for permission: %v", err)
		}
	}

	// 5. 为用户绑定角色策略
	if _, err := enforcer.AddRoleForUser("admin", "管理员"); err != nil {
		return fmt.Errorf("error adding role policy for user: %v", err)
	}

	log.Println("Casbin policies and user-role bindings initialized successfully!")
	return nil
}

// 获取默认权限列表
func getDefaultPermissions() []models.SysPermission {
	return []models.SysPermission{
		{Name: "user_create", Describe: "创建用户", Path: "/users", Method: "POST"},
		{Name: "user_update", Describe: "更新用户", Path: "/users", Method: "POST"},
		{Name: "user_delete", Describe: "删除用户", Path: "/users", Method: "DELETE"},
		{Name: "user_read", Describe: "查询用户", Path: "/users", Method: "GET"},
		{Name: "role_create", Describe: "创建角色", Path: "/roles", Method: "POST"},
		{Name: "role_update", Describe: "更新角色", Path: "/roles", Method: "PUT"},
		{Name: "role_delete", Describe: "删除角色", Path: "/roles", Method: "DELETE"},
		{Name: "role_read", Describe: "查询角色", Path: "/roles", Method: "GET"},
		{Name: "permission_create", Describe: "创建权限", Path: "/permissions", Method: "POST"},
		{Name: "permission_update", Describe: "更新权限", Path: "/permissions", Method: "PUT"},
		{Name: "permission_delete", Describe: "删除权限", Path: "/permissions", Method: "DELETE"},
		{Name: "permission_read", Describe: "查询权限", Path: "/permissions", Method: "GET"},
	}
}
