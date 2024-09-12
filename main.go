package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/SHresTho12/rag_go/db"
	"github.com/joho/godotenv"
)

func init() {
	// Load the environment file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Unable to Load the env file.", err)
	}
}

func main() {
	// Retrieve environment variables
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_DB")

	if uri == "" || dbName == "" {
		log.Fatal("Environment variables for MongoDB URI or database name are not set.")
	}

	// Connect to the MongoDB database
	client, err := db.Connect(uri)
	if err != nil {
		log.Fatal(err)
	}

	// Get the database
	database := db.NewDatabase(client, dbName)
	if database == nil {
		log.Fatal("Unable to get the database.")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
