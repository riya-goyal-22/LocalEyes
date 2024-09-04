package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Question struct {
	QId       primitive.ObjectID `bson:"q_id"`
	PostId    primitive.ObjectID `bson:"post_id"`
	UserId    primitive.ObjectID `bson:"user_id"`
	Text      string             `bson:"text"`
	Replies   []string           `bson:"replies"`
	CreatedAt time.Time          `bson:"created_at"`
}
