package model

import (
	"context"
	"log"
	"web_service_ko/pkg/config"

	"go.mongodb.org/mongo-driver/bson"
)

type Book struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Project     *Project `json:"project"`
}

type Project struct {
	ID            int    `json:"projid"`
	Title         string `json:"projtitle"`
	Description   string `json:"projdescription"`
	ProjectStatus string `json:"projstatus"`
}
type Tags struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

var COLL = config.Connection().Database("book_project").Collection("details")

func (e *Book) CreateBookDetail() *Book {
	_, err := COLL.InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal(err)
	}
	return e
}

func ShowAllBookDetails() []Book {
	cursor, err := COLL.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var Books []Book
	if err = cursor.All(context.TODO(), &Books); err != nil {
		log.Fatal(err)
	}

	return Books
}

func ShowBookDetail(Id string) *Book {
	var Books Book
	cursor := COLL.FindOne(context.TODO(), bson.M{"bookid": Id})
	cursor.Decode(&Books)
	return &Books
}

func (e *Book) UpdateBookDetail(Id string) *Book {
	var Books Book
	update := bson.M{
		"$set": e,
	}
	_, err := COLL.UpdateOne(context.TODO(), bson.M{"bookid": Id}, update)
	if err != nil {
		panic(err)
	}
	cursor1 := COLL.FindOne(context.TODO(), bson.M{"bookid": Id})
	cursor1.Decode(&Books)
	return &Books
}

func DeleteBookDetail(Id string) []Book {
	_, err := COLL.DeleteOne(context.TODO(), bson.M{"bookid": Id})
	if err != nil {
		panic(err)
	}
	cursor, err := COLL.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var Books []Book
	if err = cursor.All(context.TODO(), &Books); err != nil {
		log.Fatal(err)
	}

	return Books
}
