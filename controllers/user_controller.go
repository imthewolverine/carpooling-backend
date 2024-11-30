package controllers

import (
	"context"
	"log"

	"github.com/imthewolverine/carpooling-backend/firebase"
)

func AddUserToFirestore(userID string, userData map[string]interface{}) {
	_, err := firebase.FirestoreClient.Collection("users").Doc(userID).Set(context.Background(), userData)
	if err != nil {
		log.Fatalf("Error adding user to Firestore: %v", err)
	}
	log.Println("User added to Firestore successfully!")
}
