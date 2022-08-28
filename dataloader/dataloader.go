package dataloader

import (
	"context"
	"myapp/service"
	"net/http"
	"time"
)

type ctxKeyType string

const (
	ctxKey ctxKeyType = "dataloaders"
)

type loaders struct {
	Subscribers     *SubscribersLoader
	RoomSubscribers *RoomSubscribersLoader
	Room            *RoomLoader
}

func DataloaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}
		wait := 5 * time.Millisecond
		max := 100

		ldrs.Subscribers = &SubscribersLoader{
			wait:     wait,
			maxBatch: max,
			fetch:    service.SubscribersLoader,
		}

		ldrs.RoomSubscribers = &RoomSubscribersLoader{
			wait:     wait,
			maxBatch: max,
			fetch:    service.RoomSubscribersLoader,
		}

		ldrs.Room = &RoomLoader{
			wait:     wait,
			maxBatch: max,
			fetch:    service.RoomLoader,
		}

		dataloaderCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dataloaderCtx))
	})
}

func CtxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
