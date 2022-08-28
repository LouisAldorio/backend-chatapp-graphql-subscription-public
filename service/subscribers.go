package service

import (
	"myapp/config"
	"myapp/graph/model"

	"myapp/auth"
)

func RoomSubscribersLoaderGetRelationRoomsAndSubscribers(roomIds []int) ([]*model.SubscribersHaveRooms, error) {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var rooms []*model.SubscribersHaveRooms
	err := db.Table("subscribers_have_rooms").Where("room_id IN (?)", roomIds).Find(&rooms).Error
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func SubscribersGet(query string) ([]*model.Subscriber, error){
	db,sql := config.ConnectGorm()
	defer sql.Close()

	var subscribers []*model.Subscriber

	tx := db.Table("subscribers")

	if len(query) > 0 {
		tx = tx.Where("name LIKE ?","%" + query + "%")
	}

	err := tx.Find(&subscribers).Error
	if err != nil {
		return nil, err
	}

	return subscribers, nil
}


func SubscriberGetByWhereInIDs(ids []int) []*model.Subscriber {
	db,sql := config.ConnectGorm()
	defer sql.Close()

	var subscribers []*model.Subscriber

	err := db.Table("subscribers").Where("id IN (?)", ids).Find(&subscribers).Error
	if err != nil {
		panic(err)
	}

	return subscribers
}

func SubcriberCreate(input model.NewSubscriber) *model.Subscriber {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	defaultAvatar := "https://cdn.pixabay.com/photo/2015/10/05/22/37/blank-profile-picture-973460_640.png"
	if input.Avatar != nil {
		defaultAvatar = *input.Avatar
	}

	newSubscriber := model.Subscriber {
		Name: input.Name,
		Avatar: defaultAvatar,
		IsOnline: 2,
		Email: input.Email,
		HashedPassword: auth.HashPassword(input.Password),
	}

	err := db.Table("subscribers").Create(&newSubscriber).Error
	if err != nil {
		panic(err)
	}

	return &newSubscriber
}

func GetOnlineSubscribers() []*model.Subscriber {
	db, sql := config.ConnectGorm()
	defer sql.Close()

	var subscribers []*model.Subscriber
	err := db.Table("subscribers").Where("is_online = ?", 1).Find(&subscribers).Error
	if err != nil {
		panic(err)
	}

	return subscribers
}
