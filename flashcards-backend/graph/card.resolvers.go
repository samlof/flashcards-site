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
	cardId, err := strconv.Atoi(input.CardID)
	if err != nil {
		return nil, fmt.Errorf("parsing cardid %s to int: %v", input.CardID, err)
	}
	card, err := r.DB.Word.Query().
		Where(word.ID(cardId)).
		WithCardLogs(func(q *ent.CardLogQuery) {
			q.Where(cardlog.Reviewed(false))
		}).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting card with id %d: %v", cardId, err)
	}

	var oldLog *ent.CardLog
	// Get old log item related to this
	if len(card.Edges.CardLogs) != 0 {
		// Get old log item to calculate when this should be scheduled for
		oldLog, err = r.DB.CardLog.UpdateOne(card.Edges.CardLogs[0]).SetReviewed(true).Save(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, fmt.Errorf("getting old log item: %v", err)
		}
	}

	// Calculate when this will be scheduled next
	mod := model.LogModifiers[input.Result]
	// Default hoursSince last seen to 24
	var hoursSince float64 = 24
	if oldLog != nil {
		hoursSince = time.Since(oldLog.ScheduledFor).Hours()
	}
	scheduleFor := time.Now().Add(time.Duration(float64(time.Hour) * hoursSince * mod))

	// Insert to db
	log, err := r.DB.CardLog.Create().
		SetResult(modelconv.ModelCardResult(input.Result)).
		SetScheduledFor(scheduleFor).
		SetCard(card).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("creating log item: %v", err)
	}
	log.Edges.Card = card
	return modelconv.CardLog(log), nil
}

func (r *queryResolver) ScheduledWords(ctx context.Context, newWordCount *int) (*model.ScheduledWordsResponse, error) {
	// Get cards scheduled for review
	cards, err := r.DB.CardLog.Query().
		WithCard().
		Where(cardlog.And(cardlog.Reviewed(false), cardlog.ScheduledForLTE(time.Now()))).
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

	// Get how many new cards have been done during last 24h
	alreadyDoneIds, err := r.DB.CardLog.Query().
		Where(cardlog.CreateTimeGT(time.Now().Add(time.Hour * 24 * -1))).
		IDs(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting already done ids: %v", err)
	}
	if len(alreadyDoneIds) > 0 {
		allIds, err := r.DB.CardLog.Query().Where(cardlog.IDIn(alreadyDoneIds...)).Select(cardlog.ForeignKeys[0]).Ints(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting already done allIds: %v", err)
		}
		idCounts := make(map[int]int, len(alreadyDoneIds))
		for _, id := range allIds {
			idCounts[id]++
		}
		*newWordCount -= len(idCounts)
	}
	if *newWordCount > 0 {
		// Get new cards
		newWords, err := r.DB.Word.Query().
			Where(word.Not(word.HasCardLogs())).
			Limit(*newWordCount).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting new words: %v", err)
		}

		ret.NewWords = modelconv.WordS(newWords)
	}

	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
