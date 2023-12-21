package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() *mongo.Client{
	if err:= godotenv.Load(); err != nil{
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri ==""{
		log.Fatal("Sorry URI is blanck!")
	}
	// Set client options
	clientOption := options.Client().ApplyURI(uri)

	//Creating mongo connection
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil{
		log.Fatal(err)
	}

	// Check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}