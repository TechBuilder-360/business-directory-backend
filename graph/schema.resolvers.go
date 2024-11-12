package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/TechBuilder-360/business-directory-backend/graph/generated"
	"github.com/TechBuilder-360/business-directory-backend/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", randNumber),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	user := model.User{ID: fmt.Sprintf("T%d", randNumber), Name: input.Name}
	r.users = append(r.users, &user)
	return &user, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.DeleteUser) (*string, error) {
	users := r.users
	index := input.Index

	if len(users) < input.Index {
		return nil, errors.New("index not found")
	}

	r.users = append(users[:index], users[index+1:]...)

	return nil, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) GetUser(ctx context.Context, input model.GetUser) (*model.User, error) {
	if len(r.users) < input.ID {
		return nil, errors.New("user does not exist")
	}
	return r.users[input.ID], nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
