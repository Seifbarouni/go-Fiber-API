package main

import (
	"log"
	"projects/Go-Fiber/api/controllers"

	"github.com/gofiber/fiber/v2"
)

func main(){
	app:=fiber.New()

	app.Get("/",controllers.Redirect)  
	app.Get("/articles",controllers.Articles)
	app.Post("/articles/add/:title/:content",controllers.Add)
	app.Put("/articles/update/:id/:title/:content",controllers.Update)
	app.Delete("/articles/delete/:id",controllers.Delete)

	log.Fatal(app.Listen(":3000"))
}