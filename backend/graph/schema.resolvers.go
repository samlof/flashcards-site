package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flashcards-backend/ent"
	"flashcards-backend/ent/word"
	"flashcards-backend/graph/generated"
	"flashcards-backend/graph/model"
	"fmt"
	"strconv"
)

func (r *mutationResolver) CreateWord(ctx context.Context, input model.NewWord) (*model.Word, error) {
	word, err := r.DB.Word.Create().
		SetLangData(input.LangData).
		SetWord1(input.Word1).
		SetWord2(input.Word2).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("error saving word: %v", err)
	}

	return &model.Word{
		ID:       strconv.Itoa(word.ID),
		LangData: word.LangData,
		Word1:    word.Word1,
		Word2:    word.Word2,
	}, nil
}

func (r *queryResolver) GetWords(ctx context.Context) ([]*model.Word, error) {
	words, err := r.DB.Word.Query().Order(ent.Desc(word.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting words: %v", err)
	}
	ret := make([]*model.Word, 0, len(words))
	for _, word := range words {
		ret = append(ret, &model.Word{
			ID:        strconv.Itoa(word.ID),
			LangData:  word.LangData,
			Word1:     word.Word1,
			Word2:     word.Word2,
			CreatedAt: word.CreatedAt,
		})
	}
	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
