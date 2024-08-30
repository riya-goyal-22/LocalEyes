package services_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/constants"
	"localEyes/internal/models"
	"localEyes/internal/services"
	"localEyes/tests/mocks"
	"testing"
)

//func TestUserService_Signup(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockUserRepo := mocks.NewMockUserRepository(ctrl)
//	userService := services.NewUserService(mockUserRepo)
//
//	username := "testuser"
//	password := "password123"
//	dwellingAge := 5
//	tag := "testtag"
//
//	hashedPassword := services.HashPassword(password)
//	user := &models.User{
//		UId:          primitive.NewObjectID(),
//		Username:     username,
//		Password:     hashedPassword,
//		City:         "delhi",
//		Notification: []string{},
//		IsActive:     true,
//		DwellingAge:  dwellingAge,
//		Tag:          tag,
//	}
//
//	mockUserRepo.EXPECT().
//		Create(user).
//		Return(nil)
//
//	err := userService.Signup(username, password, dwellingAge, tag)
//
//	assert.NoError(t, err)
//}

func TestUserService_Login_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockUserRepo)

	username := "testuser"
	password := "password123"
	hashedPassword := services.HashPassword(password)

	user := &models.User{
		UId:          primitive.NewObjectID(),
		Username:     username,
		Password:     hashedPassword,
		City:         "delhi",
		Notification: []string{},
		IsActive:     true,
	}

	mockUserRepo.EXPECT().
		FindByUsernamePassword(username, hashedPassword).
		Return(user, nil)

	user.NotifyChannel = make(chan string, 5) // Ensure NotifyChannel is set

	loggedInUser, err := userService.Login(username, password)

	assert.NoError(t, err)
	assert.Equal(t, user, loggedInUser)
	assert.NotNil(t, loggedInUser.NotifyChannel)
}

func TestUserService_Login_InvalidCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockUserRepo)

	username := "testuser"
	password := "wrongpassword"
	hashedPassword := services.HashPassword(password)

	mockUserRepo.EXPECT().
		FindByUsernamePassword(username, hashedPassword).
		Return(nil, errors.New("invalid credentials"))

	user, err := userService.Login(username, password)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Contains(t, err.Error(), "Invalid ")
}

func TestUserService_Login_InactiveAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockUserRepo)

	username := "testuser"
	password := "password123"
	hashedPassword := services.HashPassword(password)

	user := &models.User{
		UId:          primitive.NewObjectID(),
		Username:     username,
		Password:     hashedPassword,
		City:         "delhi",
		Notification: []string{},
		IsActive:     false,
	}

	mockUserRepo.EXPECT().
		FindByUsernamePassword(username, hashedPassword).
		Return(user, errors.New(constants.Red+"Account is Inactive"+constants.Reset))

	result, err := userService.Login(username, password)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "Account")
}

func TestUserService_DeActivate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.NewUserService(mockUserRepo)

	userID := primitive.NewObjectID()

	mockUserRepo.EXPECT().
		UpdateActiveStatus(userID, false).
		Return(nil)

	err := userService.DeActivate(userID)

	assert.NoError(t, err)
}

func TestNotifyUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.UserService{Repo: mockRepo}

	// Test data
	userID := primitive.NewObjectID()
	title := "Test Title"

	// Expected behavior
	mockRepo.EXPECT().PushNotification(userID, title).Return(nil)

	// Call the method
	err := userService.NotifyUsers(userID, title)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}

func TestUnNotifyUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	userService := services.UserService{Repo: mockRepo}

	// Test data
	userID := primitive.NewObjectID()

	// Expected behavior
	mockRepo.EXPECT().ClearNotification(userID).Return(nil)

	// Call the method
	err := userService.UnNotifyUsers(userID)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}
