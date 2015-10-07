package main

import (
	"net/http"
	"github.com/chris-ramon/graphql-go/types"
	"github.com/sogko/graphql-go-handler"
)

// QueryType is the root query that has one field, `latestPost`, that returns a "Hello World!" string
//
// In GraphQL's type system shorthand notation (http://graphql.org/docs/typesystem),
// type Query {
//	 latestPost: String
// }

var queryType = types.NewGraphQLObjectType(types.GraphQLObjectTypeConfig{
	Name: "Query",
	Fields: types.GraphQLFieldConfigMap{
		"latestPost": &types.GraphQLFieldConfig{
			Type: types.GraphQLString,
			Resolve: func(p types.GQLFRParams) interface{} {
				return "Hello World!"
			},
		},
	},
})

var Schema, _ = types.NewGraphQLSchema(types.GraphQLSchemaConfig{
	Query: queryType,
})

func main() {

	// create a `graphl-go` HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)

	// Run the following curl command on terminal
	// $ curl -XPOST http://localhost:8080/graphql -H 'Content-Type: application/graphql' -d 'query Root{ latestPost }'
	// Expected output:
	// 	{
	//		"data": {
	//			"latestPost": "Hello World!"
	//		}
	//	}
}