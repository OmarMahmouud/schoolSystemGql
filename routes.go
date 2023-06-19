package main

import (
	"github.com/gin-gonic/gin"
)

func Routing(r *gin.Engine) {
	r.GET("/graphql", graphqlController)
	r.POST("/graphql", graphqlController)

}

func graphqlController(context *gin.Context) {
	graphqlHandler.ServeHTTP(context.Writer, context.Request)
}
