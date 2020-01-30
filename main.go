package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/js", "./js")

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "here I am",
		})
	})

	router.Run(":8000")
}
