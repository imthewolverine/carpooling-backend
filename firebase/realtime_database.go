package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var RealtimeDatabaseClient *db.Client

func InitRealtimeDatabase() {
	opt := option.WithCredentialsFile("config/secrets/carpooling-3ecca-firebase-adminsdk-icyva-d10cf5b5e3.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize Realtime Database: %v", err)
	}

	RealtimeDatabaseClient = client
	log.Println("Realtime Database initialized successfully")
}
