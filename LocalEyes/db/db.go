package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"localEyes/constants"
	"log"
	"sync"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
)

type CollectionInterface interface {
	InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
	DeleteMany(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error)
	UpdateFields(ctx context.Context, filter interface{}, updates interface{}) (*mongo.UpdateResult, error)
}

type MongoCollectionWrapper struct {
	collection *mongo.Collection
}

func NewCollectionWrapper(collection *mongo.Collection) CollectionInterface {
	return &MongoCollectionWrapper{
		collection: collection,
	}
}

func (w *MongoCollectionWrapper) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return w.collection.InsertOne(ctx, document)
}

func (w *MongoCollectionWrapper) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	return w.collection.FindOne(ctx, filter)
}

func (w *MongoCollectionWrapper) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return w.collection.Find(ctx, filter, opts...)
}

func (w *MongoCollectionWrapper) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return w.collection.DeleteOne(ctx, filter)
}

func (w *MongoCollectionWrapper) DeleteMany(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return w.collection.DeleteMany(ctx, filter)
}

func (w *MongoCollectionWrapper) UpdateFields(ctx context.Context, filter interface{}, updates interface{}) (*mongo.UpdateResult, error) {
	return w.collection.UpdateOne(ctx, filter, updates)
}

func GetMongoClient() *mongo.Client {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI(constants.URI)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}

		err = client.Ping(context.TODO(), nil)
		if err != nil {
			log.Fatal(err)
		}

		mongoClient = client
	})

	return mongoClient
}
func GetCollection(database, collection string) *mongo.Collection {
	return GetMongoClient().Database(database).Collection(collection)
}

func GetUsersCollection() *mongo.Collection {
	return GetCollection(constants.DatabaseName, constants.UsersCollection)
}

func GetPostsCollection() *mongo.Collection {
	return GetCollection(constants.DatabaseName, constants.PostsCollection)
}

func GetQuestionsCollection() *mongo.Collection {
	return GetCollection(constants.DatabaseName, constants.QuestionsCollection)
}
