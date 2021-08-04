package main

import "github.com/gin-gonic/gin"

//create a gin object
var n_gin *gin.Engine

func main() {
	n_gin = gin.Default()
	n_gin.LoadHTMLGlob("templates/*")
	initializeRoutes()
	n_gin.Run(":8000")
}
