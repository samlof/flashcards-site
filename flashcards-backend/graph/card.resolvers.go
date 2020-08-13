package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"flashcards-backend/ent"
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/ent/cardschedule"
	"flashcards-backend/ent/word"
	"flashcards-backend/graph/generated"
	"flashcards-backend/graph/model"
	"flashcards-backend/modelconv"
	"fmt"
	"math/rand"
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
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting card with id %d: %v", cardId, err)
	}

	// For retry, we shouldn't reschedule the card
	if input.Result != model.CardResultRetry {

		// Update the old scheduled item to be done
		_, err = r.DB.CardSchedule.Update().
			Where(cardschedule.And(cardschedule.HasCardWith(word.ID(cardId)), cardschedule.Reviewed(false))).
			SetReviewed(true).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("updating scheduled item: %v", err)
		}

		// Get old log item related to this
		oldLog, err := r.DB.CardLog.Query().
			Where(cardlog.HasCardWith(word.ID(cardId))).
			Order(ent.Desc(cardlog.FieldCreateTime)).
			First(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return nil, fmt.Errorf("finding old log: %v", err)
		}

		// Calculate when this will be scheduled next
		mod := model.LogModifiers[input.Result]
		// Default hoursSince last seen to 24
		var hoursSince float64 = 24
		if oldLog != nil {
			hoursSince = time.Since(oldLog.CreateTime).Hours()
		}
		hoursSince = hoursSince * mod
		// Should never be below 1 day
		if hoursSince < 24 {
			hoursSince = 24
		}
		scheduleFor := time.Now().Add(time.Duration(float64(time.Hour) * hoursSince))

		// Insert to db
		_, err = r.DB.CardSchedule.Create().
			SetScheduledFor(scheduleFor).
			SetCard(card).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("creating log item: %v", err)
		}
	}
	// Add the status to db
	dbResult := modelconv.ModelCardResult(input.Result)
	cardLog, err := r.DB.CardLog.Create().SetCard(card).SetResult(dbResult).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("adding log to db for id %d result %s: %v", cardId, dbResult, err)
	}
	cardLog.Edges.Card = card
	return modelconv.CardLog(cardLog), nil
}

func (r *queryResolver) ScheduledWords(ctx context.Context, shuffle bool) (*model.ScheduledWordsResponse, error) {
	// Get cards scheduled for review
	scheduledCards, err := r.DB.CardSchedule.Query().
		WithCard().
		Where(
			cardschedule.And(
				cardschedule.Reviewed(false),
				cardschedule.ScheduledForLTE(time.Now()),
			)).
		Order(ent.Desc(cardschedule.FieldScheduledFor)).
		Limit(500).
		All(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting logs: %v", err)
	}

	cards := make([]*model.Word, len(scheduledCards))
	for i := range scheduledCards {
		cards[i] = modelconv.Word(scheduledCards[i].Edges.Card)
	}
	ret := &model.ScheduledWordsResponse{
		Cards: cards,
	}

	// Get settings to find how many cards a day
	settings, err := r.DB.UserSettings.Query().First(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting settings: %v", err)
	}

	// If no new words wanted, return here
	newWordCount := settings.NewCardsPerDay
	if newWordCount == 0 {
		return ret, nil
	}

	// Get which cards have been done in last 24h
	alreadyDoneCardIds, err := r.DB.CardLog.Query().
		Where(cardlog.And(
			cardlog.CreateTimeGT(time.Now().Add(time.Hour*24*-1)),
			cardlog.ResultNEQ(cardlog.ResultRetry),
		)).
		Select(cardlog.ForeignKeys[0]).
		Ints(ctx)

	if err != nil {
		return nil, fmt.Errorf("getting already done ids: %v", err)
	}
	if len(alreadyDoneCardIds) > 0 {
		// Get how many times each card has been done in the past
		doneCardIds, err := r.DB.CardLog.Query().
			Where(cardlog.And(
				cardlog.HasCardWith(word.IDIn(alreadyDoneCardIds...)),
				cardlog.ResultNEQ(cardlog.ResultRetry),
			)).
			Select(cardlog.ForeignKeys[0]).Ints(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting already done doneCardIds: %v", err)
		}
		idCounts := make(map[int]int, len(alreadyDoneCardIds))
		for _, id := range doneCardIds {
			idCounts[id]++
		}
		for _, count := range idCounts {
			// If card has been done only once, it was a new card
			if count == 1 {
				newWordCount--
			}
		}
	}

	if newWordCount > 0 {
		// Get new cards
		newWords, err := r.DB.Word.Query().
			Where(word.Not(word.HasCardSchedules())).
			Limit(newWordCount).
			All(ctx)
		if err != nil {
			return nil, fmt.Errorf("getting new words: %v", err)
		}

		ret.Cards = append(ret.Cards, modelconv.WordS(newWords)...)
	}

	if shuffle {
		cards = ret.Cards
		rand.Shuffle(len(cards), func(i, j int) { cards[i], cards[j] = cards[j], cards[i] })
	}

	return ret, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
