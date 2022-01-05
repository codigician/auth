package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) Mongo {
	return Mongo{
		collection: collection,
	}
}

func (m *Mongo) Save(ctx context.Context, u *User) error {
	result, err := m.collection.InsertOne(ctx, u)
	if err != nil {
		return err
	}
	fmt.Println("inserted a single document:", result)
	return nil
}

func (m *Mongo) Get(ctx context.Context, email string) (*User, error) {
	var user *User
	filter := bson.M{"email": email}
	fmt.Println(email)
	err := m.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Found a single document: %+v/n", user)
	return user, err
}

func (m *Mongo) Update(ctx context.Context, id string, fields map[string]interface{}) error {
	filter := bson.M{"_id": id}
	updater := bson.M{"$set": fields}
	result, err := m.collection.UpdateOne(ctx, filter, updater)
	if err != nil {
		return err
	}
	fmt.Println("RESULT:", result)
	return err
}

func (m *Mongo) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	result, err := m.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Printf("Delete result: %v, Delete count: %v\n", result, result.DeletedCount)
	return err
}
