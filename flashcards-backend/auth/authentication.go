package auth

import (
	"context"
	"log"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"uid"}

type contextKey struct {
	name string
}

// A stand-in for our database backed user object
type User struct {
	Name    string
	IsAdmin bool
}

// Middleware decodes the share session cookie and packs the user uid into context
func Middleware(firebaseAuth *auth.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			hs, ok := r.Header["Authorization"]
			if !ok || len(hs) == 0 {
				next.ServeHTTP(w, r)
				return
			}
			h := hs[0]

			if !strings.HasPrefix(h, "Bearer ") {
				next.ServeHTTP(w, r)
				return
			}
			idToken := h[len("Bearer "):]

			ctx := r.Context()
			decToken, err := firebaseAuth.VerifyIDToken(ctx, idToken)
			if err != nil {
				log.Printf("Middleware error validating token %v", err)
				next.ServeHTTP(w, r)
				return
			}
			userUid := decToken.UID

			// put it in context
			ctx = context.WithValue(ctx, userCtxKey, userUid)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user uid from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}
