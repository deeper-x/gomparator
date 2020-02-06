package main

import (
	"encoding/json"
	"log"

	"net/http"

	"github.com/deeper-x/gomparator/records"
	"github.com/deeper-x/gomparator/routing"
	"github.com/gin-gonic/gin"
)

var result map[string]map[string]interface{}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/js", "./js")

	router.GET("/", func(c *gin.Context) {
		query := c.Request.URL.Query()

		res := routing.GetResult(query, routing.Schema)

		data, _ := json.Marshal(res)

		err := json.Unmarshal(data, &result)

		if err != nil {
			log.Println(err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": result["data"],
		})
	})

	conn := records.NewRepository()

	conn.SetKey("foo", "bar")
	log.Println(conn.GetKey("foo"))

	router.Run(":8000")
}
