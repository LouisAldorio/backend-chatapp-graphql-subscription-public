package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler/transport"
)

var userCtxKey = &contextKey{"subscriber"}

type contextKey struct {
	subscriberIdKey string
}

func WebSocketMiddleware(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {

	authToken, tokenErr := initPayload["Authorization"].(string)
	if !tokenErr {
		return context.Background(), fmt.Errorf("token please")
	}

	jwtToken, err := ValidateToken(authToken)
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*SubscriberClaim)
	if !ok && !jwtToken.Valid {
		return nil, err
	}

	userCtx := context.WithValue(ctx, userCtxKey, claims)

	// and return it so the resolvers can see it
	return userCtx, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get("Authorization")

		// Allow unauthenticated users in
		if authToken == "" {
			next.ServeHTTP(w, r)
			return
		}

		//validate jwt token
		jwtToken, err := ValidateToken(authToken)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		//validate claim
		claims, ok := jwtToken.Claims.(*SubscriberClaim)
		if !ok && !jwtToken.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		//return user data to req
		ctx := context.WithValue(r.Context(), userCtxKey, claims)
		reqWithCtx := r.WithContext(ctx)
		next.ServeHTTP(w, reqWithCtx)
	})

}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *SubscriberClaim {
	raw, _ := ctx.Value(userCtxKey).(*SubscriberClaim)
	return raw
}
