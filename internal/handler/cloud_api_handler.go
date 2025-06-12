package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/middleware"
	"github.com/yockii/Tianshu/pkg/config"
)

func RegisterCloudAPIRoutes(router fiber.Router) {
	r := router.Group("/sys")
	r.Get("/cloud-api-info", getCloudAPIInfo)
	r.Use(middleware.AuthMiddleware)
	r.Get("/connect-info", getConnectInfo)
}

func getCloudAPIInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "获取云API信息成功",
		"data": fiber.Map{
			"appId":      config.Cfg.Dji.AppId,
			"appKey":     config.Cfg.Dji.AppKey,
			"appLicense": config.Cfg.Dji.AppLicense,
		},
	})
}

func getConnectInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "获取连接信息成功",
		"data": fiber.Map{
			"mqttTcpAddr":  fmt.Sprintf(":%d", config.Cfg.MQTT.TcpAddr),
			"mqttWsAddr":   fmt.Sprintf(":%d", config.Cfg.MQTT.WsAddr),
			"mqttUsername": "dji",
			"mqttPassword": "dji",
			"wsAddr":       fmt.Sprintf(":%d/ws", config.Cfg.Server.Port),
		},
	})
}
