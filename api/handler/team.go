package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"toggler.in/api/database"
	"toggler.in/api/model"
)

func CreateTeam (c *fiber.Ctx) error {
	type RequestBody struct {
		UserId string `json:"userId" validate:"required"`
		Name string `json:"name" validate:"required"`
	}

	req := new(RequestBody)


	if err := c.BodyParser(req); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	db := database.DB
	user := model.User{}
	user.ID = req.UserId

	db.Limit(1).Find(&user)

	if user.Email == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "No logged in", "isLoggedIn": false})
	}

	fmt.Println(user)

	team := new(model.Team)
	team.Name = req.Name

	team.Members = append(team.Members, model.Member{
		User: user,
		Role: "admin",
	})

	return c.JSON(fiber.Map{"status": "success", "message": "Team created", "data": team})
}