package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/middleware"
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/service"
)

// ListUserPermissions returns the list of permission codes assigned to the current user
func listUserPermissions(c *fiber.Ctx) error {
	uid := c.Locals("userId").(int)
	codes, err := service.RelationService.ListUserPermissions(uint(uid))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": codes})
}

// 关系相关API
func RegisterRelationRoutes(router fiber.Router) {
	r := router.Group("/relation")
	r.Use(middleware.AuthMiddleware)
	r.Post("/user-role", middleware.RequirePermission("relation:user-role:assign"), assignRoleToUser)
	r.Delete("/user-role", middleware.RequirePermission("relation:user-role:remove"), removeRoleFromUser)
	r.Get("/user-roles", middleware.RequirePermission("relation:user-role:list"), listUserRoles)

	r.Post("/role-permission", middleware.RequirePermission("relation:role-permission:assign"), assignPermissionToRole)
	r.Delete("/role-permission", middleware.RequirePermission("relation:role-permission:remove"), removePermissionFromRole)
	r.Get("/role-permissions", middleware.RequirePermission("relation:role-permission:list"), listRolePermissions)
	r.Get("/user-permissions", listUserPermissions)
}

func assignRoleToUser(c *fiber.Ctx) error {
	type Req struct {
		UserID uint `json:"userId"`
		RoleID uint `json:"roleId"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 校验角色和用户归属租户
	tid := c.Locals("tenantId").(int)
	role, _ := service.RoleService.GetByID(req.RoleID)
	user, _ := service.UserService.GetByID(req.UserID)
	if role == nil || role.TenantID != uint(tid) || user == nil || user.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法用户或角色", "data": nil})
	}
	err := service.RelationService.AssignRoleToUser(req.UserID, req.RoleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	tid = c.Locals("tenantId").(int)
	uid := c.Locals("userId").(int)
	log := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "relation:user-role:assign", Detail: fmt.Sprintf("Assigned role %d to user %d", req.RoleID, req.UserID)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "分配成功", "data": nil})
}

func removeRoleFromUser(c *fiber.Ctx) error {
	type Req struct {
		UserID uint `json:"userId"`
		RoleID uint `json:"roleId"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 校验角色和用户归属租户
	tid := c.Locals("tenantId").(int)
	role, _ := service.RoleService.GetByID(req.RoleID)
	user, _ := service.UserService.GetByID(req.UserID)
	if role == nil || role.TenantID != uint(tid) || user == nil || user.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法用户或角色", "data": nil})
	}
	err := service.RelationService.RemoveRoleFromUser(req.UserID, req.RoleID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	tid = c.Locals("tenantId").(int)
	uid := c.Locals("userId").(int)
	log := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "relation:user-role:remove", Detail: fmt.Sprintf("Removed role %d from user %d", req.RoleID, req.UserID)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "移除成功", "data": nil})
}

func listUserRoles(c *fiber.Ctx) error {
	// support listing roles for any user via query param 'userId', else current user
	uid := c.Locals("userId").(int)
	if q := c.Query("userId"); q != "" {
		if id, err := strconv.Atoi(q); err == nil {
			uid = id
		}
	}
	// 校验用户归属租户
	tid := c.Locals("tenantId").(int)
	user, err := service.UserService.GetByID(uint(uid))
	if err != nil || user == nil || user.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法用户", "data": nil})
	}
	roles, err := service.RelationService.ListRolesByUser(uint(uid))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": roles})
}

func assignPermissionToRole(c *fiber.Ctx) error {
	type Req struct {
		RoleID       uint `json:"roleId"`
		PermissionID uint `json:"permissionId"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 校验角色归属租户
	tid := c.Locals("tenantId").(int)
	role, err := service.RoleService.GetByID(req.RoleID)
	if err != nil || role == nil || role.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法角色", "data": nil})
	}
	err = service.RelationService.AssignPermissionToRole(req.RoleID, req.PermissionID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	tid = c.Locals("tenantId").(int)
	uid := c.Locals("userId").(int)
	log := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "relation:role-permission:assign", Detail: fmt.Sprintf("Assigned permission %%d to role %%d", req.PermissionID, req.RoleID)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "分配成功", "data": nil})
}

func removePermissionFromRole(c *fiber.Ctx) error {
	type Req struct {
		RoleID       uint `json:"roleId"`
		PermissionID uint `json:"permissionId"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 校验角色归属租户
	tid := c.Locals("tenantId").(int)
	role, err := service.RoleService.GetByID(req.RoleID)
	if err != nil || role == nil || role.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法角色", "data": nil})
	}
	err = service.RelationService.RemovePermissionFromRole(req.RoleID, req.PermissionID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	tid = c.Locals("tenantId").(int)
	uid := c.Locals("userId").(int)
	log := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "relation:role-permission:remove", Detail: fmt.Sprintf("Removed permission %%d from role %%d", req.PermissionID, req.RoleID)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "移除成功", "data": nil})
}

func listRolePermissions(c *fiber.Ctx) error {
	ridStr := c.Query("roleId")
	rid, _ := strconv.ParseUint(ridStr, 10, 64)
	// 校验角色归属租户
	tid := c.Locals("tenantId").(int)
	role, err := service.RoleService.GetByID(uint(rid))
	if err != nil || role == nil || role.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法角色", "data": nil})
	}
	perms, err := service.RelationService.ListPermissionsByRole(uint(rid))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": perms})
}
