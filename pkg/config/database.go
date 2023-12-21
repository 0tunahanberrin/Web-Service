package config

import (
	"context"
	"fmt"
	"time"
	"web_service_ko/pkg/helper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host     = "localhost"
	port     = 27017
	user     = ""
	password = ""
	dbName   = "test"
)

func MongoDatabaseConnection() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := fmt.Sprintf("mongodb://%s:%d", host, port)
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	helper.ErrorPanic(err)

	database := client.Database(dbName)
	return database
}
