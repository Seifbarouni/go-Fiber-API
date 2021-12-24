package services

import (
	"errors"

	m "projects/Go-Fiber/api/models"
)

type Articles []*m.Article

var ArticleList = Articles{
	&m.Article{ID: 1, Title: "test1", AuthorId: 1,Content: "content test1"},
	&m.Article{ID: 2, Title: "test2", AuthorId: 2,Content: "content test2"},
	&m.Article{ID: 3, Title: "test3", AuthorId: 2,Content: "content test3"},
}

func GetArticles() Articles {
	return ArticleList
}

func GetArticleById(id int) (*m.Article, error) {
	for _, article := range ArticleList {
		if article.ID == id {
			return article, nil
		}
	}
	return nil, errors.New("article not found")
}

func AddArticle(art *m.Article) error {
	art.ID = len(ArticleList) + 1
	if art.Content == "" || art.Title == "" {
		return errors.New("invalid body")
	}
	ArticleList = append(ArticleList, art)
	return nil
}

func UpdateArticle(articleToUpdate m.Article) (*m.Article, error) {
	for _, article := range ArticleList {
		if article.ID == articleToUpdate.ID {
			if articleToUpdate.Title != "" {
				article.Title = articleToUpdate.Title
			}
			if articleToUpdate.Content != "" {
				article.Content = articleToUpdate.Content
			}
			return article, nil
		}
	}

	return nil, errors.New("article not found")
}

func DeleteArticle(id int) error {
	for i, article := range ArticleList {
		if article.ID == id {
			ArticleList = append(ArticleList[:i], ArticleList[i+1:]...)
			update()
			return nil
		}
	}
	return errors.New("article not found")
}

func update() {
	for i, article := range ArticleList {
		article.ID = i + 1
	}
}
