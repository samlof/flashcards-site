package graph

import (
	"context"
	"flashcards-backend/ent"
	"flashcards-backend/ent/enttest"
	"flashcards-backend/ent/migrate"
	"flashcards-backend/graph/model"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

func makeResolver(t *testing.T) *Resolver {
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

	return &Resolver{
		DB: client,
	}
}

func TestAddNewWord(t *testing.T) {
	resolver := makeResolver(t)

	ctx := context.Background()
	newWord := model.NewWord{
		Lang1: "fi",
		Lang2: "en",
		Word1: "jäätelö",
		Word2: "icecream",
	}
	word, err := resolver.Mutation().CreateWord(ctx, newWord)
	if err != nil {
		t.Errorf("adding word: %v", err)
	}

	// Try to get the same word
	returnedWord, err := resolver.Query().GetWords(ctx)
	if err != nil {
		t.Errorf("getting words: %v", err)
	}

	if returnedWord[0].ID != word.ID {
		t.Error("Didn't get same word")
	}
}

func TestDeleteWord(t *testing.T) {
	resolver := makeResolver(t)

	ctx := context.Background()
	newWord := model.NewWord{
		Lang1: "fi",
		Lang2: "en",
		Word1: "jäätelö",
		Word2: "icecream",
	}
	word, err := resolver.Mutation().CreateWord(ctx, newWord)
	if err != nil {
		t.Errorf("adding word: %v", err)
	}

	// Try to get the same word
	returnedWord, err := resolver.Query().GetWords(ctx)
	if err != nil {
		t.Errorf("getting words: %v", err)
	}

	if returnedWord[0].ID != word.ID {
		t.Error("Didn't get same word")
	}

	_, err = resolver.Mutation().DeleteWord(ctx, word.ID)
	if err != nil {
		t.Errorf("deleting words: %v", err)
	}

	// Get words again
	returnedWord, err = resolver.Query().GetWords(ctx)
	if err != nil {
		t.Errorf("getting words: %v", err)
	}
	if len(returnedWord) > 0 {
		t.Errorf("Word exists in db even afte deleting it")
	}
}

func TestAddDuplicateWord(t *testing.T) {
	resolver := makeResolver(t)

	ctx := context.Background()
	newWord := model.NewWord{
		Lang1: "fi",
		Lang2: "en",
		Word1: "jäätelö",
		Word2: "icecream",
	}
	word1, err := resolver.Mutation().CreateWord(ctx, newWord)
	if err != nil {
		t.Errorf("adding word: %v", err)
	}
	word2, err := resolver.Mutation().CreateWord(ctx, newWord)
	if err != nil {
		t.Errorf("adding duplicate word: %v", err)
	}
	if word1.ID != word2.ID {
		t.Error("Duplicate word should return the same word")
	}

	// Swap language and words
	newWord.Lang1, newWord.Lang2 = newWord.Lang2, newWord.Lang1
	newWord.Word1, newWord.Word2 = newWord.Word2, newWord.Word1

	word3, err := resolver.Mutation().CreateWord(ctx, newWord)
	if err != nil {
		t.Errorf("adding reversed word: %v", err)
	}
	if word1.ID != word3.ID {
		t.Error("Reversed word should return the same word")
	}
}

func TestUpdateWord(t *testing.T) {
	resolver := makeResolver(t)

	ctx := context.Background()
	newWord := model.NewWord{
		Lang1: "fi",
		Lang2: "en",
		Word1: "jäätelö",
		Word2: "icecream",
	}
	word, err := resolver.Mutation().CreateWord(ctx, newWord)
	if err != nil {
		t.Errorf("Failed adding word: %v", err)
	}
	updateWord := model.UpdateWord{
		ID:    word.ID,
		Lang1: "fi",
		Lang2: "en",
		Word1: "jäätelöä",
		Word2: "icecream",
	}

	// Try to get the same word
	_, err = resolver.Mutation().UpdateWord(ctx, updateWord)
	if err != nil {
		t.Errorf("Failed updating word: %v", err)
	}

	newWord.Word1, newWord.Word2 = "Tuoli", "Chair"
	word2, err := resolver.Mutation().CreateWord(ctx, newWord)
	if err != nil {
		t.Errorf("adding word: %v", err)
	}
	updateWord.ID = word2.ID
	_, err = resolver.Mutation().UpdateWord(ctx, updateWord)
	if err == nil {
		t.Error("Should not update to a duplicate word")
	}

	updateWord.Lang1, updateWord.Lang2 = updateWord.Lang2, updateWord.Lang1
	updateWord.Word1, updateWord.Word2 = updateWord.Word2, updateWord.Word1
	_, err = resolver.Mutation().UpdateWord(ctx, updateWord)
	if err == nil {
		t.Error("Should not update to a reversed word")
	}
}

/*  Helper for running query
func runQuery(t *testing.T, query string) json.RawMessage {
    source := &ast.Source{Input: query}
    doc, gerr := parser.ParseQuery(source)
    if gerr != nil {
        t.Fatal("query failed to parse:", gerr)
    }

    errs := validator.Validate(exec.Schema(), doc)
    if len(errs) > 0 {
        t.Fatal("error validating query:", errs)
    }

    reqCtx := graphql.NewRequestContext(doc, query, map[string]interface{}{})
    ctx := context.Background()
    ctx = graphql.WithRequestContext(ctx, reqCtx)

    if len(doc.Operations) == 0 {
        t.Fatal("no graphql operations defined")
    }
    // assuming an anonymous operation
    operation := doc.Operations.ForName("")

    var resp *graphql.Response{}
    switch operation.Operation {
    case ast.Query:
        resp = exec.Query(ctx, operation)
    case ast.Mutation:
        resp = exec.Mutation(ctx, operation)
    }

    if len(resp.Errors) > 0 {
        t.Fatal("errors executing graphql operation:", resp.Errors)
    }
    return resp.Data
}
*/
