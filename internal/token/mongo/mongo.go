package mongo

import (
	"context"
	"fmt"

	"github.com/codigician/auth/internal/token"
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

func (m *Mongo) Save(ctx context.Context, rt *token.RefreshToken) error {
	objectID, err := primitive.ObjectIDFromHex(rt.ID)
	if err != nil {
		return err
	}
	result, err := m.collection.InsertOne(ctx, RefreshToken{
		ID:             objectID,
		Token:          rt.Token,
		ExpirationDate: rt.ExpirationDate,
	})
	fmt.Println("inserted a single document:", result)
	return err
}

func (m *Mongo) Get(ctx context.Context, id string) (*token.RefreshToken, error) {
	var rt RefreshToken
	filter := bson.M{"_id": id}
	err := m.collection.FindOne(ctx, filter).Decode(&rt)
	return &token.RefreshToken{
		ID:             rt.ID.Hex(),
		Token:          rt.Token,
		ExpirationDate: rt.ExpirationDate,
	}, err
}
