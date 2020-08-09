// Package modelconv includes helpers to transform models
package modelconv

import (
	"flashcards-backend/ent"
	"flashcards-backend/graph/model"
)

func UserSettings(settings *ent.UserSettings) *model.UserSettings {
	return &model.UserSettings{
		NewCardsPerDay: settings.NewCardsPerDay,
	}
}
