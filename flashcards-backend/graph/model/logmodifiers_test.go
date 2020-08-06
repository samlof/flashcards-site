package model

import (
	"testing"
)

func TestModifiersAllEnumHandled(t *testing.T) {
	for _, res := range AllCardResult {
		_, ok := LogModifiers[res]
		if !ok {
			t.Errorf("logModifier for %v wasn't defined", res)
		}
	}
}
