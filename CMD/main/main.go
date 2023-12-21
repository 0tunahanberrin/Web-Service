package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
	"web_service_ko/pkg/helper"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var wg sync.WaitGroup

type Book struct {
	ID          string   `json:"id"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Project     *Project `json:"project"`
}

type Project struct {
	ProjID          string `json:"projid"`
	ProjTitle       string `json:"projtitle"`
	ProjDescription string `json:"projdescription"`
	ProjStatus      string `json:"projstatus"`
}

func init() {
	// MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	helper.ErrorPanic(err)

	database := client.Database("booksDB")
	collection = database.Collection("books")
}

func MongoDatabaseConnection() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// MongoDB: Get all books
	cursor, err := collection.Find(context.Background(), bson.M{})
	helper.ErrorPanic(err)
	defer cursor.Close(context.Background())

	var books []Book
	err = cursor.All(context.Background(), &books)
	helper.ErrorPanic(err)

	json.NewEncoder(w).Encode(books)
}

func PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	helper.ErrorPanic(err)

	// Add data to MongoDB concurrently using Goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()

		insertResult, err := collection.InsertOne(context.Background(), newBook)
		if err != nil {
			if writeErr, ok := err.(mongo.WriteException); ok {
				if len(writeErr.WriteErrors) > 0 {
					for _, writeError := range writeErr.WriteErrors {
						switch writeError.Code {
						case 11000:
							http.Error(w, "Duplicate key error. Book with the same ID already exists.", http.StatusBadRequest)
							return
						default:
							http.Error(w, fmt.Sprintf("Error inserting data into MongoDB: %v", writeError), http.StatusInternalServerError)
							return
						}
					}
				}
			}

			http.Error(w, fmt.Sprintf("Error inserting data into MongoDB: %v", err), http.StatusInternalServerError)
			return
		}

		insertedID := insertResult.InsertedID

		response := map[string]interface{}{"message": "Book received successfully", "insertedID": insertedID}
		json.NewEncoder(w).Encode(response)
	}()

	wg.Wait()
}
func main() {
	r := gin.Default()

	r.GET("/book", func(c *gin.Context) {
		GetBook(c.Writer, c.Request)
	})

	r.POST("/book", func(c *gin.Context) {
		PostBook(c.Writer, c.Request)
	})

	fmt.Println("Server running on: http://localhost:8000")
	log.Fatal(r.Run(":8000"))
}
