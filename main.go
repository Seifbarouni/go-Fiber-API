package main

import (
	"log"
	"os"
	c "projects/Go-Fiber/api/controllers"
	db "projects/Go-Fiber/api/data"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	jwtware "github.com/gofiber/jwt/v3"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize database
	db.InitializeDB()
	// Config
	app := fiber.New(fiber.Config{
		Concurrency: 256 * 1024,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        6,
		LimitReached: func(c *fiber.Ctx) error {
			return c.JSON(map[string]string{"error": "too many requests, limit reached"})
		},
	}))

	// Auth routes
	app.Post("/register", c.Register)
	app.Post("/login", c.Login)

	// Protected routes
	articles := app.Group("/api/v1")
	articles.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "unauthorized"})
		},
	}))
	articles.Get("/", c.Redirect)
	articles.Get("/articles", c.Articles)
	articles.Get("/articles/:id", c.GetArticleById)
	articles.Post("/articles", c.Add)
	articles.Put("/articles", c.Update)
	articles.Delete("/articles/:id", c.Delete)

	log.Fatal(app.Listen(":3000"))
}
