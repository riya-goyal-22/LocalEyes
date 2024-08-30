package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"localEyes/constants"
	"localEyes/internal/interfaces"
	"log"
	"sync"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
)

type MongoCollectionWrapper struct {
	Collection *mongo.Collection
}

func NewCollectionWrapper(collectionInterface interfaces.CollectionInterface) interfaces.CollectionInterface {
	return collectionInterface
}

func (w *MongoCollectionWrapper) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	return w.Collection.InsertOne(ctx, document)
}

func (w *MongoCollectionWrapper) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	return w.Collection.FindOne(ctx, filter)
}

func (w *MongoCollectionWrapper) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return w.Collection.Find(ctx, filter, opts...)
}

func (w *MongoCollectionWrapper) DeleteOne(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return w.Collection.DeleteOne(ctx, filter)
}

func (w *MongoCollectionWrapper) DeleteMany(ctx context.Context, filter interface{}) (*mongo.DeleteResult, error) {
	return w.Collection.DeleteMany(ctx, filter)
}

func (w *MongoCollectionWrapper) UpdateFields(ctx context.Context, filter interface{}, updates interface{}) (*mongo.UpdateResult, error) {
	return w.Collection.UpdateMany(ctx, filter, updates)
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
