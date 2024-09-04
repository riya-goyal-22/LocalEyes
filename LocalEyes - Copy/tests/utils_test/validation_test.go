package utils_test

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"localEyes/internal/models"
	"localEyes/tests/mocks"
	"localEyes/utils"
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		expected bool
	}{
		{"password1@", true},    // Valid password
		{"short1@", true},       // valid
		{"longpassword", false}, // No special character
		{"Password@", false},    // No number
	}

	for _, test := range tests {
		result := utils.ValidatePassword(test.password)
		assert.Equal(t, test.expected, result)
	}
}

func TestValidateFilter(t *testing.T) {
	tests := []struct {
		filter   string
		expected bool
	}{
		{"food", true},
		{"travel", true},
		{"shopping", true},
		{"other", true},
		{"invalid", false},
	}

	for _, test := range tests {
		result := utils.ValidateFilter(test.filter)
		assert.Equal(t, test.expected, result)
	}
}

func TestValidateUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)

	// Case 1: Username not found
	mockRepo.EXPECT().FindByUsername("nonexistent_user").Return(nil, mongo.ErrNoDocuments)

	isValid := utils.ValidateUsername("nonexistent_user", mockRepo)
	assert.True(t, isValid, "Expected username to be valid when not found")

	// Case 2: Username found
	mockRepo.EXPECT().FindByUsername("existing_user").Return(&models.User{}, nil)

	isValid = utils.ValidateUsername("existing_user", mockRepo)
	assert.False(t, isValid, "Expected username to be invalid when found")
}
