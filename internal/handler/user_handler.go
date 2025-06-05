package handler

import (
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/yockii/Tianshu/internal/middleware"
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/service"
	"github.com/yockii/Tianshu/pkg/cache"
	"github.com/yockii/Tianshu/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

// 用户相关API
func RegisterUserRoutes(app *fiber.App) {
	r := app.Group("/api/user")
	r.Post("/register", userRegister)
	r.Post("/login", userLogin)
	// protected routes
	r.Use(middleware.AuthMiddleware)
	r.Get("/profile", userProfile)
	r.Put("/profile", userUpdate)
	r.Get("/list", userList)
}

func userRegister(c *fiber.Ctx) error {
	// 新注册即为租户注册+管理员注册
	type RegisterRequest struct {
		TenantName    string `json:"tenantName"`
		Domain        string `json:"domain"`
		Logo          string `json:"logo"`
		Theme         string `json:"theme"`
		WelcomeText   string `json:"welcomeText"`
		AdminUsername string `json:"adminUsername"`
		AdminEmail    string `json:"adminEmail"`
		AdminPassword string `json:"adminPassword"`
		AdminPhone    string `json:"adminPhone"`
	}
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 1. 创建租户
	tenant := model.Tenant{
		Name:        req.TenantName,
		Domain:      req.Domain,
		Logo:        req.Logo,
		Theme:       req.Theme,
		WelcomeText: req.WelcomeText,
	}
	if err := service.TenantService.Create(&tenant); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": "租户创建失败: " + err.Error(), "data": nil})
	}
	// 2. 创建管理员用户 with hashed password
	passHash, err := bcrypt.GenerateFromPassword([]byte(req.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "密码加密失败"})
	}
	admin := model.User{
		TenantID:     tenant.ID,
		Username:     req.AdminUsername,
		Email:        req.AdminEmail,
		Phone:        req.AdminPhone,
		PasswordHash: string(passHash),
		IsSuperAdmin: true,
		Status:       1,
	}
	if err := service.UserService.Create(&admin); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": "管理员创建失败: " + err.Error(), "data": nil})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "注册成功", "data": fiber.Map{"tenant": tenant, "admin": admin}})
}

func userLogin(c *fiber.Ctx) error {
	// 登录逻辑：域名或tenantId + 用户名 + 密码
	type LoginRequest struct {
		Domain   string `json:"domain"`
		TenantId uint   `json:"tenantId"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误"})
	}
	var ten *model.Tenant
	var err error
	if req.Domain != "" {
		// 优先使用域名
		ten, err = service.TenantService.GetByDomain(req.Domain)
	} else if req.TenantId != 0 {
		ten, err = service.TenantService.GetByID(req.TenantId)
	} else {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "请提供租户域名或ID"})
	}
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"code": 404, "message": "租户不存在"})
	}
	// 2. 根据用户名查用户
	user, err := service.UserService.GetByUsername(ten.ID, req.Username)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"code": 404, "message": "用户不存在"})
	}
	// 3. 校验密码
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"code": 401, "message": "用户名或密码错误"})
	}
	// 4. 生成 Redis sessionKey 并存储，返回包含 sessionKey 的 JWT
	// 4.1 生成 sessionKey
	sessKey := "session:" + uuid.New().String()
	conn := cache.Pool.Get()
	defer conn.Close()
	// 存储 tenantId 和 userId
	if _, err := conn.Do("HMSET", sessKey, "tenantId", ten.ID, "userId", user.ID); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": "Redis写入失败"})
	}
	// 设置过期时间（秒）
	ttl := int(time.Duration(config.Cfg.JWT.ExpireHours) * time.Hour / time.Second)
	if _, err := conn.Do("EXPIRE", sessKey, ttl); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": "Redis设置过期失败"})
	}
	// 4.2 生成 JWT
	secret := config.Cfg.JWT.Secret
	claims := jwt.MapClaims{
		"sessionKey": sessKey,
		"exp":        time.Now().Add(time.Hour * time.Duration(config.Cfg.JWT.ExpireHours)).Unix(),
	}
	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": "token生成失败"})
	}
	// 返回包含用户数据和 token
	return c.JSON(fiber.Map{"code": 200, "message": "登录成功", "data": fiber.Map{"token": tokenStr, "user": user}})
}

func userProfile(c *fiber.Ctx) error {
	// 已由中间件注入 userId
	uid := c.Locals("userId").(int)
	user, err := service.UserService.GetByID(uint(uid))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "用户不存在"})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": user})
}

func userUpdate(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "参数错误"})
	}
	// Override IDs from context to prevent tampering
	uid := c.Locals("userId").(int)
	user.ID = uint(uid)
	tid := c.Locals("tenantId").(int)
	user.TenantID = uint(tid)
	// 如果请求中包含 PasswordHash 字段，则对其进行哈希
	if user.PasswordHash != "" {
		if hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost); err == nil {
			user.PasswordHash = string(hash)
		}
	}
	if err := service.UserService.Update(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "更新成功", "data": user})
}

func userList(c *fiber.Ctx) error {
	// Use tenantId from context instead of query
	tid := c.Locals("tenantId").(int)
	offset, _ := strconv.Atoi(c.Query("offset", "0"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	list, total, err := service.UserService.List(uint(tid), offset, limit)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"code": 200, "message": "查询成功", "data": fiber.Map{"list": list, "total": total}})
}
