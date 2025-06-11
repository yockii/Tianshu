package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/pkg/config"
)

func RegisterCloudAPIRoutes(router fiber.Router) {
	router.Get("/sys/cloud-api-info", getCloudAPIInfo)
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
