// Package modelconv includes helpers to transform models
package modelconv

import (
	"flashcards-backend/ent/cardlog"
	"testing"
)

func Test_cardResultHandlesAllEnums(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("%v", err)
		}
	}()
	for _, enum := range cardlog.ResultAll() {
		cardResult(enum)
	}
}
