package service

import (
	"context"
	"myapp/auth"
	"myapp/graph/model"
	"sync"

	"myapp/utils"
	"time"
)

type Broker struct {
	ChatMessages  map[int][]*model.Message
	ChatObservers map[int][]*model.Subscriber
	Mutex         sync.Mutex
}

func (r *Broker) New() *Broker {

	return &Broker{
		ChatMessages:  map[int][]*model.Message{},
		ChatObservers: map[int][]*model.Subscriber{},
	}
}

func (r *Broker) Subscribe(ctx context.Context, roomId int, subscriberId int) <-chan []*model.Message {

	msgs := make(chan []*model.Message, 1)
	r.ChatMessages[roomId] = []*model.Message{}

	go func() {
		<-ctx.Done()
		r.Mutex.Lock()
		r.ChatObservers[roomId] = utils.FilterSubscribers(r.ChatObservers[roomId], subscriberId)
		r.Mutex.Unlock()
	}()

	r.Mutex.Lock()
	r.ChatObservers[roomId] = append(r.ChatObservers[roomId], &model.Subscriber{
		ID:       subscriberId,
		Messages: msgs,
	})
	r.Mutex.Unlock()

	r.ChatObservers[roomId][utils.FindNewJoinedSubscribers(r.ChatObservers[roomId], subscriberId)].Messages <- r.ChatMessages[roomId]

	return msgs
}

func (r *Broker) PostNewMessage(ctx context.Context, subcriberId int, content string, roomId int) string {

	subscriber := auth.ForContext(ctx)

	newMessage := &model.Message{
		ID:           utils.GenerateUniqueID(),
		SubscriberID: subscriber.SubscriberID,
		Content:      content,
		CreatedAt:    time.Now().UTC().Format("2006-01-02 15:04:05"),
	}

	r.ChatMessages[roomId] = append(r.ChatMessages[roomId], newMessage)

	InsertChatToMongoDB(newMessage, roomId)

	r.Mutex.Lock()
	for _, observer := range r.ChatObservers[roomId] {
		observer.Messages <- r.ChatMessages[roomId]
	}
	r.Mutex.Unlock()

	return newMessage.ID
}
