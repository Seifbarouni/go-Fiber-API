package main

import (
	"log"
	"projects/Go-Fiber/api/controllers"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main(){
	// Config
	app:=fiber.New()

	app.Use(csrf.New())

	app.Use(cors.New(cors.Config{
    AllowOrigins: "http://localhost:3000",
    AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	app.Use(limiter.New(limiter.Config{
		Duration:10* time.Second ,
		Max:6,
		LimitReached: func(c *fiber.Ctx)error{
			return c.JSON(map[string]string{"error":"too many requests, limit reached"})
		},
	}))

	// Routes
	app.Get("/",controllers.Redirect)  
	app.Get("/articles",controllers.Articles)
	app.Get("/articles/:id",controllers.GetArticleById)
	app.Post("/articles",controllers.Add)
	app.Put("/articles",controllers.Update)
	app.Delete("/articles/:id",controllers.Delete)

	log.Println("Listening on port 3000")
	log.Fatal(app.Listen(":3000"))
}