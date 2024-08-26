package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	PostId    primitive.ObjectID `bson:"id"`
	UId       primitive.ObjectID `bson:"userId"`
	Title     string             `bson:"title"`
	Type      string             `bson:"type"`
	Content   string             `bson:"content"`
	Likes     int                `bson:"likes"`
	CreatedAt time.Time          `bson:"created_at"`
}
