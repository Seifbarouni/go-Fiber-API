package main

import (
	"log"
	"os"
	"projects/Go-Fiber/api/controllers"
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
	// Config
	app := fiber.New(fiber.Config{
		Concurrency: 256*1024,
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
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)

	// Protected routes
	articles := app.Group("/api/v1")
	articles.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(map[string]string{"error": "unauthorized"})
		},
	}))
	articles.Get("/", controllers.Redirect)
	articles.Get("/articles", controllers.Articles)
	articles.Get("/articles/:id", controllers.GetArticleById)
	articles.Post("/articles", controllers.Add)
	articles.Put("/articles", controllers.Update)
	articles.Delete("/articles/:id", controllers.Delete)

	log.Fatal(app.Listen(":3000"))
}
