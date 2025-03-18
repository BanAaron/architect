package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BanAaron/architect/database"
	"github.com/joho/godotenv"
)

const (
	architect = "Architect"
	keywords  = "Keywords"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mongoURI, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		log.Fatal("MONGO_URI environment variable not set")
	}

	database.Connect(mongoURI)
	defer database.Disconnect()

	collection := database.GetCollection(architect, keywords)
	fmt.Printf(collection.Name())

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Hello, World!"))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/put", func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		id, err := collection.InsertOne(ctx, database.NewKeywordDocument("Aaron", []string{"Barratt"}))
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		log.Println(id)
		writer.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("/teapot", func(writer http.ResponseWriter, request *http.Request) {
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
		writer.WriteHeader(http.StatusTeapot)
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server")
	}
}
