package main

import (
	"fmt"
	"testing"
)

func TestGetAllArticles(t *testing.T) {
	fmt.Println("Testing Articles")
	aList := getAllArticles()
	if len(aList) != len(articleList) {
		t.Fail()
	}

	for i, v := range aList {
		if v.ID != articleList[i].ID || v.Title != articleList[i].Title ||
			v.Content != articleList[i].Content {
			t.Fail()
			break
		}
	}
}
