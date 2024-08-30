package services_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/internal/models"
	"localEyes/internal/services"
	"localEyes/mocks"
	"testing"
)

func TestPostService_UpdateMyPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	postID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	title := "Updated Title"
	content := "Updated Content"

	mockPostRepo.EXPECT().
		UpdateUserPost(postID, userID, title, content).
		Return(nil)

	err := postService.UpdateMyPost(postID, userID, title, content)

	assert.NoError(t, err)
}

func TestPostService_GiveAllPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	expectedPosts := []*models.Post{
		{Title: "Title1", Content: "Content1"},
		{Title: "Title2", Content: "Content2"},
	}

	mockPostRepo.EXPECT().
		GetAllPosts().
		Return(expectedPosts, nil)

	posts, err := postService.GiveAllPosts()

	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)
}

func TestPostService_GiveAllPosts_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	mockPostRepo.EXPECT().
		GetAllPosts().
		Return(nil, errors.New("fetch error"))

	posts, err := postService.GiveAllPosts()

	assert.Error(t, err)
	assert.Nil(t, posts)
	assert.Contains(t, err.Error(), "fetch error")
}

func TestPostService_GiveMyPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	userID := primitive.NewObjectID()
	expectedPosts := []*models.Post{
		{Title: "Title1", Content: "Content1"},
	}

	filter := bson.M{"userId": userID}
	mockPostRepo.EXPECT().
		GetPostsByFilter(filter).
		Return(expectedPosts, nil)

	posts, err := postService.GiveMyPosts(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)
}

func TestPostService_DeleteMyPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	userID := primitive.NewObjectID()
	postID := primitive.NewObjectID()

	filter := bson.M{"userId": userID, "id": postID}
	mockPostRepo.EXPECT().
		DeleteOneDoc(filter).
		Return(nil)

	err := postService.DeleteMyPost(userID, postID)

	assert.NoError(t, err)
}

func TestPostService_Like(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	postID := primitive.NewObjectID()

	mockPostRepo.EXPECT().
		UpdateLike(postID).
		Return(nil)

	err := postService.Like(postID)

	assert.NoError(t, err)
}

func TestPostService_GiveFilteredPosts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	filterType := "General"
	expectedPosts := []*models.Post{
		{Title: "Title1", Content: "Content1"},
	}

	filter := bson.M{"type": filterType}
	mockPostRepo.EXPECT().
		GetPostsByFilter(filter).
		Return(expectedPosts, nil)

	posts, err := postService.GiveFilteredPosts(filterType)

	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)
}

func TestPostService_PostIdExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	postID := primitive.NewObjectID()

	filter := bson.M{"id": postID}
	mockPostRepo.EXPECT().
		GetPostsByFilter(filter).
		Return([]*models.Post{{}}, nil)

	exists, err := postService.PostIdExist(postID)

	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestPostService_PostIdExist_NotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPostRepo := mocks.NewMockPostRepository(ctrl)
	postService := services.NewPostService(mockPostRepo)

	postID := primitive.NewObjectID()

	filter := bson.M{"id": postID}
	mockPostRepo.EXPECT().
		GetPostsByFilter(filter).
		Return(nil, nil)

	exists, err := postService.PostIdExist(postID)

	assert.NoError(t, err)
	assert.False(t, exists)
}
