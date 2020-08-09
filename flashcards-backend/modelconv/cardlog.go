// Package modelconv includes helpers to transform models
package modelconv

import (
	"fmt"
	"strconv"

	"flashcards-backend/ent"
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/graph/model"
)

func cardResult(res cardlog.Result) model.CardResult {
	switch res {
	case cardlog.ResultAverage:
		return model.CardResultAverage
	case cardlog.ResultBad:
		return model.CardResultBad
	case cardlog.ResultGood:
		return model.CardResultGood
	case cardlog.ResultRetry:
		return model.CardResultRetry
	}
	panic(fmt.Sprintf("Not supported cardlogResult: %v", res))
}

// ModelCardResult converts model.CardResult to cardlog.Result
func ModelCardResult(res model.CardResult) cardlog.Result {
	switch res {
	case model.CardResultAverage:
		return cardlog.ResultAverage
	case model.CardResultBad:
		return cardlog.ResultBad
	case model.CardResultGood:
		return cardlog.ResultGood
	case model.CardResultRetry:
		return cardlog.ResultRetry
	}
	panic(fmt.Sprintf("Not supported cardlogResult: %v", res))
}

// CardLog converts ent.CardLog to model.CardLog
func CardLog(log *ent.CardLog) *model.CardLog {
	if log == nil {
		return nil
	}
	return &model.CardLog{
		ID:         strconv.Itoa(log.ID),
		LastResult: cardResult(log.Result),
		Word:       Word(log.Edges.Card),
		CreateTime: log.CreateTime,
	}
}

// CardLogS converts a slice of ent.CardLog to model.CardLog
func CardLogS(logs []*ent.CardLog) []*model.CardLog {
	models := make([]*model.CardLog, 0, len(logs))
	for _, card := range logs {
		models = append(models, CardLog(card))
	}
	return models
}
