package middleware

import (
	"fmt"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gomodule/redigo/redis"
	"github.com/yockii/Tianshu/pkg/cache"
	"github.com/yockii/Tianshu/pkg/config"
)

// AuthMiddleware validates JWT, stores tenantId and userId in Redis with a session key and attaches data to context
func AuthMiddleware(c *fiber.Ctx) error {
	// Extract token
	auth := c.Get("Authorization")
	if auth == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "缺少Authorization头"})
	}
	var tokenStr string
	fmt.Sscanf(auth, "Bearer %s", &tokenStr)

	// Parse JWT
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.JWT.Secret), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "无效的Token"})
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token Claims解析错误"})
	}

	// Get sessionKey from JWT claims
	sessionKey, ok := claims["sessionKey"].(string)
	if !ok || sessionKey == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "无效的SessionKey"})
	}
	// Fetch tenantId and userId from Redis
	conn := cache.Pool.Get()
	defer conn.Close()
	data, err := redis.StringMap(conn.Do("HGETALL", sessionKey))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Redis读取失败"})
	}
	tid, err := strconv.Atoi(data["tenantId"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "解析租户ID失败"})
	}
	uid, err := strconv.Atoi(data["userId"])
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "解析用户ID失败"})
	}
	// Attach to context locals
	c.Locals("sessionKey", sessionKey)
	c.Locals("tenantId", tid)
	c.Locals("userId", uid)

	return c.Next()
}
