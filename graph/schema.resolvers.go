package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/auth"
	"myapp/graph/generated"
	"myapp/graph/model"
)

func (r *mutationResolver) Message(ctx context.Context) (*model.MessageOps, error) {
	return &model.MessageOps{}, nil
}

func (r *mutationResolver) Room(ctx context.Context) (*model.RoomOps, error) {
	return &model.RoomOps{}, nil
}

func (r *mutationResolver) Subscriber(ctx context.Context) (*model.SubscriberOps, error) {
	return &model.SubscriberOps{}, nil
}

func (r *mutationResolver) Invitation(ctx context.Context) (*model.InvitationOps, error) {
	return &model.InvitationOps{}, nil
}

func (r *queryResolver) Message(ctx context.Context) (*model.MessageQuery, error) {
	return &model.MessageQuery{}, nil
}

func (r *queryResolver) Room(ctx context.Context) (*model.RoomQuery, error) {
	return &model.RoomQuery{}, nil
}

func (r *queryResolver) Subscriber(ctx context.Context) (*model.SubscriberQuery, error) {
	return &model.SubscriberQuery{}, nil
}

func (r *queryResolver) Invitation(ctx context.Context) (*model.InvitationQuery, error) {
	return &model.InvitationQuery{}, nil
}

func (r *subscriptionResolver) Messages(ctx context.Context, roomID int) (<-chan []*model.Message, error) {
	subscriber := auth.ForContext(ctx)
	return Broker.Subscribe(ctx, roomID, subscriber.SubscriberID), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
