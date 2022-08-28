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

func (r *invitationOpsResolver) Invite(ctx context.Context, obj *model.InvitationOps, roomID int, receiverID int) (*model.PendingInvitation, error) {
	subscriber := auth.ForContext(ctx)
	return service.Invite(receiverID, roomID, subscriber.SubscriberID)
}

func (r *invitationOpsResolver) AcceptInvitation(ctx context.Context, obj *model.InvitationOps, invitationID int, roomID int) (bool, error) {
	subscriber := auth.ForContext(ctx)
	return service.AcceptInvitation(roomID, invitationID, subscriber.SubscriberID)
}

func (r *invitationQueryResolver) SentInvitationRequest(ctx context.Context, obj *model.InvitationQuery) ([]*model.PendingInvitation, error) {
	subscriber := auth.ForContext(ctx)
	return service.SendInvitationRequestGet(subscriber.SubscriberID)
}

func (r *invitationQueryResolver) ReceivedInvitationRequest(ctx context.Context, obj *model.InvitationQuery) ([]*model.PendingInvitation, error) {
	subscriber := auth.ForContext(ctx)
	return service.ReceivedInvitationRequestGet(subscriber.SubscriberID)
}

func (r *pendingInvitationResolver) Inviter(ctx context.Context, obj *model.PendingInvitation) (*model.Subscriber, error) {
	return dataloader.CtxLoaders(ctx).Subscribers.Load(obj.InviterID)
}

func (r *pendingInvitationResolver) Receiver(ctx context.Context, obj *model.PendingInvitation) (*model.Subscriber, error) {
	return dataloader.CtxLoaders(ctx).Subscribers.Load(obj.ReceiverID)
}

func (r *pendingInvitationResolver) Room(ctx context.Context, obj *model.PendingInvitation) (*model.Room, error) {
	return dataloader.CtxLoaders(ctx).Room.Load(obj.RoomID)
}

// InvitationOps returns generated.InvitationOpsResolver implementation.
func (r *Resolver) InvitationOps() generated.InvitationOpsResolver { return &invitationOpsResolver{r} }

// InvitationQuery returns generated.InvitationQueryResolver implementation.
func (r *Resolver) InvitationQuery() generated.InvitationQueryResolver {
	return &invitationQueryResolver{r}
}

// PendingInvitation returns generated.PendingInvitationResolver implementation.
func (r *Resolver) PendingInvitation() generated.PendingInvitationResolver {
	return &pendingInvitationResolver{r}
}

type invitationOpsResolver struct{ *Resolver }
type invitationQueryResolver struct{ *Resolver }
type pendingInvitationResolver struct{ *Resolver }
