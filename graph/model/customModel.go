package model

type Subscriber struct {
	ID             int             `json:"id" gorm:"column:id"`
	Name           string          `json:"name" gorm:"column:name"`
	Avatar         string          `json:"avatar" gorm:"column:avatar"`
	IsOnline       int             `json:"is_online" gorm:"column:is_online"`
	Email          string          `json:"email" gorm:"column:email"`
	HashedPassword string          `json:"hashed_password" gorm:"column:hashed_password"`
	Messages       chan []*Message `gorm:"-"`
}

type Message struct {
	ID           string `json:"id" bson:"id"`
	SubscriberID int    `json:"subcriber_id" bson:"subscriber_id"`
	Content      string `json:"content" bson:"content"`
	CreatedAt    string `json:"created_at" bson:"created_at"`
}

type MongoRoom struct {
	RoomID int        `bson:"roomd_id"`
	Chat   []*Message `bson:"chat"`
}

type Room struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Avatar      string        `json:"avatar"`
	Subscribers []*Subscriber `json:"subscribers" gorm:"-"`
}
