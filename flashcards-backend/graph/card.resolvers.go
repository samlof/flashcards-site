package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flashcards-backend/ent"
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/word"
	"flashcards-backend/graph/generated"
	"flashcards-backend/graph/model"
	"flashcards-backend/modelconv"
	"fmt"
	"strconv"
	"time"
)

func (r *mutationResolver) CardStatus(ctx context.Context, input model.CardStatus) (*model.CardLog, error) {
	idInt, err := strconv.Atoi(input.CardID)
	if err != nil {
		return nil, fmt.Errorf("parsing id %s to int: %v", input.CardID, err)
	}
	// Get old log item to calculate when this should be scheduled for
	oldLog, err := r.DB.CardLog.Query().
		Where(cardlog.ID(idInt)).
		Order(ent.Desc(cardlog.FieldCreateTime)).
		First(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting old log item: %v", err)
	}
	// Calculate when this will be scheduled next

	mod := logModifiers[input.Result]
	// Default hoursSince last seen to 24
	hoursSince := 24 * time.Hour
	if oldLog != nil {
		hoursSince = time.Since(oldLog.ScheduledFor)
	}
	scheduleFor := time.Now().Add(time.Duration(float64(time.Hour*hoursSince) * mod))

	// Insert to db
	log, err := r.DB.CardLog.Create().
		SetResult(modelconv.ModelCardResult(input.Result)).
		SetScheduledFor(scheduleFor).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating log item: %v", err)
	}
	return modelconv.CardLog(log), nil
}

func (r *queryResolver) ScheduledWords(ctx context.Context, newWordCount *int) (*model.ScheduledWordsResponse, error) {
	// Get cards scheduled for review
	cards, err := r.DB.CardLog.Query().
		WithCard().
		Where(cardlog.ScheduledForLTE(time.Now())).
		Order(ent.Desc(cardlog.FieldScheduledFor)).
		Limit(500).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting logs: %v", err)
	}
	reviews := modelconv.CardLogS(cards)
	ret := &model.ScheduledWordsResponse{
		Reviews: reviews,
	}

	// If no new words wanted, return here
	if newWordCount == nil {
		return ret, nil
	}

	// Get new cards
	newWords, err := r.DB.Word.Query().
		Where(word.Not(word.HasCardLogs())).
		Limit(*newWordCount).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting new words: %v", err)
	}

	ret.NewWords = modelconv.WordS(newWords)

	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
