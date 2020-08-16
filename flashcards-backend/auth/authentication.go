package auth

import (
	"context"
	"flashcards-backend/ent"
	"flashcards-backend/ent/user"
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}
var errorCtxKey = &contextKey{"error"}

type contextKey struct {
	name string
}

// Middleware decodes the share session cookie and packs the user uid into context
func Middleware(firebaseAuth *auth.Client, client *ent.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			hs, ok := r.Header["Authorization"]
			if !ok || len(hs) == 0 {
				r = setError(r, ctx, "missing Authorization header")
				next.ServeHTTP(w, r)
				return
			}
			h := hs[0]

			if !strings.HasPrefix(h, "Bearer ") {
				r = setError(r, ctx, "missing bearer")
				next.ServeHTTP(w, r)
				return
			}
			idToken := h[len("Bearer "):]

			decToken, err := firebaseAuth.VerifyIDToken(ctx, idToken)
			if err != nil {
				r = setError(r, ctx, "middleware error validating token %v", err)
				next.ServeHTTP(w, r)
				return
			}
			userUid := decToken.UID

			// put it in context
			dbUser, err := client.User.Query().Where(user.FirebaseUid(userUid)).First(ctx)
			if ent.IsNotFound(err) {
				dbUser, err = client.User.Create().SetFirebaseUid(userUid).Save(ctx)
				if err != nil {
					r = setError(r, ctx, "error creating user for uid %s: %v", userUid, err)
					next.ServeHTTP(w, r)
					return
				}
			} else if err != nil {
				r = setError(r, ctx, "error getting user with uid %s: %v", userUid, err)
				next.ServeHTTP(w, r)
				return
			}
			settingsExists, err := dbUser.QuerySettings().Exist(ctx)
			if err != nil {
				r = setError(r, ctx, "error getting settings for uid %s: %v", userUid, err)
				next.ServeHTTP(w, r)
				return
			}
			if !settingsExists {
				_, err := client.UserSettings.Create().SetUser(dbUser).Save(ctx)
				if err != nil {
					r = setError(r, ctx, "error creating settings for uid %s: %v", userUid, err)
					next.ServeHTTP(w, r)
					return
				}
			}

			ctx = context.WithValue(ctx, userCtxKey, dbUser)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func setError(r *http.Request, ctx context.Context, msg string, a ...interface{}) *http.Request {
	ctx = context.WithValue(ctx, errorCtxKey, fmt.Errorf(msg, a...))
	return r.WithContext(ctx)
}

// ForContext finds the user uid from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *ent.User {
	raw, _ := ctx.Value(userCtxKey).(*ent.User)
	return raw
}

// ForContextErr finds the happened error from the context.
func ForContextErr(ctx context.Context) error {
	raw, _ := ctx.Value(errorCtxKey).(error)
	return raw
}
