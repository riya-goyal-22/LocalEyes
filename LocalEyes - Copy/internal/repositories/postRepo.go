package repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"localEyes/internal/db"
	"localEyes/internal/interfaces"
	"localEyes/internal/models"
)

type MongoPostRepository struct {
	Collection interfaces.CollectionInterface
}

func NewMongoPostRepository() interfaces.PostRepository {
	return &MongoPostRepository{
		Collection: db.NewCollectionWrapper(&db.MongoCollectionWrapper{Collection: db.GetPostsCollection()}),
	}
}

func (r *MongoPostRepository) Create(post *models.Post) error {
	_, err := r.Collection.InsertOne(context.Background(), post)
	return err
}

func (r *MongoPostRepository) GetAllPosts() ([]*models.Post, error) {
	var posts []*models.Post

	cursor, err := r.Collection.Find(context.Background(), bson.M{}, options.Find())
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
	_, err := r.Collection.DeleteOne(context.Background(), filter)
	return err
}

func (r *MongoPostRepository) DeleteByUId(UId primitive.ObjectID) error {
	_, err := r.Collection.DeleteMany(context.Background(), bson.M{"userId": UId})
	return err
}

func (r *MongoPostRepository) GetPostsByFilter(filter interface{}) ([]*models.Post, error) {
	cursor, err := r.Collection.Find(context.Background(), filter)
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
	_, err := r.Collection.UpdateFields(context.Background(), filter, update)
	return err
}

func (r *MongoPostRepository) UpdateLike(PId primitive.ObjectID) error {
	filter := bson.M{"id": PId}
	updates := bson.M{
		"likes": 1,
	}
	update := bson.M{"$inc": updates}
	_, err := r.Collection.UpdateFields(context.Background(), filter, update)
	return err
}
