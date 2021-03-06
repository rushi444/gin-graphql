package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rushi444/gin-graphql/handlers"
)

const defaultPort = ":4000"

func main() {
	r := gin.Default()
	r.POST("/graphql", handlers.GraphQLHandler())
	r.GET("/", handlers.PlaygroundHandler())
	r.Run(defaultPort)
}
