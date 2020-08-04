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
	}
	panic(fmt.Sprintf("Not supported cardlogResult: %v", res))
}

// CardLog converts ent.CardLog to model.CardLog
func CardLog(log *ent.CardLog) *model.CardLog {
	return &model.CardLog{
		ID:         strconv.Itoa(log.ID),
		LastResult: cardResult(log.Result),
		Word:       Word(log.Edges.Card),
	}
}
