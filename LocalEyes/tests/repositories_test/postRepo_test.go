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

func TestMongoPostRepository_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	postRepo := &repositories.MongoPostRepository{Collection: mockCollection}

	post := &models.Post{
		UId:       primitive.NewObjectID(),
		PostId:    primitive.NewObjectID(),
		Title:     "Test Post",
		Content:   "This is a test post.",
		Type:      "general",
		CreatedAt: time.Now(),
		Likes:     0,
	}

	mockCollection.EXPECT().
		InsertOne(context.Background(), post).
		Return(&mongo.InsertOneResult{}, nil)

	err := postRepo.Create(post)
	assert.NoError(t, err)
}

func TestMongoPostRepository_DeleteOneDoc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	postRepo := &repositories.MongoPostRepository{Collection: mockCollection}

	filter := bson.M{"postId": primitive.NewObjectID()}

	mockCollection.EXPECT().
		DeleteOne(context.Background(), filter).
		Return(&mongo.DeleteResult{}, nil)

	err := postRepo.DeleteOneDoc(filter)
	assert.NoError(t, err)
}

func TestMongoPostRepository_DeleteByUId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	postRepo := &repositories.MongoPostRepository{Collection: mockCollection}

	userID := primitive.NewObjectID()

	mockCollection.EXPECT().
		DeleteMany(context.Background(), bson.M{"userId": userID}).
		Return(&mongo.DeleteResult{}, nil)

	err := postRepo.DeleteByUId(userID)
	assert.NoError(t, err)
}

func TestMongoPostRepository_UpdateUserPost(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	postRepo := &repositories.MongoPostRepository{Collection: mockCollection}

	postID := primitive.NewObjectID()
	userID := primitive.NewObjectID()
	title := "Updated Title"
	content := "Updated Content"

	filter := bson.M{"id": postID, "userId": userID}
	update := bson.M{"$set": bson.M{"title": title, "content": content}}

	mockCollection.EXPECT().
		UpdateFields(context.Background(), filter, update).
		Return(&mongo.UpdateResult{}, nil)

	err := postRepo.UpdateUserPost(postID, userID, title, content)
	assert.NoError(t, err)
}

func TestMongoPostRepository_UpdateLike(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	postRepo := &repositories.MongoPostRepository{Collection: mockCollection}

	postID := primitive.NewObjectID()
	filter := bson.M{"id": postID}
	update := bson.M{"$inc": bson.M{"likes": 1}}

	mockCollection.EXPECT().
		UpdateFields(context.Background(), filter, update).
		Return(&mongo.UpdateResult{}, nil)

	err := postRepo.UpdateLike(postID)
	assert.NoError(t, err)
}

//// MockCursor is a simple mock implementation of mongo.Cursor
//type MockCursor struct {
//	Documents []models.Post
//	Index     int
//}
//
//func (m *MockCursor) Next(ctx context.Context) bool {
//	if m.Index < len(m.Documents) {
//		m.Index++
//		return true
//	}
//	return false
//}
//
//func (m *MockCursor) Decode(v interface{}) error {
//	if m.Index == 0 || m.Index > len(m.Documents) {
//		return mongo.ErrNoDocuments
//	}
//	*v.(*models.Post) = m.Documents[m.Index-1]
//	return nil
//}
//
//func (m *MockCursor) Err() error {
//	return nil
//}
//
//func (m *MockCursor) Close(ctx context.Context) error {
//	return nil
//}
//
//func TestGetAllPosts(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	// Create mock collection
//	mockCollection := mocks.NewMockCollectionInterface(ctrl)
//
//	// Create an instance of MongoPostRepository with the mock collection
//	repo := &repositories.MongoPostRepository{
//		Collection: mockCollection,
//	}
//
//	// Define test data
//	posts := []*models.Post{
//		{Title: "Post 1", Content: "Content 1"},
//		{Title: "Post 2", Content: "Content 2"},
//	}
//
//	// Create a mock cursor with predefined documents
//	mockCursor := &MockCursor{Documents: []models.Post{*posts[0], *posts[1]}}
//
//	// Set up mock expectations
//	mockCollection.EXPECT().Find(gomock.Any(), bson.M{}, gomock.Any()).Return(mockCursor, nil)
//
//	// Call the method
//	result, err := repo.GetAllPosts()
//
//	// Assertions
//	assert.NoError(t, err)
//	assert.Equal(t, len(posts), len(result))
//	assert.Equal(t, posts[0].Title, result[0].Title)
//	assert.Equal(t, posts[1].Title, result[1].Title)
//}
