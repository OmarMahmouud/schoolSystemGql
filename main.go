package main

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var schema graphql.Schema

// GraphQL endpoint
var graphqlHandler = handler.New(&handler.Config{
	Schema:   &schema,
	Pretty:   true,
	GraphiQL: true,
})

func main() {
	ConnectToDataBase()
	Migrate()
	r := gin.Default()

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(mutationType),
	}

	schema, _ = graphql.NewSchema(schemaConfig)

	Routing(r)

	r.Run(":8080")
}
