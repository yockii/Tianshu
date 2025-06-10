package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/middleware"
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/service"
)

// Permission相关API
func RegisterPermissionRoutes(router fiber.Router) {
	r := router.Group("/permission")
	r.Use(middleware.AuthMiddleware)
	r.Post("", middleware.RequirePermission("permission:create"), createPermission)
	r.Get("/list", middleware.RequirePermission("permission:list"), listPermissions)
}

func createPermission(c *fiber.Ctx) error {
	type Req struct {
		Code        string `json:"code"`
		Description string `json:"description"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	perm := model.Permission{Code: req.Code, Description: req.Description}
	if err := service.PermissionService.Create(&perm); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	uid := c.Locals("userId").(int)
	log := model.OperationLog{TenantID: 0, UserID: uint(uid), Action: "permission:create", Detail: fmt.Sprintf("Created permission code:%s", perm.Code)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "创建成功", "data": perm})
}

func listPermissions(c *fiber.Ctx) error {
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	perms, total, err := service.PermissionService.List(offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": fiber.Map{"list": perms, "total": total}})
}
