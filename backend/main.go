package main

import (
	"context"
	"log"
	"os"
	"time"

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
	kw := database.NewKeywordDocument("Aaron", []string{"Ford", "BMW", "Audi"})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err = collection.InsertOne(ctx, kw)
	if err != nil {
		log.Printf("Failed to insert row: %s", kw)
	}

	rows := []*database.KeywordDocument{}
	rows = append(rows, kw)
}
