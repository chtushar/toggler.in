package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"toggler.in/api/config"
	"toggler.in/api/database"
	"toggler.in/api/model"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func userExists(id string) bool {
	db := database.DB
	var user model.User
	user.ID = id
	db.Limit(1).Find(&user)

	return user.Email != ""
}

func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		ID string
		FirstName string
		LastName string
		Email string
	}

	db := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	err = attachCookie(c,&AuthCookieData{
		ID: user.ID,
		Email: user.Email,
	})

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	newUser := NewUser{
		ID: user.ID,
		Email:    user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
	}

	return c.JSON(fiber.Map{"user": newUser})
}

func GetUserStatus(c *fiber.Ctx) error {
	tokenString := c.Cookies("token")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return []byte(config.Config("JWT_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "No logged in", "data": err, "isLoggedIn": false})
		}

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Bad Reuqest", "data": err, "isLoggedIn": false})
	}

	if claims.Valid() != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "No logged in", "data": err, "isLoggedIn": false})
	}

	userId := fmt.Sprintf("%v", claims["id"])

	if !userExists(userId) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "No logged in", "data": err, "isLoggedIn": false})
	}

	return c.JSON(fiber.Map{"isLoggedIn": true})
}