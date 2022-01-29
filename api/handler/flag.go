package handler

import (
	"github.com/gofiber/fiber/v2"
)

func CreateFlag () error	{
	return nil
}

func GetUserFlags (c *fiber.Ctx) error {
	type RequestBody struct {
		UserID string `json:"userId" validate:"required"`
	}

	body := new(RequestBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	return c.JSON(fiber.Map{"message": "ok", "id": body.UserID})
}