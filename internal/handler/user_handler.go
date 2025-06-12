package handler

import (
	"fmt"
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
func RegisterUserRoutes(router fiber.Router) {
	r := router.Group("/user")
	r.Post("/register", userRegister)
	r.Post("/login", userLogin)
	// protected routes
	r.Use(middleware.AuthMiddleware)
	// self profile
	r.Get("/profile", userProfile)
	r.Put("/profile", userUpdate)
	// tenant-level user management
	r.Get("/list", middleware.RequirePermission("user:list"), userList)
	r.Post("/create", middleware.RequirePermission("user:create"), createUser)
	r.Put("/:id", middleware.RequirePermission("user:update"), updateUserByID)
	r.Delete("/:id", middleware.RequirePermission("user:delete"), deleteUser)
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
		TenantName string `json:"tenantName"`
		Username   string `json:"username"`
		Password   string `json:"password"`
	}
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误"})
	}
	var ten *model.Tenant
	var err error
	if req.TenantName == "" {
		// 获取域名
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "租户名称不能为空"})
	} else {
		ten, err = service.TenantService.GetByName(req.TenantName)
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
	// 获取并附加用户权限列表
	perms, _ := service.RelationService.ListUserPermissions(user.ID)
	return c.JSON(fiber.Map{"code": 200, "message": "登录成功", "data": fiber.Map{"token": tokenStr, "user": user, "permissions": perms}})
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

// createUser allows tenant admins to create users under their tenant
func createUser(c *fiber.Ctx) error {
	tid := c.Locals("tenantId").(int)
	uid := c.Locals("userId").(int)
	type Req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
		Status   int    `json:"status"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": "密码加密失败", "data": nil})
	}
	user := model.User{TenantID: uint(tid), Username: req.Username, Email: req.Email, Phone: req.Phone, PasswordHash: string(hash), Status: req.Status}
	if err := service.UserService.Create(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// 默认角色自动分配
	if defRole, err := service.RoleService.GetDefaultRole(uint(tid)); err == nil {
		_ = service.RelationService.AssignRoleToUser(user.ID, defRole.ID)
		// 记录默认角色分配日志
		logDef := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "user:assign_default_role", Detail: fmt.Sprintf("Assigned default role ID:%d to user ID:%d", defRole.ID, user.ID)}
		service.OperationLogService.Create(&logDef)
	}
	// write operation log
	log := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "user:create", Detail: fmt.Sprintf("Created user ID:%d", user.ID)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "创建成功", "data": user})
}

// updateUserByID updates any user within the same tenant
func updateUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 校验用户归属租户
	tid := c.Locals("tenantId").(int)
	existing, err2 := service.UserService.GetByID(uint(id))
	if err2 != nil || existing == nil || existing.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法用户", "data": nil})
	}
	var req model.User
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// enforce tenant
	req.ID = uint(id)
	req.TenantID = uint(tid)
	// hash new password if provided
	if req.PasswordHash != "" {
		if h, e := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), bcrypt.DefaultCost); e == nil {
			req.PasswordHash = string(h)
		}
	}
	if err := service.UserService.Update(&req); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	uid := c.Locals("userId").(int)
	tid = c.Locals("tenantId").(int)
	log := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "user:update", Detail: fmt.Sprintf("Updated user ID:%d", req.ID)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "更新成功", "data": req})
}

// deleteUser deletes a user within the same tenant
func deleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"code": 400, "message": "参数错误", "data": nil})
	}
	// 校验用户归属租户
	tid := c.Locals("tenantId").(int)
	existing, err2 := service.UserService.GetByID(uint(id))
	if err2 != nil || existing == nil || existing.TenantID != uint(tid) {
		return c.Status(403).JSON(fiber.Map{"code": 403, "message": "非法用户", "data": nil})
	}
	if err := service.UserService.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
	}
	// write operation log
	uid := c.Locals("userId").(int)
	log := model.OperationLog{TenantID: uint(tid), UserID: uint(uid), Action: "user:delete", Detail: fmt.Sprintf("Deleted user ID:%d", id)}
	service.OperationLogService.Create(&log)
	return c.JSON(fiber.Map{"code": 200, "message": "删除成功", "data": nil})
}
