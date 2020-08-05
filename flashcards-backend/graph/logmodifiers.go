package graph

import (
	"flashcards-backend/graph/model"
)

// logModifiers maps result into a modifier to be used with scheduled_next
var logModifiers map[model.CardResult]float64 = map[model.CardResult]float64{
	model.CardResultAverage: 1.5,
	model.CardResultBad:     0.75,
	model.CardResultGood:    2,
}
