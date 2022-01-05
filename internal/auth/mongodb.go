package auth

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/distribution/uuid"
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

func (m Mongo) Update(id uuid.UUID, key string, value string) error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	collection := client.Database("auth").Collection("users")

	filter := bson.M{"id": bson.M{"$eq": id}}
	updater := bson.M{"$set": bson.M{key: value}}
	result, err := collection.UpdateOne(context.TODO(), filter, updater)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("RESULT:", result)
	return nil
}

func (m Mongo) List() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")
	collection := client.Database("auth").Collection("users")

	var results []*User
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		var user User
		if err = cursor.Decode(&user); err != nil {
			log.Fatal(err)
		}
		results = append(results, &user)
	}
	cursor.Close(context.TODO())

	for _, r := range results {
		fmt.Println("One user:", r.ID, r.Firstname, r.Lastname, r.Email, r.HashedPassword)
	}
}
