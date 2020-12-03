package domain

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection = "mongodb://localhost:27017"

const dbName = "go-todo"
const collectionName = "todoList"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connection)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// mongo connection
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	if error := client.Ping(ctx, nil); err != nil {
		log.Fatal(error)
	}

	fmt.Println("Connected to MongoDB =D")

	collection = client.Database(dbName).Collection(collectionName)
	color := "\033[32m"
	colorReset := "\033[0m"
	fmt.Printf("Collection %s %s %s instance created!\n", color, collectionName, colorReset)
}

func GetAll() []bson.D {
	cxt := context.TODO()
	filter := bson.D{{}}
	cursor, err := collection.Find(cxt, filter)

	defer cursor.Close(cxt)

	if err != nil {
		panic(err)
	}

	var result []bson.D

	if err := cursor.All(cxt, &result); err != nil {
		panic(err)
	}

	return result
}
