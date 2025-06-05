package handler

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/middleware"
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/service"
)

// 租户相关API
func RegisterTenantRoutes(app *fiber.App) {
	r := app.Group("/api/tenant")
	// Protected tenant routes require authentication
	r.Use(middleware.AuthMiddleware)
	r.Get("/profile", tenantProfile)
	r.Put("/profile", tenantUpdate)
	r.Get("/list", tenantList)
}

func tenantProfile(c *fiber.Ctx) error {
	// 获取当前租户ID
	tid := c.Locals("tenantId").(int)
	tenant, err := service.TenantService.GetByID(uint(tid))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"code": 404, "message": "租户不存在"})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": tenant})
}

func tenantUpdate(c *fiber.Ctx) error {
	// 请求包含租户基本信息和定制化信息
	type UpdateRequest struct {
		Name          string `json:"name"`
		Logo          string `json:"logo"`
		Theme         string `json:"theme"`
		WelcomeText   string `json:"welcomeText"`
		Customization struct {
			Logo        string `json:"logo"`
			SiteName    string `json:"siteName"`
			ThemeColor  string `json:"themeColor"`
			Favicon     string `json:"favicon"`
			ExtraConfig string `json:"extraConfig"`
		} `json:"customization"`
	}
	var req UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "参数错误"})
	}
	// 获取当前租户ID，并更新租户基本信息
	tid := c.Locals("tenantId").(int)
	tenant := model.Tenant{
		ID:          uint(tid),
		Name:        req.Name,
		Logo:        req.Logo,
		Theme:       req.Theme,
		WelcomeText: req.WelcomeText,
	}
	if err := service.TenantService.Update(&tenant); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error()})
	}
	// 更新定制化信息
	// 默认 ExtraConfig 为 {} 以保证 JSON 有效
	extraCfg := strings.TrimSpace(req.Customization.ExtraConfig)
	if extraCfg == "" {
		extraCfg = "{}"
	}
	tc := model.TenantCustomization{
		TenantID:    uint(tid),
		Logo:        req.Customization.Logo,
		SiteName:    req.Customization.SiteName,
		ThemeColor:  req.Customization.ThemeColor,
		Favicon:     req.Customization.Favicon,
		ExtraConfig: extraCfg,
	}
	if err := service.TenantCustomizationService.Update(&tc); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": "更新定制化失败: " + err.Error()})
	}
	// 返回最新完整租户信息
	updated, err := service.TenantService.GetByID(uint(tid))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": "获取更新后租户失败: " + err.Error()})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "更新成功", "data": updated})
}

func tenantList(c *fiber.Ctx) error {
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	list, total, err := service.TenantService.List(offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": fiber.Map{"list": list, "total": total}})
}
