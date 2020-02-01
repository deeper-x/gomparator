package routing

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// Fields are the graphql url targets
var Fields = graphql.Fields{
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
