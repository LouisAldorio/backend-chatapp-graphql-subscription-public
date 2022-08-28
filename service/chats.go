package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetChatFromMongoDB(roomId int) []*model.Message {

	response := []*model.Message{}

	client := config.MongodbConnect()
	collection := client.Database("subscription_chat").Collection("rooms")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	findOptions := options.FindOneOptions{}
	findOptions.SetSort(bson.D{
		primitive.E{
			Key: "_id",
			Value:  -1,
		},
	})

	cur := collection.FindOne(ctx, bson.M{"room_id": roomId}, &findOptions)
	var result bson.M
	err := cur.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*model.Message{}
		}
	}

	for _, v := range result["chat"].(primitive.A) {
		currentChatObject := v.(primitive.M)

		subscriberId, err := strconv.Atoi(fmt.Sprintf("%v", currentChatObject["subscriber_id"]))
		if err != nil {
			panic(err)
		}
		response = append(response, &model.Message{
			ID:           fmt.Sprintf("%v", currentChatObject["id"]),
			SubscriberID: subscriberId,
			Content:      fmt.Sprintf("%v", currentChatObject["content"]),
			CreatedAt:    fmt.Sprintf("%v", currentChatObject["created_at"]),
		})
	}

	return response
}

func GetChatFromMongoDBWithPagination(page int, limit int, roomId int) []*model.Message {
	total := GetChatTotal(roomId)

	response := []*model.Message{}

	if total == 0 {
		return response
	}

	offset := total - (page * limit)
	preciseLimit := limit
	if offset < 0 {
		preciseLimit = limit + offset
		offset = 0		
	}

	client := config.MongodbConnect()
	collection := client.Database("subscription_chat").Collection("rooms")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	operationFindMatchingDocument := bson.D{
		primitive.E{
			Key: "$match",
			Value: bson.M{
				"room_id": roomId,
			},			 
		},
	}

	//$addfields has alias which is $set
	operationPagination := bson.D{
		primitive.E{
			Key: "$addFields",
			Value: bson.M{
				"chat": bson.M {
					"$slice": []interface{}{
						"$chat",
						offset,
						preciseLimit,
					},	
				},
			},
		},
	}

	pipeline := mongo.Pipeline{
		operationFindMatchingDocument, 
		operationPagination,
	}

	cur,err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		panic(err)
	}

	var result []bson.M
	err = cur.All(ctx, &result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return []*model.Message{}
		}
	}

	//pretty sure the result will only be one, cause we are getting it by room id
	roomChats := result[0]["chat"]

	for _,v := range roomChats.(primitive.A) {
		currentChatObject := v.(primitive.M)

		subscriberId, err := strconv.Atoi(fmt.Sprintf("%v", currentChatObject["subscriber_id"]))
		if err != nil {
			panic(err)
		}
		response = append(response, &model.Message{
			ID:           fmt.Sprintf("%v", currentChatObject["id"]),
			SubscriberID: subscriberId,
			Content:      fmt.Sprintf("%v", currentChatObject["content"]),
			CreatedAt:    fmt.Sprintf("%v", currentChatObject["created_at"]),
		})
	}

	return response
}


func GetChatTotal(roomId int) int {
	client := config.MongodbConnect()
	collection := client.Database("subscription_chat").Collection("rooms")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	findOptions := options.FindOneOptions{}
	findOptions.SetSort(bson.D{
		primitive.E{
			Key: "_id",
			Value:  -1,
		},
	})

	cur := collection.FindOne(ctx, bson.M{"room_id": roomId}, &findOptions)
	var result bson.M
	err := cur.Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0
		}
	}

	chats := result["chat"].(primitive.A)

	return len(chats)
}






func InsertChatToMongoDB(newMessage *model.Message, roomId int) {
	
	client := config.MongodbConnect()
	collection := client.Database("subscription_chat").Collection("rooms")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	upsert := false
	after := options.After
	findOptions := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	update := bson.M{
		"$push": bson.M{
			"chat": newMessage,
		},
	}

	cur := collection.FindOneAndUpdate(ctx, bson.M{
		"room_id": roomId,
	}, update, &findOptions)

	var result bson.M
	err := cur.Decode(&result)
	if err != nil {
		panic(err)
	}
}
