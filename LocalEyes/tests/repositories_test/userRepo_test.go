package repositories_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"localEyes/internal/models"
	"localEyes/internal/repositories"
	"localEyes/mocks" // Adjust import path if needed
)

func TestMongoUserRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	userRepo := &repositories.MongoUserRepository{Collection: mockCollection}

	user := &models.User{
		UId:      primitive.NewObjectID(),
		Username: "testuser",
		Password: "hashedpassword",
	}

	mockCollection.EXPECT().
		InsertOne(context.Background(), user).
		Return(&mongo.InsertOneResult{}, nil)

	err := userRepo.Create(user)
	assert.NoError(t, err)

}

func TestMongoUserRepository_DeleteByUId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	userRepo := &repositories.MongoUserRepository{Collection: mockCollection}

	userId := primitive.NewObjectID()

	mockCollection.EXPECT().
		DeleteOne(context.Background(), bson.M{"id": userId}).
		Return(&mongo.DeleteResult{}, nil)

	err := userRepo.DeleteByUId(userId)
	assert.NoError(t, err)
}

func TestMongoUserRepository_UpdateActiveStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	userRepo := &repositories.MongoUserRepository{Collection: mockCollection}

	userId := primitive.NewObjectID()
	status := false

	mockCollection.EXPECT().
		UpdateFields(context.Background(), bson.M{"id": userId}, bson.M{"$set": bson.M{"is_active": status}}).
		Return(&mongo.UpdateResult{}, nil)

	err := userRepo.UpdateActiveStatus(userId, status)
	assert.NoError(t, err)
}
