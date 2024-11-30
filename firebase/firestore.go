package firebase

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func InitFirestore() {
	opt := option.WithCredentialsFile("config/secrets/carpooling-3ecca-firebase-adminsdk-icyva-d10cf5b5e3.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase: %v", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize Firestore: %v", err)
	}

	FirestoreClient = client
	log.Println("Firestore initialized successfully")
}