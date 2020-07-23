// Package graph includes gqlgen
package graph

import "flashcards-backend/ent"

//go:generate go run github.com/99designs/gqlgen

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB *ent.Client
}
