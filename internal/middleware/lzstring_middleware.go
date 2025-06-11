package middleware

import (
	lzstring "github.com/daku10/go-lz-string"
	"github.com/gofiber/fiber/v2"
	"github.com/yockii/Tianshu/internal/utils"
)

func LZStringMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		encryptedBody := c.Body()
		if len(encryptedBody) == 0 {
			return handleResponse(c) // 如果没有请求体，直接调用下一个中间件
		}

		// 如果头设置了不解密，则不处理解密
		if c.Get("X-Decrypt") != "false" {
			decryptedBody, err := lzstring.DecompressFromUTF16(utils.EncodeStringToUTF16(string(encryptedBody)))
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"code": 400, "message": "Decompression failed", "data": nil})
			}
			c.Request().SetBody([]byte(decryptedBody))
		}

		return handleResponse(c) // 如果没有请求体，直接调用下一个中间件
	}
}

func handleResponse(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return err
	}

	// 如果头设置了不加密，则不加密
	if c.Get("X-Encrypt") == "false" {
		return nil
	}

	originalBody := c.Response().Body()
	if len(originalBody) == 0 {
		return nil
	}

	encryptedBodyUint6, err := lzstring.CompressToUTF16((string(originalBody)))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": 500, "message": "Encrypt data failed", "data": nil})
	}

	encryptedBody := utils.DecodeUTF16ToString(encryptedBodyUint6)

	c.Response().SetBody([]byte(encryptedBody))
	return nil
}
