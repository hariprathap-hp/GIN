package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create a gin object
var n_gin *gin.Engine

func main() {
	n_gin = gin.Default()
	n_gin.LoadHTMLGlob("templates/*")
	initializeRoutes()
	n_gin.Run(":8000")
}

func render(c *gin.Context, data gin.H, template string) {
	accept := c.Request.Header.Get("accept")
	switch accept {
	case "application/json":
		fmt.Println("Parsing JSON")
		c.JSON(200, data["payload"])
	case "application/xml":
		c.XML(200, data["payload"])
	default:
		fmt.Println("Parsing HTML")
		c.HTML(http.StatusOK, template, data)
	}
}
