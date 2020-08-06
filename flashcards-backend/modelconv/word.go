// Package modelconv includes helpers to transform models
package modelconv

import (
	"strconv"

	"flashcards-backend/ent"
	"flashcards-backend/graph/model"
)

// Word converts ent.Word to model.Word
func Word(word *ent.Word) *model.Word {
	if word == nil {
		return nil
	}
	return &model.Word{
		ID:         strconv.Itoa(word.ID),
		Lang1:      word.Lang1,
		Lang2:      word.Lang2,
		Word1:      word.Word1,
		Word2:      word.Word2,
		CreateTime: word.CreateTime,
		UpdateTime: word.CreateTime,
	}
}

// WordS converts a slice of ent.Word to model.Word
func WordS(words []*ent.Word) []*model.Word {
	models := make([]*model.Word, 0, len(words))
	for _, card := range words {
		models = append(models, Word(card))
	}
	return models
}
