package main

import (
	"github.com/gin-gonic/gin"
)

func initializeRoutes() {
	// Handle the index route
	n_gin.GET("/", showIndexPage)
}

func indexHandler(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"title": "Home Page",
	})
}
