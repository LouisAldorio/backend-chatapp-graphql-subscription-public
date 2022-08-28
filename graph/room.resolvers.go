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

func (r *roomResolver) Subscribers(ctx context.Context, obj *model.Room) ([]*model.Subscriber, error) {
	return dataloader.CtxLoaders(ctx).RoomSubscribers.Load(obj.ID)
}

func (r *roomOpsResolver) Create(ctx context.Context, obj *model.RoomOps, input model.NewRoom) (*model.Room, error) {
	subscriber := auth.ForContext(ctx)
	return service.RoomCreate(input, subscriber.SubscriberID)
}

func (r *roomOpsResolver) AddNewSubscriberToRoom(ctx context.Context, obj *model.RoomOps, roomID int) (*model.Room, error) {
	subscriber := auth.ForContext(ctx)
	return service.AddNewSubscriberToRoom(roomID, subscriber.SubscriberID)
}

func (r *roomQueryResolver) RoomsByLoggedInUser(ctx context.Context, obj *model.RoomQuery) ([]*model.Room, error) {
	subscriber := auth.ForContext(ctx)
	return service.RoomsGetBySubscriberId(subscriber.SubscriberID)
}

func (r *roomQueryResolver) AllRoomsExceptLoggedInUserRooms(ctx context.Context, obj *model.RoomQuery) ([]*model.Room, error) {
	subscriber := auth.ForContext(ctx)
	return service.RoomsGetByNotSubscriberId(subscriber.SubscriberID)
}

// Room returns generated.RoomResolver implementation.
func (r *Resolver) Room() generated.RoomResolver { return &roomResolver{r} }

// RoomOps returns generated.RoomOpsResolver implementation.
func (r *Resolver) RoomOps() generated.RoomOpsResolver { return &roomOpsResolver{r} }

// RoomQuery returns generated.RoomQueryResolver implementation.
func (r *Resolver) RoomQuery() generated.RoomQueryResolver { return &roomQueryResolver{r} }

type roomResolver struct{ *Resolver }
type roomOpsResolver struct{ *Resolver }
type roomQueryResolver struct{ *Resolver }
