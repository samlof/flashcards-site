package auth

import (
	"context"
	"flashcards-backend/ent"
	"math/rand"
	"strconv"
)

// SetUser sets a user into context. Used with tests
func SetUser(ctx context.Context, client *ent.Client) (context.Context, *ent.User) {
	randomId := strconv.Itoa(rand.Intn(100000))
	user := client.User.Create().SetEmail("test@test.email").SetFirebaseUid("randomuid" + randomId).SaveX(ctx)
	ctx = context.WithValue(ctx, userCtxKey, user)
	
	return ctx, user
}