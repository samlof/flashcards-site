package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flashcards-backend/graph/generated"
	"flashcards-backend/graph/model"
	"flashcards-backend/modelconv"
	"fmt"
)

func (r *queryResolver) ScheduledWords(ctx context.Context) ([]*model.CardLog, error) {
	cards, err := r.DB.CardLog.Query().WithCard().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting logs: %v", err)
	}
	ret := make([]*model.CardLog, 0, len(cards))
	for _, card := range cards {
		ret = append(ret, modelconv.CardLog(card))
	}
	return ret, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
