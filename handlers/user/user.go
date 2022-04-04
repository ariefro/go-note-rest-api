package userHandler

import (
	"fmt"
	"strconv"

	"github.com/ariefro/notes-server/database"
	"github.com/ariefro/notes-server/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func validToken(t *jwt.Token, id string) bool {
	n, err := strconv.Atoi(id)
	if err != nil {
		return false
	}

	claims := t.Claims.(jwt.MapClaims)
	uid := int(claims["user_id"].(float64))

	return uid == n
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	id := c.Params("userId")
	db.Find(&user, "id = ?", id)
	if user.Username == "" {
		return c.Status(404).JSON(fiber.Map{
			"status": "error",
			"message": "No user found with ID " + id,
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "User found",
		"data": user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email string `json:"email"`
	}

	db := database.DB
	user := new(model.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Review your input",
			"data": err,
		}) 
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Could not hash password",
			"data": err,
		})
	}

	user.Password = hash
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Could not create user",
			"data": err,
		})
	}

	newUser := NewUser {
		Email: user.Email,
		Username: user.Username,
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "Created user",
		"data": newUser,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUserInput struct {
		Names string `json:"names"`
	}

	var uui UpdateUserInput
	err := c.BodyParser(&uui)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status": "error",
			"message": "Review your input",
			"data": err,
		}) 
	}

	id := c.Params("id")

	fmt.Println(id)
	fmt.Println(c.Locals("User"))
	token := c.Locals("user").(*jwt.Token)
	fmt.Println(token)

	if !validToken(token, id) {
		return c.Status(500).JSON(fiber.Map{
			"status": "error", 
			"message": "Invalid token id", 
			"data": nil})
	}

	db := database.DB
	var user model.User

	db.First(&user, id)
	user.Names = uui.Names
	db.Save(&user)

	return c.JSON(fiber.Map{
		"status": "success",
		"message": "User successfully updated",
		"data": user,
	})
}