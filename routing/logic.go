package routing

import (
	"net/url"

	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: Fields}
var schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

// Schema is the params schema
var Schema, _ = graphql.NewSchema(schemaConfig)

// GetResult is returning graphql result
func GetResult(v url.Values, s graphql.Schema) *graphql.Result {
	query := v["query"][0]

	params := graphql.Params{Schema: s, RequestString: query}

	r := graphql.Do(params)

	return r
}
