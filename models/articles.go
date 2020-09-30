package models

type Article struct {
	ID      int    `json:id`
	Title   string `json:title`
	Content string `json:content`
}

type Articles []*Article

var ArticleList = Articles{
	&Article{ID: 1, Title: "test1", Content: "content test1"},
	&Article{ID: 2, Title: "test2", Content: "content test2"},
	&Article{ID: 3, Title: "test3", Content: "content test3"},
}

func GetArticles() Articles {
	return ArticleList
}

func AddArticle(title string, content string) {
	art := Article{Title: title, Content: content}
	ArticleList = append(ArticleList, &art)
	update()
}

func UpdateArticle(id int, title string, content string) {
	for _, article := range ArticleList {
		if article.ID == id {
			if title == "" {
				continue
			} else {
				article.Title = title
			}

			if content == "" {
				continue
			} else {
				article.Content = content
			}
			break
		}
	}

}

func DeleteArticle(id int) {
	for i, article := range ArticleList {
		if article.ID == id {
			ArticleList = append(ArticleList[:i], ArticleList[i+1:]...)
		}
	}
	update()
}

func update() {
	for i, article := range ArticleList {
		article.ID = i + 1
	}
}