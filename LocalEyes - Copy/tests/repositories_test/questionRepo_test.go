package repositories_test

import (
	"context"
	"localEyes/tests/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"localEyes/internal/models"
	"localEyes/internal/repositories"
)

func TestMongoQuestionRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	questionRepo := &repositories.MongoQuestionRepository{Collection: mockCollection}

	question := &models.Question{
		QId:       primitive.NewObjectID(),
		PostId:    primitive.NewObjectID(),
		UserId:    primitive.NewObjectID(),
		Text:      "What is this?",
		Replies:   []string{},
		CreatedAt: time.Now(),
	}

	mockCollection.EXPECT().
		InsertOne(context.Background(), question).
		Return(&mongo.InsertOneResult{}, nil)

	err := questionRepo.Create(question)
	assert.NoError(t, err)
}

func TestMongoQuestionRepository_DeleteOneDoc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	questionRepo := &repositories.MongoQuestionRepository{Collection: mockCollection}

	filter := bson.M{"q_id": primitive.NewObjectID()}

	mockCollection.EXPECT().
		DeleteOne(context.Background(), filter).
		Return(&mongo.DeleteResult{}, nil)

	err := questionRepo.DeleteOneDoc(filter)
	assert.NoError(t, err)
}

func TestMongoQuestionRepository_DeleteByPId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	questionRepo := &repositories.MongoQuestionRepository{Collection: mockCollection}

	postId := primitive.NewObjectID()

	mockCollection.EXPECT().
		DeleteMany(context.Background(), bson.M{"post_id": postId}).
		Return(&mongo.DeleteResult{}, nil)

	err := questionRepo.DeleteByPId(postId)
	assert.NoError(t, err)
}

func TestMongoQuestionRepository_UpdateQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	questionRepo := &repositories.MongoQuestionRepository{Collection: mockCollection}

	questionId := primitive.NewObjectID()
	answer := "This is the answer"

	mockCollection.EXPECT().
		UpdateFields(context.Background(), bson.M{"q_id": questionId}, bson.M{"$push": bson.M{"replies": answer}}).
		Return(&mongo.UpdateResult{}, nil)

	err := questionRepo.UpdateQuestion(questionId, answer)
	assert.NoError(t, err)
}
