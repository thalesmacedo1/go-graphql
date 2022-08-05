package graph

import "github.com/thalesmacedo1/go-graphql/graph/model"

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Categories []*model.Category
	Courses    []*model.Course
	Chapters   []*model.Chapter
}
