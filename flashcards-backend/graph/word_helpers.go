// Package graph includes gqlgen
package graph

import (
	"context"
	"flashcards-backend/ent"
	"flashcards-backend/ent/word"
	"fmt"
)

// testLanguageSwap returns whether lang1 and 2 should be swapped or not
func testLanguageSwap(ctx context.Context, ent *ent.Client, lang1, lang2 string) (bool, error) {
	exists, err := ent.Word.
		Query().
		Where(
			word.And(
				word.Lang1EqualFold(lang1),
				word.Lang2EqualFold(lang2))).
		Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("checking lang code: %v", err)
	}
	if exists {
		return false, nil
	}
	exists, err = ent.Word.
		Query().
		Where(
			word.And(
				word.Lang1EqualFold(lang2),
				word.Lang2EqualFold(lang1))).
		Exist(ctx)
	if err != nil {
		return false, fmt.Errorf("checking lang code: %v", err)
	}
	return exists, nil
}
