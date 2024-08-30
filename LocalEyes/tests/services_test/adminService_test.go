package services_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/internal/models"
	"localEyes/internal/services"
	mocks2 "localEyes/tests/mocks"
	"testing"
)

func TestAdminService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks2.NewMockUserRepository(ctrl)
	adminService := services.NewAdminService(mockUserRepo, nil, nil)

	username := "admin"
	password := "password"
	hashedPassword := services.HashPassword(password)
	expectedAdmin := &models.Admin{
		User: models.User{
			Username: username,
			Password: password,
		},
	}

	mockUserRepo.EXPECT().
		FindAdminByUsernamePassword(username, hashedPassword).
		Return(expectedAdmin, nil)

	admin, err := adminService.Login(password)

	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, admin)
}

func TestAdminService_Login_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks2.NewMockUserRepository(ctrl)
	adminService := services.NewAdminService(mockUserRepo, nil, nil)

	username := "admin"
	password := "wrongpassword"
	hashedPassword := services.HashPassword(password)

	mockUserRepo.EXPECT().
		FindAdminByUsernamePassword(username, hashedPassword).
		Return(nil, errors.New("invalid credentials"))

	admin, err := adminService.Login(password)

	assert.Error(t, err)
	assert.Nil(t, admin)
	assert.Contains(t, err.Error(), "Invalid username or password")
}

func TestAdminService_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks2.NewMockUserRepository(ctrl)
	adminService := services.NewAdminService(mockUserRepo, nil, nil)

	expectedUsers := []*models.User{
		{
			Username: "user1",
			Password: "password",
		},
	}

	mockUserRepo.EXPECT().
		GetAllUsers().
		Return(expectedUsers, nil)

	users, err := adminService.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestAdminService_GetAllUsers_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks2.NewMockUserRepository(ctrl)
	adminService := services.NewAdminService(mockUserRepo, nil, nil)

	mockUserRepo.EXPECT().
		GetAllUsers().
		Return(nil, errors.New("fetch error"))

	users, err := adminService.GetAllUsers()

	assert.Error(t, err)
	assert.Nil(t, users)
	assert.Contains(t, err.Error(), "fetch error")
}

func TestAdminService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks2.NewMockUserRepository(ctrl)
	mockPostRepo := mocks2.NewMockPostRepository(ctrl)
	adminService := services.NewAdminService(mockUserRepo, mockPostRepo, nil)

	userID := primitive.NewObjectID()

	// Set up expectations
	mockUserRepo.EXPECT().
		DeleteByUId(userID).
		Return(nil)
	mockPostRepo.EXPECT().
		DeleteByUId(userID).
		Return(nil)

	err := adminService.DeleteUser(userID)

	assert.NoError(t, err)
}

func TestAdminService_DeleteUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks2.NewMockUserRepository(ctrl)
	mockPostRepo := mocks2.NewMockPostRepository(ctrl)
	adminService := services.NewAdminService(mockUserRepo, mockPostRepo, nil)

	userID := primitive.NewObjectID()

	// Set up expectations with errors
	mockUserRepo.EXPECT().
		DeleteByUId(userID).
		Return(errors.New("user delete error"))
	mockPostRepo.EXPECT().
		DeleteByUId(userID).
		Return(nil)

	err := adminService.DeleteUser(userID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "user delete error")
}

func TestAdminService_ReActivate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks2.NewMockUserRepository(ctrl)
	adminService := services.NewAdminService(mockUserRepo, nil, nil)

	userID := primitive.NewObjectID()

	// Set up expectations
	mockUserRepo.EXPECT().
		UpdateActiveStatus(userID, true).
		Return(nil)

	err := adminService.ReActivate(userID)

	assert.NoError(t, err)
}

func TestAdminService_ReActivate_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks2.NewMockUserRepository(ctrl)
	adminService := services.NewAdminService(mockUserRepo, nil, nil)

	userID := primitive.NewObjectID()

	// Set up expectations with errors
	mockUserRepo.EXPECT().
		UpdateActiveStatus(userID, true).
		Return(errors.New("update error"))

	err := adminService.ReActivate(userID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "update error")
}

func TestAdminService_GetAllQuestions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks2.NewMockQuestionRepository(ctrl)
	adminService := services.NewAdminService(nil, nil, mockQuestionRepo)

	expectedQuestions := []*models.Question{
		{Text: "Text"},
	}

	mockQuestionRepo.EXPECT().
		GetAllQuestions().
		Return(expectedQuestions, nil)

	questions, err := adminService.GetAllQuestions()

	assert.NoError(t, err)
	assert.Equal(t, expectedQuestions, questions)
}

func TestAdminService_GetAllQuestions_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks2.NewMockQuestionRepository(ctrl)
	adminService := services.NewAdminService(nil, nil, mockQuestionRepo)

	mockQuestionRepo.EXPECT().
		GetAllQuestions().
		Return(nil, errors.New("fetch error"))

	questions, err := adminService.GetAllQuestions()

	assert.Error(t, err)
	assert.Nil(t, questions)
	assert.Contains(t, err.Error(), "fetch error")
}

func TestAdminService_DeleteQuestion(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks2.NewMockQuestionRepository(ctrl)
	adminService := services.NewAdminService(nil, nil, mockQuestionRepo)

	questionID := primitive.NewObjectID()

	// Set up expectations
	mockQuestionRepo.EXPECT().
		DeleteOneDoc(bson.M{"q_id": questionID}).
		Return(nil)

	err := adminService.DeleteQuestion(questionID)

	assert.NoError(t, err)
}

func TestAdminService_DeleteQuestion_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockQuestionRepo := mocks2.NewMockQuestionRepository(ctrl)
	adminService := services.NewAdminService(nil, nil, mockQuestionRepo)

	questionID := primitive.NewObjectID()

	// Set up expectations with errors
	mockQuestionRepo.EXPECT().
		DeleteOneDoc(bson.M{"q_id": questionID}).
		Return(errors.New("delete error"))

	err := adminService.DeleteQuestion(questionID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "delete error")
}

func TestAdminService_DeletePost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks2.NewMockPostRepository(ctrl)
	mockQuesRepo := mocks2.NewMockQuestionRepository(ctrl)
	adminService := services.NewAdminService(nil, mockPostRepo, mockQuesRepo)

	postID := primitive.NewObjectID()

	// Set up expectations
	mockPostRepo.EXPECT().
		DeleteOneDoc(bson.M{"id": postID}).
		Return(nil)
	mockQuesRepo.EXPECT().
		DeleteByPId(postID).
		Return(nil)

	err := adminService.DeletePost(postID)

	assert.NoError(t, err)
}

func TestAdminService_DeletePost_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks2.NewMockPostRepository(ctrl)
	mockQuesRepo := mocks2.NewMockQuestionRepository(ctrl)
	adminService := services.NewAdminService(nil, mockPostRepo, mockQuesRepo)

	postID := primitive.NewObjectID()

	// Set up expectations with errors
	mockPostRepo.EXPECT().
		DeleteOneDoc(bson.M{"id": postID}).
		Return(errors.New("delete error"))
	mockQuesRepo.EXPECT().
		DeleteByPId(postID).
		Return(nil)

	err := adminService.DeletePost(postID)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "delete error")
}

func TestAdminService_GetAllPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks2.NewMockPostRepository(ctrl)
	adminService := services.NewAdminService(nil, mockPostRepo, nil)

	expectedPosts := []*models.Post{ /* initialize with test data */ }

	mockPostRepo.EXPECT().
		GetAllPosts().
		Return(expectedPosts, nil)

	posts, err := adminService.GetAllPosts()

	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)
}

func TestAdminService_GetAllPosts_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks2.NewMockPostRepository(ctrl)
	adminService := services.NewAdminService(nil, mockPostRepo, nil)

	mockPostRepo.EXPECT().
		GetAllPosts().
		Return(nil, errors.New("fetch error"))

	posts, err := adminService.GetAllPosts()

	assert.Error(t, err)
	assert.Nil(t, posts)
	assert.Contains(t, err.Error(), "fetch error")
}
