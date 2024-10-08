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

type MongoUserRepository struct {
	Collection interfaces.CollectionInterface
}

func NewMongoUserRepository() interfaces.UserRepository {
	return &MongoUserRepository{
		Collection: db.NewCollectionWrapper(&db.MongoCollectionWrapper{Collection: db.GetUsersCollection()}),
	}
}

func (r *MongoUserRepository) Create(user *models.User) error {
	_, err := r.Collection.InsertOne(context.Background(), user)
	return err
}

func (r *MongoUserRepository) FindByUId(UId primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(context.Background(), bson.M{"id": UId}).Decode(&user)
	return &user, err
}

func (r *MongoUserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return &user, err
}

func (r *MongoUserRepository) FindByUsernamePassword(username, password string) (*models.User, error) {
	var user models.User
	err := r.Collection.FindOne(context.Background(), bson.M{"username": username, "password": password}).Decode(&user)
	return &user, err
}

func (r *MongoUserRepository) FindAdminByUsernamePassword(username, password string) (*models.Admin, error) {
	var user models.Admin
	err := r.Collection.FindOne(context.Background(), bson.M{"username": username, "password": password}).Decode(&user)
	return &user, err
}

func (r *MongoUserRepository) GetAllUsers() ([]*models.User, error) {
	var users []*models.User

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
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *MongoUserRepository) DeleteByUId(UId primitive.ObjectID) error {
	_, err := r.Collection.DeleteOne(context.Background(), bson.M{"id": UId})
	return err
}

func (r *MongoUserRepository) UpdateActiveStatus(UId primitive.ObjectID, status bool) error {
	filter := bson.M{"id": UId}
	updates := bson.M{
		"is_active": status,
	}
	update := bson.M{"$set": updates}
	_, err := r.Collection.UpdateFields(context.Background(), filter, update)
	return err
}

func (r *MongoUserRepository) PushNotification(UId primitive.ObjectID, title string) error {
	filter := bson.M{"id": bson.M{"$ne": UId}}
	updates := bson.M{
		"notification": "New post :" + title,
	}
	update := bson.M{"$push": updates}
	_, err := r.Collection.UpdateFields(context.Background(), filter, update)
	return err
}

func (r *MongoUserRepository) ClearNotification(UId primitive.ObjectID) error {
	filter := bson.M{"id": UId}
	updates := bson.M{
		"notification": []string{},
	}
	update := bson.M{"$set": updates}
	_, err := r.Collection.UpdateFields(context.Background(), filter, update)
	return err
}
