package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/rushi444/gin-graphql/database"
	"github.com/rushi444/gin-graphql/graph/generated"
	"github.com/rushi444/gin-graphql/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	db := database.GetDatabase()

	todo := model.Todo{}
	user := model.User{}

	todo.ID = uuid.New().String()
	todo.Text = input.Text
	todo.Done = false
	todo.UserID = input.UserID

	db.First(&user, "id = ?", todo.UserID)
	todo.User = &user

	db.Create(&todo)

	return &todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	db := database.GetDatabase()

	db.Find(&r.todos)
	for _, todo := range r.todos {
		user := model.User{}
		db.Where(&model.User{ID: todo.UserID}).Find(&user)
		todo.User = &user
	}
	return r.todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
