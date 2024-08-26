package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"localEyes/db"
	"localEyes/internal/models"
)

type MongoPostRepository struct {
	collection db.CollectionInterface
}

func NewMongoPostRepository() PostRepository {
	return &MongoPostRepository{
		collection: db.NewCollectionWrapper(db.GetPostsCollection()),
	}
}

func (r *MongoPostRepository) Create(post *models.Post) error {
	_, err := r.collection.InsertOne(context.Background(), post)
	return err
}

func (r *MongoPostRepository) GetAllPosts() ([]*models.Post, error) {
	var posts []*models.Post

	cursor, err := r.collection.Find(context.Background(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(cursor, context.Background())

	for cursor.Next(context.Background()) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *MongoPostRepository) DeleteOneDoc(filter interface{}) error {
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *MongoPostRepository) DeleteByUId(UId primitive.ObjectID) error {
	_, err := r.collection.DeleteMany(context.Background(), bson.M{"userId": UId})
	return err
}

func (r *MongoPostRepository) GetPostsByFilter(filter interface{}) ([]*models.Post, error) {
	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			fmt.Println(err)
		}
	}(cursor, context.Background())

	var posts []*models.Post
	for cursor.Next(context.Background()) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *MongoPostRepository) UpdateUserPost(PId, UId primitive.ObjectID, title string, content string) error {
	filter := bson.M{"id": PId, "userId": UId}
	updates := bson.M{
		"title":   title,
		"content": content,
	}
	update := bson.M{"$set": updates}
	_, err := r.collection.UpdateFields(context.Background(), filter, update)
	return err
}

func (r *MongoPostRepository) UpdateLike(PId primitive.ObjectID) error {
	filter := bson.M{"id": PId}
	updates := bson.M{
		"likes": 1,
	}
	update := bson.M{"$inc": updates}
	_, err := r.collection.UpdateFields(context.Background(), filter, update)
	return err
}
