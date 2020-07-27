// Package modelconv includes helpers to transform models
package modelconv

import (
	"strconv"

	"flashcards-backend/ent"
	"flashcards-backend/graph/model"
)

// Word converts ent.Word to model.Word
func Word(word *ent.Word) *model.Word {
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
