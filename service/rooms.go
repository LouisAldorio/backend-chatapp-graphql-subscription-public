package service

import (
	"context"
	"myapp/api"
	"myapp/config"
	"myapp/graph/model"
	"myapp/utils"
	"time"

	"github.com/vektah/gqlparser/gqlerror"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

func RoomsGetBySubscriberId(subscriberId int) ([]*model.Room, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var rooms []*model.Room
	relations, err := GetSubsciberAndRoomRelation(subscriberId)
	if err != nil {
		return nil, err
	}

	roomIds := []int{}
	for _, v := range relations {
		roomIds = append(roomIds, v.RoomID)
	}

	err = db.Table("rooms").Where("id IN (?)", utils.UniqueIntSlice(roomIds)).Find(&rooms).Error

	return rooms, err
}

func RoomsGetByNotSubscriberId(subscriberId int) ([]*model.Room, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var rooms []*model.Room
	relations, err := GetSubsciberAndRoomRelation(subscriberId)
	if err != nil {
		return nil, err
	}

	roomIds := []int{}
	for _, v := range relations {
		roomIds = append(roomIds, v.RoomID)
	}

	err = db.Table("rooms").Where("id NOT IN (?)", utils.UniqueIntSlice(roomIds)).Find(&rooms).Error

	return rooms, err
}

func GetSubsciberAndRoomRelation(subscriberId int) ([]*model.SubscribersHaveRooms, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var rooms []*model.SubscribersHaveRooms
	err := db.Table("subscribers_have_rooms").Where("subscriber_id = ?", subscriberId).Find(&rooms).Error
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func RoomGetById(id int) (*model.Room, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var room model.Room

	err := db.Table("rooms").Where("id = ?", id).First(&room).Error

	return &room, err
}

func RoomCreate(input model.NewRoom, subscriberId int) (*model.Room, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	if err := db.Table("rooms").Where("name = ?", input.Name).First(&model.Room{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			defaultAvatar := api.RandomOneImageFromPicsum()
			if input.Avatar != nil {
				defaultAvatar = *input.Avatar
			}

			newRoom := model.Room{
				Name:   input.Name,
				Avatar: defaultAvatar,
			}

			err := db.Table("rooms").Create(&newRoom).Error
			if err != nil {
				return &model.Room{}, err
			}

			MongoRoomCreate(newRoom.ID)

			_, err = AddNewSubscriberToRoom(newRoom.ID, subscriberId)
			if err != nil {
				return nil, err
			}

			return RoomGetById(newRoom.ID)
		}
	}

	return &model.Room{}, &gqlerror.Error{
		Message: "Room With the following name has already been taken!",
	}
}

func MongoRoomCreate(roomId int) {

	client := config.MongodbConnect()
	collection := client.Database("subscription_chat").Collection("rooms")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, bson.D{
		primitive.E{
			Key:   "room_id",
			Value: roomId,
		},
		primitive.E{
			Key:   "chat",
			Value: []*model.Message{},
		},
	})

	if err != nil {
		panic(err)
	}
}

func AddNewSubscriberToRoom(roomId int, subscriberId int) (*model.Room, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	err := db.Table("subscribers_have_rooms").Create(map[string]interface{}{
		"room_id":       roomId,
		"subscriber_id": subscriberId,
	}).Error

	if err != nil {
		panic(err)
	}

	return RoomGetById(roomId)
}

func RoomGetByWhereInID(roomIds []int) []*model.Room {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var rooms []*model.Room
	err := db.Table("rooms").Where("id IN (?)", roomIds).Find(&rooms).Error
	if err != nil {
		panic(err)
	}

	return rooms
}
