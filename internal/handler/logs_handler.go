package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/middleware"
	"github.com/yockii/Tianshu/internal/service"
)

// 操作日志相关API
func RegisterLogRoutes(app *fiber.App) {
	r := app.Group("/api/logs")
	r.Use(middleware.AuthMiddleware)
	r.Get("", middleware.RequirePermission("logs:list"), listLogs)
}

// listLogs already logs via operation logs elsewhere
func listLogs(c *fiber.Ctx) error {
	tid := c.Locals("tenantId").(int)
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	logs, total, err := service.OperationLogService.List(uint(tid), offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": fiber.Map{"list": logs, "total": total}})
}
