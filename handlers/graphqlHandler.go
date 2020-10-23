package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/rushi444/gin-graphql/graph"
	"github.com/rushi444/gin-graphql/graph/generated"
)

// GraphQLHandler : GraphQL Server
func GraphQLHandler() gin.HandlerFunc {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}
