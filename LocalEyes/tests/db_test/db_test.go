// build+ !test
package db_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"localEyes/internal/db"
	"localEyes/tests/mocks"
	"testing"
)

func TestInsertOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	db := db.NewCollectionWrapper(mockCollection)

	document := map[string]interface{}{"key": "value"}
	expectedResult := &mongo.InsertOneResult{InsertedID: "12345"}
	mockCollection.EXPECT().InsertOne(gomock.Any(), document).Return(expectedResult, nil)

	result, err := db.InsertOne(context.Background(), document)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestFindOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	db := db.NewCollectionWrapper(mockCollection)

	filter := map[string]interface{}{"key": "value"}
	expectedResult := &mongo.SingleResult{}
	mockCollection.EXPECT().FindOne(gomock.Any(), filter).Return(expectedResult)

	result := db.FindOne(context.Background(), filter)
	assert.Equal(t, expectedResult, result)
}

func TestDeleteOne(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	db := db.NewCollectionWrapper(mockCollection)

	filter := map[string]interface{}{"key": "value"}
	expectedResult := &mongo.DeleteResult{DeletedCount: 1}
	mockCollection.EXPECT().DeleteOne(gomock.Any(), filter).Return(expectedResult, nil)

	result, err := db.DeleteOne(context.Background(), filter)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestUpdateFields(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	db := db.NewCollectionWrapper(mockCollection)

	filter := map[string]interface{}{"key": "value"}
	updates := map[string]interface{}{"$set": map[string]interface{}{"field": "newValue"}}
	expectedResult := &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}
	mockCollection.EXPECT().UpdateFields(gomock.Any(), filter, updates).Return(expectedResult, nil)

	result, err := db.UpdateFields(context.Background(), filter, updates)
	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestFind(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCollection := mocks.NewMockCollectionInterface(ctrl)
	db1 := db.NewCollectionWrapper(mockCollection)

	filter := map[string]interface{}{"key": "value"}
	opts := &options.FindOptions{}
	expectedCursor := &mongo.Cursor{}
	mockCollection.EXPECT().Find(gomock.Any(), filter, opts).Return(expectedCursor, nil)

	cursor, err := db1.Find(context.Background(), filter, opts)
	assert.NoError(t, err)
	assert.Equal(t, expectedCursor, cursor)
}
