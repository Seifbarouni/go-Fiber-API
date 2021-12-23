package controllers

import (
	"net/http"
	m "projects/Go-Fiber/api/models"
	s "projects/Go-Fiber/api/services"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	newUser := new(m.User)
	if err := c.BodyParser(newUser); err != nil {
		return c.Status(http.StatusTeapot).JSON(map[string]string{"error": "invalid body"})
	}
	if newUser.Email == "" || newUser.Password == "" || newUser.Name == "" {
		return c.Status(http.StatusTeapot).JSON(map[string]string{"error": "invalid body"})
	}
	if token, err := s.Register(newUser); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	} else {
		return c.Status(http.StatusCreated).JSON(map[string]string{"success": "user added!", "token": token})
	}
}

func Login(c *fiber.Ctx) error {
	user := new(m.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusTeapot).JSON(map[string]string{"error": "invalid body"})
	}
	if user.Email == "" || user.Password == "" {
		return c.Status(http.StatusTeapot).JSON(map[string]string{"error": "invalid body"})
	}
	if token, err := s.Login(user); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
	} else {
		return c.JSON(map[string]string{"success": "successful login!", "token": token})
	}

}
