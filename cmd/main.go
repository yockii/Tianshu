package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/yockii/Tianshu/internal/handler"
	"github.com/yockii/Tianshu/internal/middleware"
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/mqtt"
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

	if config.Cfg.MQTT.UseEmbedded {
		// 启动内嵌 MQTT 客户端，连接大疆 MQTT 网关
		mqtt.Start()
		defer mqtt.Close()
	}

	app := fiber.New()

	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(middleware.LZStringMiddleware())

	// 注册路由
	// 用户端
	{
		userEndpoint := app.Group("/api/v1/")
		handler.RegisterCloudAPIRoutes(userEndpoint)

		handler.RegisterTenantRoutes(userEndpoint)
		handler.RegisterUserRoutes(userEndpoint)
		// Role, Permission, Relation and Log routes
		handler.RegisterRoleRoutes(userEndpoint)
		handler.RegisterPermissionRoutes(userEndpoint)
		handler.RegisterRelationRoutes(userEndpoint)
		handler.RegisterLogRoutes(userEndpoint)
	}

	portStr := strconv.Itoa(config.Cfg.Server.Port)
	log.Printf("Cockpit Dashboard running at :%s", portStr)
	log.Fatal(app.Listen(":" + portStr))
}
