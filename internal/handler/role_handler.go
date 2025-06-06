package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/middleware"
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/service"
)

// 角色相关API
func RegisterRoleRoutes(app *fiber.App) {
	r := app.Group("/api/role")
	r.Use(middleware.AuthMiddleware)
	r.Post("", middleware.RequirePermission("role:create"), createRole)
	r.Get("/list", middleware.RequirePermission("role:list"), listRoles)
	r.Put("/:id", middleware.RequirePermission("role:update"), updateRole) // 添加更新路由
	r.Delete("/:id", middleware.RequirePermission("role:delete"), deleteRole)
}

func createRole(c *fiber.Ctx) error {
	type Req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		IsDefault   bool   `json:"isDefault"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 获取租户ID
	tid := c.Locals("tenantId").(int)
	// 默认角色标记前清除其他默认
	if req.IsDefault {
		_ = service.RoleService.UnsetDefaultRoles(uint(tid))
	}
	role := model.Role{TenantID: uint(tid), Name: req.Name, Description: req.Description, IsDefault: req.IsDefault}
	if err := service.RoleService.Create(&role); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	uid := c.Locals("userId").(int)
	log := model.OperationLog{TenantID: role.TenantID, UserID: uint(uid), Action: "role:create", Detail: fmt.Sprintf("Created role ID:%d", role.ID)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "创建成功", "data": role})
}

// updateRole 更新角色属性，包括默认标记处理
func updateRole(c *fiber.Ctx) error {
	// 路径参数
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 校验角色归属租户
	tid := c.Locals("tenantId").(int)
	existing, err := service.RoleService.GetByID(uint(id))
	if err != nil || existing == nil || existing.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法角色", "data": nil})
	}
	type Req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		IsDefault   bool   `json:"isDefault"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 处理默认角色
	if req.IsDefault {
		_ = service.RoleService.UnsetDefaultRoles(uint(tid))
	}
	role := model.Role{ID: uint(id), TenantID: uint(tid), Name: req.Name, Description: req.Description, IsDefault: req.IsDefault}
	if err := service.RoleService.Update(&role); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	uid := c.Locals("userId").(int)
	logEntry := model.OperationLog{TenantID: role.TenantID, UserID: uint(uid), Action: "role:update", Detail: fmt.Sprintf("Updated role ID:%d", role.ID)}
	service.OperationLogService.Create(&logEntry)
	return c.JSON(fiber.Map{"code": 200, "message": "更新成功", "data": role})
}

func listRoles(c *fiber.Ctx) error {
	tid := c.Locals("tenantId").(int)
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	roles, total, err := service.RoleService.List(uint(tid), offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": fiber.Map{"list": roles, "total": total}})
}

func deleteRole(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 校验角色归属租户
	tid := c.Locals("tenantId").(int)
	existing, err := service.RoleService.GetByID(uint(id))
	if err != nil || existing == nil || existing.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法角色", "data": nil})
	}
	if err := service.RoleService.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	uid := c.Locals("userId").(int)
	log := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "role:delete", Detail: fmt.Sprintf("Deleted role ID:%d", id)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "删除成功", "data": nil})
}
