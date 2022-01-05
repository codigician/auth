package auth

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/distribution/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct{}

func (m Mongo) Save(u *User) error {
	client, ctx, collection := m.connect()
	insertResult, err := collection.InsertOne(ctx, u)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single document:", insertResult.InsertedID)
	err = m.disconnect(client, ctx)
	return err
}

func (m Mongo) Find(email string) (*User, error) {
	client, ctx, collection := m.connect()
	var result *User
	filter := bson.M{"email": email}
	fmt.Println(email)
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Found a single document: %+v/n", result)
	err = m.disconnect(client, ctx)
	return result, err
}

func (m Mongo) Update(id uuid.UUID, key string, value string) error {
	client, ctx, collection := m.connect()
	filter := bson.M{"id": bson.M{"$eq": id}}
	updater := bson.M{"$set": bson.M{key: value}}
	result, err := collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return err
	}
	fmt.Println("RESULT:", result)
	err = m.disconnect(client, ctx)
	return err
}

func (m Mongo) List() {
	client, ctx, collection := m.connect()
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
	m.disconnect(client, ctx)
}

func (m Mongo) connect() (mongo.Client, context.Context, mongo.Collection) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	usersCollection := client.Database("auth").Collection("users")
	return *client, ctx, *usersCollection
}

func (m Mongo) disconnect(client mongo.Client, ctx context.Context) error {
	err := client.Disconnect(ctx)
	return err
}
