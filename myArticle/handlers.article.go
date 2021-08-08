package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Below is html oly index page rendering
/*func showIndexPage(c *gin.Context) {
	articles := getAllArticles()

	// Call the HTML method of the Context to render a template
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"index.html",
		// Pass the data that the page uses
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)

}*/

//modify the showIndexPage to render in any of json. xml or html based on the request's accept type
func showIndexPage(c *gin.Context) {
	articles := getAllArticles()
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles,
	}, "index.html")
}

func getArticle(ctx *gin.Context) {
	articleID, err := strconv.Atoi(ctx.Param("article_id"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	article, err := getArticleByID(articleID)

	if err != nil {
		fmt.Println(err)
	}
	ctx.HTML(200,
		"article.html",
		gin.H{
			"payload": article,
		})
}
