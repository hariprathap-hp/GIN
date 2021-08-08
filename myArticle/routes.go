package main

func initializeRoutes() {
	// Handle the index route
	n_gin.GET("/", showIndexPage)
	n_gin.GET("/article/view/:article_id", getArticle)
}
