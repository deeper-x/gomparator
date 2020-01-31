package main

import (
	"encoding/json"
	"fmt"
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
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// curl -i http://127.0.0.1:8000/?query={hello(id:100)}
				return fmt.Sprintf("Hello! passed id is: %d", p.Args["id"]), nil
			},
		},
		"foo": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "And content of foo is...", nil
			},
		},
		"bar": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Oh yeah, content of bar is...", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}

	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatal(err)
	}

	router.GET("/", func(c *gin.Context) {
		param := c.Request.URL.Query()

		query := param["query"][0]

		params := graphql.Params{Schema: schema, RequestString: query}

		r := graphql.Do(params)

		data, _ := json.Marshal(r)

		var ro map[string]map[string]interface{}
		err = json.Unmarshal(data, &ro)

		if err != nil {
			log.Println(err)
		}

		res := ro["data"]

		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": res,
		})
	})

	router.Run(":8000")
}
