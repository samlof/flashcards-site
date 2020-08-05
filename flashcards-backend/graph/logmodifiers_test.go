package graph

import (
	"flashcards-backend/graph/model"
	"testing"
)

func TestModifiersAllEnumHandled(t *testing.T) {
	for _, res := range model.AllCardResult {
		_, ok := logModifiers[res]
		if !ok {
			t.Errorf("logModifier for %v wasn't defined", res)
		}
	}
}
