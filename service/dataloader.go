package service

import (
	"myapp/graph/model"
)


func SubscribersLoader(keys []int) ([]*model.Subscriber, []error) {
	subscribers := SubscriberGetByWhereInIDs(keys)

	subscribersMappedById := map[int]*model.Subscriber{}

	for _,v := range subscribers {
		subscribersMappedById[v.ID] = v
	}

	result := make([]*model.Subscriber, len(keys))
	for i,id := range keys {
		result[i] = subscribersMappedById[id]
	}

	return result, nil
}

func RoomSubscribersLoader(keys []int) ([][]*model.Subscriber, []error) {
	relations,err := RoomSubscribersLoaderGetRelationRoomsAndSubscribers(keys)
	if err != nil {
		panic(err)
	}

	subscribersIdsMappedByRoomId := map[int][]int{}
	subscribersIds := []int{}
	for _,v := range relations {
		subscribersIdsMappedByRoomId[v.RoomID] = append(subscribersIdsMappedByRoomId[v.RoomID], v.SubscriberID)
		subscribersIds = append(subscribersIds, v.SubscriberID)
	}

	subscribers := SubscriberGetByWhereInIDs(subscribersIds)
	subcribersMappedById := map[int]*model.Subscriber{}

	for _,v := range subscribers {
		subcribersMappedById[v.ID] = v
	}

	result := make([][]*model.Subscriber, len(keys))
	for i, id := range keys {
		temp := []*model.Subscriber{}
		for _,v := range subscribersIdsMappedByRoomId[id] {
			temp = append(temp, subcribersMappedById[v])
		}

		result[i] = temp
	}

	return result, nil
}


func RoomLoader(keys []int) ([]*model.Room, []error){
	rooms := RoomGetByWhereInID(keys)

	roomMappedById := map[int]*model.Room{}
	for _,v := range rooms {
		roomMappedById[v.ID] = v
	}

	result := make([]*model.Room, len(keys))
	for i,id := range keys {
		result[i] = roomMappedById[id]
	}

	return result, nil
}