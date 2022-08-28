package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/auth"
	"myapp/dataloader"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *messageResolver) Subscriber(ctx context.Context, obj *model.Message) (*model.Subscriber, error) {
	return dataloader.CtxLoaders(ctx).Subscribers.Load(obj.SubscriberID)
}

func (r *messageOpsResolver) PostMessage(ctx context.Context, obj *model.MessageOps, content string, roomID int) (string, error) {
	subscriber := auth.ForContext(ctx)
	return Broker.PostNewMessage(ctx, subscriber.SubscriberID, content, roomID), nil
}

func (r *messagePaginationResolver) Total(ctx context.Context, obj *model.MessagePagination) (int, error) {
	return service.GetChatTotal(obj.RoomID), nil
}

func (r *messagePaginationResolver) Nodes(ctx context.Context, obj *model.MessagePagination) ([]*model.Message, error) {
	return service.GetChatFromMongoDBWithPagination(obj.Page, obj.Limit, obj.RoomID), nil
}

func (r *messageQueryResolver) Messages(ctx context.Context, obj *model.MessageQuery, page int, limit int, roomID int) (*model.MessagePagination, error) {
	return &model.MessagePagination{
		Page:   page,
		Limit:  limit,
		RoomID: roomID,
	}, nil
}

// Message returns generated.MessageResolver implementation.
func (r *Resolver) Message() generated.MessageResolver { return &messageResolver{r} }

// MessageOps returns generated.MessageOpsResolver implementation.
func (r *Resolver) MessageOps() generated.MessageOpsResolver { return &messageOpsResolver{r} }

// MessagePagination returns generated.MessagePaginationResolver implementation.
func (r *Resolver) MessagePagination() generated.MessagePaginationResolver {
	return &messagePaginationResolver{r}
}

// MessageQuery returns generated.MessageQueryResolver implementation.
func (r *Resolver) MessageQuery() generated.MessageQueryResolver { return &messageQueryResolver{r} }

type messageResolver struct{ *Resolver }
type messageOpsResolver struct{ *Resolver }
type messagePaginationResolver struct{ *Resolver }
type messageQueryResolver struct{ *Resolver }
