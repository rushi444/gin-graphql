package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/graphql/playground"
)


// PlaygroundHandler : GraphQL Playground
func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}