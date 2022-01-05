package mongo

import (
	"context"
	"fmt"

	"github.com/codigician/auth/internal/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) *Mongo {
	return &Mongo{collection}
}

func (m *Mongo) Save(ctx context.Context, u *auth.User) error {
	result, err := m.collection.InsertOne(ctx, User{
		ID:             primitive.NewObjectID(),
		Firstname:      u.Firstname,
		Lastname:       u.Lastname,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
	})
	fmt.Println("inserted a single document:", result)
	return err
}

func (m *Mongo) Get(ctx context.Context, email string) (*auth.User, error) {
	var user User
	filter := bson.M{"email": email}
	err := m.collection.FindOne(ctx, filter).Decode(&user)
	return &auth.User{
		ID:             user.ID.Hex(),
		Firstname:      user.Firstname,
		Lastname:       user.Lastname,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
	}, err
}

func (m *Mongo) Update(ctx context.Context, id string, fields map[string]interface{}) error {
	filter := bson.M{"_id": id}
	updater := bson.M{"$set": fields}
	result, err := m.collection.UpdateOne(ctx, filter, updater)
	fmt.Println("RESULT:", result)
	return err
}

func (m *Mongo) Delete(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	result, err := m.collection.DeleteOne(ctx, filter)
	fmt.Printf("Delete result: %v, Delete count: %v\n", result, result.DeletedCount)
	return err
}
