package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	Config struct {
		URI string
	}

	Mongo struct {
		client *mongo.Client
	}
)

func New(conf *Config) (*Mongo, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(conf.URI))
	return &Mongo{client}, err
}

func (m *Mongo) Collection(database, collection string) *mongo.Collection {
	return m.client.Database(database).Collection(collection)
}

func (m *Mongo) Disconnect(ctx context.Context) error {
	return m.client.Disconnect(ctx)
}
