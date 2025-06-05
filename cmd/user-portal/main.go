// 用户/租户端服务入口
package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/handler"
	"github.com/yockii/Tianshu/internal/model"
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

	app := fiber.New()

	handler.RegisterTenantRoutes(app)
	handler.RegisterUserRoutes(app)
	// TODO: 注册用户端专属路由

	// Use configured server port
	portStr := strconv.Itoa(config.Cfg.Server.UserPortalPort)
	log.Printf("User Portal running at :%s", portStr)
	log.Fatal(app.Listen(":" + portStr))
}
