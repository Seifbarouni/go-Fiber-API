package controllers

import (
	"net/http"
	"projects/Go-Fiber/api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)


func Redirect(c *fiber.Ctx)error{
	return c.Redirect("/articles",http.StatusFound)
}

func Articles(c* fiber.Ctx)error{
	articles:=models.GetArticles()
	return c.JSON(articles)
}  


func Add(c *fiber.Ctx)error{
	title:=c.Params("title") 
	content:=c.Params("content")
	models.AddArticle(title,content) 
	return c.Redirect("/",http.StatusOK)
}

func Update(c *fiber.Ctx)error{
	title:=c.Params("title")
	content:=c.Params("content")
	idstr:=c.Params("id")

	id,err:=strconv.Atoi(idstr)
	if err != nil {
		return c.SendString("Invalid ID")
	}
	models.UpdateArticle(id,title,content)
	return c.Redirect("/",http.StatusOK)
}

func Delete(c *fiber.Ctx)error{
	idstr:=c.Params("id")
	id,err:=strconv.Atoi(idstr)
	if err != nil {
		return c.SendString("Invalid ID")
	}
	models.DeleteArticle(id)
	return c.Redirect("/",http.StatusOK)
	
}