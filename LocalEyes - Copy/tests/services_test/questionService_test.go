package services_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/internal/models"
	"localEyes/internal/services"
	"localEyes/tests/mocks"
	"testing"
)

//func TestQuestionService_AskQuestion(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockQuestionRepo := mocks.NewMockQuestionRepository(ctrl)
//	questionService := services.NewQuestionService(mockQuestionRepo)
//
//	userID := primitive.NewObjectID()
//	postID := primitive.NewObjectID()
//	content := "Question content"
//	question := &models.Question{}
//
//	mockQuestionRepo.EXPECT().
//		Create(question).
//		Return(nil)
//
//	err := questionService.AskQuestion(userID, postID, content)
//
//	assert.NoError(t, err)
//}

func TestQuestionService_DeleteQuesByPId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks.NewMockQuestionRepository(ctrl)
	questionService := services.NewQuestionService(mockQuestionRepo)

	postID := primitive.NewObjectID()

	mockQuestionRepo.EXPECT().
		DeleteByPId(postID).
		Return(nil)

	err := questionService.DeleteQuesByPId(postID)

	assert.NoError(t, err)
}

func TestQuestionService_DeleteUserQues(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks.NewMockQuestionRepository(ctrl)
	questionService := services.NewQuestionService(mockQuestionRepo)

	userID := primitive.NewObjectID()
	questionID := primitive.NewObjectID()

	filter := bson.M{"q_id": questionID, "user_id": userID}
	mockQuestionRepo.EXPECT().
		DeleteOneDoc(filter).
		Return(nil)

	err := questionService.DeleteUserQues(userID, questionID)

	assert.NoError(t, err)
}

func TestQuestionService_GetPostQuestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks.NewMockQuestionRepository(ctrl)
	questionService := services.NewQuestionService(mockQuestionRepo)

	postID := primitive.NewObjectID()
	expectedQuestions := []*models.Question{
		{QId: primitive.NewObjectID(), PostId: postID, Text: "Question 1"},
		{QId: primitive.NewObjectID(), PostId: postID, Text: "Question 2"},
	}

	mockQuestionRepo.EXPECT().
		GetQuestionsByPId(postID).
		Return(expectedQuestions, nil)

	questions, err := questionService.GetPostQuestions(postID)

	assert.NoError(t, err)
	assert.Equal(t, expectedQuestions, questions)
}

func TestQuestionService_GetPostQuestions_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks.NewMockQuestionRepository(ctrl)
	questionService := services.NewQuestionService(mockQuestionRepo)

	postID := primitive.NewObjectID()

	mockQuestionRepo.EXPECT().
		GetQuestionsByPId(postID).
		Return(nil, errors.New("fetch error"))

	questions, err := questionService.GetPostQuestions(postID)

	assert.Error(t, err)
	assert.Nil(t, questions)
	assert.Contains(t, err.Error(), "fetch error")
}

func TestQuestionService_AddAnswer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks.NewMockQuestionRepository(ctrl)
	questionService := services.NewQuestionService(mockQuestionRepo)

	questionID := primitive.NewObjectID()
	answer := "This is an answer"

	mockQuestionRepo.EXPECT().
		UpdateQuestion(questionID, answer).
		Return(nil)

	err := questionService.AddAnswer(questionID, answer)

	assert.NoError(t, err)
}
