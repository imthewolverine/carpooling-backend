package main

import (
	"log"

	"github.com/imthewolverine/carpooling-backend/firebase"
)

func main() {
	// Initialize Firestore
	firebase.InitFirestore()
	defer firebase.FirestoreClient.Close()

	// Initialize Realtime Database
	firebase.InitRealtimeDatabase()

	// Start your application
	log.Println("Application is running...")
}
