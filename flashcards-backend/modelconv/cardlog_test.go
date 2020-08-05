// Package modelconv includes helpers to transform models
package modelconv

import (
	"flashcards-backend/ent/cardlog"
	"flashcards-backend/graph/model"
	"testing"
)

func TestCardResult_HandlesAllEnums(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("%v", err)
		}
	}()
	for _, enum := range cardlog.AllResult {
		res := cardResult(enum)

		if !res.IsValid() {
			t.Errorf("%v returned invalid result: %v", enum, res)
		}
	}
}

func TestModelCardResult_HandlesAllEnums(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Errorf("%v", err)
		}
	}()

	validEnum := func(res cardlog.Result) bool {
		for _, enum := range cardlog.AllResult {
			if enum == res {
				return true
			}
		}
		return false
	}
	for _, enum := range model.AllCardResult {
		res := ModelCardResult(enum)

		if !validEnum(res) {
			t.Errorf("%v returned invalid result: %v", enum, res)
		}
	}
}
