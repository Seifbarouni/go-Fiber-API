package controllers

import (
	"net/http"
	"strconv"

	m "projects/Go-Fiber/api/models"
	s "projects/Go-Fiber/api/services"

	"github.com/gofiber/fiber/v2"
)


func Redirect(c *fiber.Ctx)error{
	return c.Redirect("/articles",http.StatusFound)
}

func Articles(c* fiber.Ctx)error{
	articles:=s.GetArticles()
	return c.JSON(map[string]interface{}{"articles":articles})
}  

func GetArticleById(c *fiber.Ctx)error{
	idStr:=c.Params("id")
	id,err:=strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(map[string]string{"error":"Invalid ID"})
	}
	article,notFound:=s.GetArticleById(id)
	if notFound != nil {
		return c.JSON(map[string]string{"error":"Article not found"})
	}
	return c.JSON(*article)
}


func Add(c *fiber.Ctx)error{
	newArticle := new(m.Article)

	if err := c.BodyParser(newArticle); err != nil {
		return c.JSON(map[string]string{"error":"Invalid body"})
	}
	if err:=s.AddArticle(newArticle) ; err != nil {
		return c.JSON(map[string]string{"error":err.Error()})
	}
	return c.JSON(map[string]string{"success":"Article added"})
}

func Update(c *fiber.Ctx)error{
	articleToUpdate := new(m.Article)
	if err := c.BodyParser(articleToUpdate); err != nil {
		return c.JSON(map[string]string{"error":"Invalid body"})
	}
	updatedArticle,err:=s.UpdateArticle(*articleToUpdate)
	if err != nil {
		return c.JSON(map[string]string{"error":"Article not found"})
	}
	return c.JSON(*updatedArticle)
}

func Delete(c *fiber.Ctx)error{
	idStr:=c.Params("id")
	id,err:=strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(map[string]string{"error":"Invalid ID"})
	}
	err=s.DeleteArticle(id)
	if err != nil {
		return c.JSON(map[string]string{"error":"Article not found"})
	}
	return c.JSON(map[string]string{"success":"Article deleted"})
}