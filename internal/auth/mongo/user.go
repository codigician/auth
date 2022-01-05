package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	Firstname      string             `bson:"firstname"`
	Lastname       string             `bson:"lastname"`
	Email          string             `bson:"email"`
	HashedPassword string             `bson:"hashed_password"`
}
