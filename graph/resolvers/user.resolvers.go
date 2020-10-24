package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/rushi444/gin-graphql/database"
	"github.com/rushi444/gin-graphql/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := database.GetDatabase()

	user := model.User{}

	user.ID = uuid.New().String()
	user.Name = input.Name

	db.Create(&user)

	return &user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	db := database.GetDatabase()

	db.Find(&r.users)
	return r.users, nil
}
