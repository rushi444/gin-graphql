package graph

import (
	"github.com/rushi444/gin-graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver : Resolver Struct
type Resolver struct {
	todos []*model.Todo
	users []*model.User
}
