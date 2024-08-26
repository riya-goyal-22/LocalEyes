package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UId           primitive.ObjectID `bson:"id"`
	Username      string             `bson:"username"`
	Password      string             `bson:"password"`
	City          string             `bson:"city"`
	DwellingAge   int                `bson:"dwelling_age"`
	IsActive      bool               `bson:"is_active"`
	Notification  []string           `bson:"notification"`
	Tag           string             `bson:"tag"`
	NotifyChannel chan string        `bson:"-"` //ignore
	//IsAdmin       bool        `bson:"is_admin"`
}
