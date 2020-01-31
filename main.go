package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"

	"net/http"
)

type query struct{}

func (q *query) Hello() string {
	return "hello!"
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/js", "./js")

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hi, there...", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatal(err)
	}

	query := `
		{
			hello
		}
	`

	params := graphql.Params{Schema: schema, RequestString: query}

	r := graphql.Do(params)

	data, _ := json.Marshal(r)

	var ro map[string]map[string]interface{}
	err = json.Unmarshal(data, &ro)

	if err != nil {
		log.Println(err)
	}

	res := ro["data"]

	log.Println("result:", res["hello"])

	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "here I am",
		})
	})

	router.Run(":8000")
}
