package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
)

var (
	fcmClient *messaging.Client
)

// InitializeAppDefault will initialize firebase app
func InitializeAppDefault() *firebase.App {
	// you can directly set the serviceAccountKey json file here
	// opts := []option.ClientOption{option.WithCredentialsFile("serviceAccountKey.json")}

	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	//Firebase cloud messaging client
	fcmClient, err = app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Firebase messaging can't be initialized error: %v\n", err)
	}

	return app
}

// FcmClient returns a firebase cloud messaging client for authentication
func FcmClient() *messaging.Client {
	return fcmClient
}
