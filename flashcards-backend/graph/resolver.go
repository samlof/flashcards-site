// Package graph includes gqlgen
package graph

//go:generate go run github.com/99designs/gqlgen

import (
	"flashcards-backend/ent"
)

// Resolver serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB *ent.Client
}
