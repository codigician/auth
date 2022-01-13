package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type RefreshToken struct {
	ID             primitive.ObjectID `bson:"_id"`
	Token          string             `bson:"token"`
	ExpirationDate int64              `bson:"expiration_date"`
}
