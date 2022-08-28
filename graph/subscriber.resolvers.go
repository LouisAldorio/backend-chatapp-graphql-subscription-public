package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *subscriberOpsResolver) Register(ctx context.Context, obj *model.SubscriberOps, input model.NewSubscriber) (*string, error) {
	return service.Register(input)
}

func (r *subscriberOpsResolver) Login(ctx context.Context, obj *model.SubscriberOps, email string, password string) (*string, error) {
	return service.Login(email, password)
}

func (r *subscriberQueryResolver) OnlineSubscribers(ctx context.Context, obj *model.SubscriberQuery) ([]*model.Subscriber, error) {
	return service.GetOnlineSubscribers(), nil
}

func (r *subscriberQueryResolver) Subscribers(ctx context.Context, obj *model.SubscriberQuery, query string) ([]*model.Subscriber, error) {
	return service.SubscribersGet(query)
}

// SubscriberOps returns generated.SubscriberOpsResolver implementation.
func (r *Resolver) SubscriberOps() generated.SubscriberOpsResolver { return &subscriberOpsResolver{r} }

// SubscriberQuery returns generated.SubscriberQueryResolver implementation.
func (r *Resolver) SubscriberQuery() generated.SubscriberQueryResolver {
	return &subscriberQueryResolver{r}
}

type subscriberOpsResolver struct{ *Resolver }
type subscriberQueryResolver struct{ *Resolver }
