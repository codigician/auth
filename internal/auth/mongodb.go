package auth

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct{}

func (m Mongo) Save(u *User) error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	collection := client.Database("auth").Collection("users")

	insertResult, err := collection.InsertOne(context.TODO(), u)
	fmt.Println("inserted a single document:", insertResult.InsertedID)

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to mongodb closed")

	return err
}

func (m Mongo) Find(e string) (*User, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	fmt.Println("connected to mongodb")
	collection := client.Database("auth").Collection("users")

	var result *User
	filter := bson.M{"email": e}
	fmt.Println(e)
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Found a single document: %+v/n", result)
	fmt.Printf(result.Firstname, result.Lastname, result.Email, result.HashedPassword)

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to mongodb closed")

	return result, nil
}
