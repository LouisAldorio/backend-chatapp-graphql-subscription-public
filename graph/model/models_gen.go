// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type InvitationOps struct {
	Invite           *PendingInvitation `json:"invite"`
	AcceptInvitation bool               `json:"accept_invitation"`
}

type InvitationQuery struct {
	SentInvitationRequest     []*PendingInvitation `json:"sent_invitation_request"`
	ReceivedInvitationRequest []*PendingInvitation `json:"received_invitation_request"`
}

type MessageOps struct {
	PostMessage string `json:"post_message"`
}

type MessagePagination struct {
	Page   int        `json:"page"`
	Limit  int        `json:"limit"`
	RoomID int        `json:"room_id"`
	Total  int        `json:"total"`
	Nodes  []*Message `json:"nodes"`
}

type MessageQuery struct {
	Messages *MessagePagination `json:"messages"`
}

type NewRoom struct {
	Name   string  `json:"name"`
	Avatar *string `json:"avatar"`
}

type NewSubscriber struct {
	Name     string  `json:"name"`
	Avatar   *string `json:"avatar"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
}

type PendingInvitation struct {
	ID         int         `json:"id"`
	InviterID  int         `json:"inviter_id"`
	ReceiverID int         `json:"receiver_id"`
	RoomID     int         `json:"room_id"`
	Inviter    *Subscriber `json:"inviter"`
	Receiver   *Subscriber `json:"receiver"`
	Room       *Room       `json:"room"`
}

type RoomOps struct {
	Create                 *Room `json:"create"`
	AddNewSubscriberToRoom *Room `json:"add_new_subscriber_to_room"`
}

type RoomQuery struct {
	RoomsByLoggedInUser             []*Room `json:"rooms_by_logged_in_user"`
	AllRoomsExceptLoggedInUserRooms []*Room `json:"all_rooms_except_logged_in_user_rooms"`
}

type SubscriberOps struct {
	Register *string `json:"register"`
	Login    *string `json:"login"`
}

type SubscriberQuery struct {
	OnlineSubscribers []*Subscriber `json:"online_subscribers"`
	Subscribers       []*Subscriber `json:"subscribers"`
}

type SubscribersHaveRooms struct {
	ID           int `json:"id"`
	RoomID       int `json:"room_id"`
	SubscriberID int `json:"subscriber_id"`
}
