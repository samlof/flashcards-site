package graph

import (
	"context"
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/graph/model"
	"testing"
)

func TestCardStatus(t *testing.T) {
	resolver := makeResolver(t)

	ctx := context.Background()
	newWord := model.NewWord{
		Lang1: "fi",
		Lang2: "en",
		Word1: "jäätelö",
		Word2: "icecream",
	}
	word, err := resolver.Mutation().CreateWord(ctx, newWord)
	if err != nil {
		t.Errorf("adding word: %v", err)
	}
	cardStatus := model.CardStatus{
		CardID: word.ID,
		Result: model.CardResultAverage,
	}
	ret, err := resolver.Mutation().CardStatus(ctx, cardStatus)
	if err != nil {
		t.Errorf("adding status: %v", ret)
	}
}

// Resolver depends on FK[0] being exactly this
func TestCardLogForeignKey(t *testing.T) {
	fk := cardlog.ForeignKeys[0]
	if fk != "card_log_card" {
		t.Errorf("cardlog -> card FK has changed. Should be %s, was %s", "card_log_card", fk)
	}
}
