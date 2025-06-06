// filepath: internal/middleware/authz_middleware.go
package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/service"
)

// RequirePermission returns a middleware that checks if the current user has the specified permission code
func RequirePermission(code string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// retrieve userId from context
		uidVal := c.Locals("userId")
		if uidVal == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"code": 401, "message": "未登录", "data": nil})
		}
		uid, ok := uidVal.(int)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"code": 401, "message": "用户ID错误", "data": nil})
		}
		allowed, err := service.RelationService.CheckUserPermission(uint(uid), code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": 500, "message": err.Error(), "data": nil})
		}
		if !allowed {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"code": 403, "message": "无权限", "data": nil})
		}
		return c.Next()
	}
}
