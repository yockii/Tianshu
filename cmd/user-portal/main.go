// 用户/租户端服务入口
package main

import (
	"errors"
	"log"
	"strconv"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/handler"
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/service"
	"github.com/yockii/Tianshu/pkg/cache"
	"github.com/yockii/Tianshu/pkg/config"
	"github.com/yockii/Tianshu/pkg/db"
)

func main() {
	if err := config.InitConfig(""); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	if err := db.InitDB(); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	if err := db.AutoMigrateModels(model.Models); err != nil {
		log.Fatalf("自动迁移表结构失败: %v", err)
	}
	cache.InitRedis()
	// 初始化权限数据
	initPermissions()

	app := fiber.New()

	handler.RegisterTenantRoutes(app)
	handler.RegisterUserRoutes(app)
	// Role, Permission, Relation and Log routes
	handler.RegisterRoleRoutes(app)
	handler.RegisterPermissionRoutes(app)
	handler.RegisterRelationRoutes(app)
	handler.RegisterLogRoutes(app)

	// Use configured server port
	portStr := strconv.Itoa(config.Cfg.Server.UserPortalPort)
	log.Printf("User Portal running at :%s", portStr)
	log.Fatal(app.Listen(":" + portStr))
}

// initPermissions seeds default permissions into database if not exists
func initPermissions() {
	// 权限列表
	perms := []model.Permission{
		{Code: "user:list", Description: "用户列表"},
		{Code: "user:create", Description: "创建用户"},
		{Code: "user:update", Description: "更新用户"},
		{Code: "user:delete", Description: "删除用户"},
		{Code: "role:create", Description: "创建角色"},
		{Code: "role:list", Description: "角色列表"},
		{Code: "role:delete", Description: "删除角色"},
		{Code: "permission:create", Description: "创建权限"},
		{Code: "permission:list", Description: "权限列表"},
		{Code: "relation:user-role:assign", Description: "分配用户角色"},
		{Code: "relation:user-role:remove", Description: "移除用户角色"},
		{Code: "relation:user-role:list", Description: "用户角色列表"},
		{Code: "relation:role-permission:assign", Description: "分配角色权限"},
		{Code: "relation:role-permission:remove", Description: "移除角色权限"},
		{Code: "relation:role-permission:list", Description: "角色权限列表"},
		{Code: "logs:list", Description: "操作日志列表"},
	}
	for _, p := range perms {
		_, err := service.PermissionService.GetByCode(p.Code)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				if err2 := service.PermissionService.Create(&p); err2 != nil {
					log.Printf("初始化权限 %s 失败: %v", p.Code, err2)
				}
			} else {
				log.Printf("查询权限 %s 出错: %v", p.Code, err)
			}
		}
	}
}
