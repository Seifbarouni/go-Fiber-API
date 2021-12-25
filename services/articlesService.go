package services

import (
	"errors"

	d "projects/Go-Fiber/api/data"
	m "projects/Go-Fiber/api/models"
)

type Articles []m.Article

func GetArticles() Articles {
	var articles Articles
	if result := d.DB.Find(&articles); result.Error != nil || result.RowsAffected == 0 {
		return []m.Article{}
	}
	return articles
}

func GetArticlesByAuthorId(id int) Articles {
	var articles Articles
	if result := d.DB.Where("author_id = ?", id).Find(&articles); result.Error != nil || result.RowsAffected == 0 {
		return []m.Article{}
	}
	return articles
}

func GetArticleById(id int) (*m.Article, error) {
	var article m.Article
	result := d.DB.First(&article, id)
	if result.Error != nil {
		return nil, errors.New("article not found")
	}
	return &article, nil
}

func AddArticle(art *m.Article) error {
	if art.Content == "" || art.Title == "" || art.AuthorId == 0 {
		return errors.New("invalid body")
	}
	result := d.DB.Create(art)
	return result.Error
}

func UpdateArticle(articleToUpdate m.Article) (*m.Article, error) {

	var article m.Article
	if result := d.DB.First(&article, articleToUpdate.ID); result.Error != nil {
		return nil, errors.New("article not found")
	}

	if articleToUpdate.Title != "" {
		article.Title = articleToUpdate.Title
	}
	if articleToUpdate.Content != "" {
		article.Content = articleToUpdate.Content
	}
	if articleToUpdate.AuthorId != 0 {
		article.AuthorId = articleToUpdate.AuthorId
	}
	if result := d.DB.Save(&article); result.Error != nil {
		return nil, errors.New("cannot update article")
	}
	return &article, nil
}

func DeleteArticle(id int) error {
	if result := d.DB.Delete(&m.Article{}, id); result.Error != nil || result.RowsAffected == 0 {
		return errors.New("cannot delete article")
	}
	return nil
}
