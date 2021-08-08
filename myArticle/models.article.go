package main

import (
	"errors"
	"fmt"
)

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

func getAllArticles() []article {
	fmt.Println("Get All Articles")
	return articleList
}

func getArticleByID(art int) (*article, error) {
	for _, v := range articleList {
		if v.ID == art {
			return &v, nil
		}
	}
	return nil, errors.New("Article not found")
}
