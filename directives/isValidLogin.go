package directives

import (
	"context"
	"fmt"
	"myapp/auth"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func IsValidLogin(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {

	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("Access denied, Please Login to continue!")
	}

	currentTime := time.Now().UTC().UnixNano() / int64(time.Millisecond)

	if user.ExpiresAt <= currentTime {
		return nil, fmt.Errorf("Session Expired")
	}
	
	return next(ctx)
}

