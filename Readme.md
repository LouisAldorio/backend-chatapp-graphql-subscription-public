1. Create Subscriber
2. Create Room
3. Insert Subscriber to Room
4. Start Chatting


<!-- setProjection take care of the pagination in array held by the document key
findOptions := options.FindOneOptions{}
findOptions.SetSort(bson.D{{"_id", -1}})
findOptions.SetProjection(bson.M{
 	"chat": bson.M{
 		"$slice": []int{
 			offset,
 			limit,
 		},		
 	},
})
cur := collection.FindOne(ctx, bson.M{
 	"room_id": roomId,
}, &findOptions) -->



<!-- pipeline operation used to reverse array
operationReverseArrayOfFoundDocumentField := bson.D{
    primitive.E{
        Key: "$addFields",
        Value: bson.M{
            "chat": bson.M{
                "$reverseArray": "$chat",
            },
        },
    },
} -->