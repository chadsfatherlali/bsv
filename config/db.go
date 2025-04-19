package config

import (
	"blockchain_votation_system/constants"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func InitDB() {
	dbUrl := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))

	if err != nil {
		log.Fatal("Mongo Connect Error:", err)
	}

	Client = client
	DB = client.Database(dbName)
}
