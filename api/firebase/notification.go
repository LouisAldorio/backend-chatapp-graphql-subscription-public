package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"

	"google.golang.org/api/option"
)

func SendToToken() {

	ctx := context.Background()
	opt := option.WithCredentialsFile("subscription-chat-app-firebase-adminsdk-90j0u-02533c4dcc.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println(err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// This registration token comes from the client FCM SDKs.
	registrationToken := os.Getenv("FIREBASE_REGISTRATION_TOKEN")

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"title": "Whosapp",
			"body":  "Incoming Message",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}