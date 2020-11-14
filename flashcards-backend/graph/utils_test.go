package graph_test

import (
	"flashcards-backend/ent"
	"flashcards-backend/ent/enttest"
	"flashcards-backend/ent/migrate"
	"flashcards-backend/graph"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func makeResolver(t *testing.T) *graph.Resolver {
	t.Helper()
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	}

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)

	t.Cleanup(func() {
		err := client.Close()
		if err != nil {
			t.Errorf("error closing ent client: %v", err)
		}
	})

	return &graph.Resolver{
		DB: client,
	}
}
