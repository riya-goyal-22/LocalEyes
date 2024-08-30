package interfaces

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
	UpdateFields(ctx context.Context, filter interface{}, updates interface{}) (*mongo.UpdateResult, error)
}
