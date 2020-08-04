package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"flashcards-backend/ent"
	"flashcards-backend/ent/word"
	"flashcards-backend/graph/generated"
	"flashcards-backend/graph/model"
	"flashcards-backend/modelconv"
	"fmt"
	"strconv"
)

func (r *mutationResolver) CreateWord(ctx context.Context, input model.NewWord) (*model.Word, error) {
	shouldSwap, err := testLanguageSwap(ctx, r.DB, input.Lang1, input.Lang2)
	if err != nil {
		return nil, fmt.Errorf("checking swap: %v", err)
	}
	if shouldSwap {
		input.Lang1, input.Lang2 = input.Lang2, input.Lang1
		input.Word1, input.Word2 = input.Word2, input.Word1
	}

	// Check duplicate
	exists, err := r.DB.Word.Query().
		Where(word.And(
			word.Lang1EqualFold(input.Lang1), word.Lang2EqualFold(input.Lang2),
			word.Word1EqualFold(input.Word1), word.Word2EqualFold(input.Word2))).
		Limit(1).
		All(ctx)

	if err != nil {
		return nil, fmt.Errorf("checking duplicate: %v", err)
	}

	// If found a duplicate, return the existing one
	if len(exists) > 0 {
		return modelconv.Word(exists[0]), nil
	}

	word, err := r.DB.Word.Create().
		SetLang1(input.Lang1).
		SetLang2(input.Lang2).
		SetWord1(input.Word1).
		SetWord2(input.Word2).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("saving word: %v", err)
	}

	return modelconv.Word(word), nil
}

func (r *mutationResolver) DeleteWord(ctx context.Context, id string) (string, error) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return "", fmt.Errorf("parsing id %s to int: %v", id, err)
	}
	err = r.DB.Word.DeleteOneID(idInt).Exec(ctx)
	if err != nil {
		return "", fmt.Errorf("deleting word: %v", err)
	}
	return id, nil
}

func (r *mutationResolver) UpdateWord(ctx context.Context, input model.UpdateWord) (*model.Word, error) {
	idInt, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, fmt.Errorf("parsing id %s to int: %v", input.ID, err)
	}

	shouldSwap, err := testLanguageSwap(ctx, r.DB, input.Lang1, input.Lang2)
	if err != nil {
		return nil, fmt.Errorf("checking swap: %v", err)
	}
	if shouldSwap {
		input.Lang1, input.Lang2 = input.Lang2, input.Lang1
		input.Word1, input.Word2 = input.Word2, input.Word1
	}
	// Check duplicate
	exists, err := r.DB.Word.
		Query().
		Where(
			word.And(
				word.Lang1EqualFold(input.Lang1), word.Lang2EqualFold(input.Lang2),
				word.Word1EqualFold(input.Word1), word.Word2EqualFold(input.Word2))).
		Exist(ctx)
	if err != nil {
		return nil, fmt.Errorf("checking duplicate: %v", err)
	}
	if exists {
		return nil, errors.New("word already exists")
	}

	word, err := r.DB.Word.
		UpdateOneID(idInt).
		SetWord1(input.Word1).
		SetWord2(input.Word2).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("updating word: %v", err)
	}

	return modelconv.Word(word), nil
}

func (r *queryResolver) GetWords(ctx context.Context) ([]*model.Word, error) {
	words, err := r.DB.Word.Query().Order(ent.Desc(word.FieldCreateTime)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting words: %v", err)
	}
	ret := make([]*model.Word, 0, len(words))
	for _, word := range words {
		ret = append(ret, modelconv.Word(word))
	}
	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
