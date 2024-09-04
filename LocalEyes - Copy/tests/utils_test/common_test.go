package utils_test

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/utils"
	"os"
	"testing"
)

func TestPromptInput(t *testing.T) {
	expectedInput := "test input"
	// Mocking stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Simulate user input
	w.WriteString(expectedInput + "\n")
	w.Close()

	// Call the function
	input := utils.PromptInput("Enter something: ")

	// Verify the output
	assert.Equal(t, expectedInput, input)
}

func TestGetChoice(t *testing.T) {
	expectedChoice := "2"
	// Mocking stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	// Simulate user input
	w.WriteString(expectedChoice + "\n")
	w.Close()

	// Call the function
	choice := utils.GetChoice()

	// Verify the output
	assert.Equal(t, 2, choice)
}

func TestPromptID(t *testing.T) {
	// Ensure the test is being run
	t.Log("Running TestPromptID...")

	validID := primitive.NewObjectID().Hex()

	// Mocking stdin
	r, w, _ := os.Pipe()
	defer func() {
		_ = w.Close()
		_ = r.Close()
	}()
	os.Stdin = r

	// Simulate user input
	_, _ = w.WriteString(validID + "\n")

	// Call the function
	id, err := utils.PromptID("Enter ID: ")

	// Verify the output
	assert.NoError(t, err)
	assert.Equal(t, validID, id.Hex())

	// Ensure the test completed
	t.Log("TestPromptID completed.")
}

func TestPromptID_Invalid(t *testing.T) {
	invalidID := "invalidID"

	// Debug: Ensure this is running
	t.Log("Running TestPromptID_Invalid...")

	r, w, _ := os.Pipe()
	defer func() {
		_ = r.Close()
		_ = w.Close()
	}()
	os.Stdin = r

	_, _ = w.WriteString(invalidID + "\n")
	_ = w.Close()

	id, err := utils.PromptID("Enter ID: ")

	// Debug: Check the values of err and id
	t.Logf("err: %v, id: %v", err, id)

	// Verify the output
	assert.Error(t, err, "Expected an error for invalid ObjectID")
	assert.Equal(t, primitive.NilObjectID, id, "Expected a NilObjectID for invalid input")

}
