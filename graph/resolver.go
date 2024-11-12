package graph

import "github.com/TechBuilder-360/business-directory-backend/graph/model"

//go:generate go run github.com/arsmn/fastgql generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
	users []*model.User
}
