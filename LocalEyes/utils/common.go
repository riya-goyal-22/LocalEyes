package utils

import (
	"bufio"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/constants"
	"os"
)

func PromptInput(prompt string) string {

	// Create a new Scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Display the prompt message
	fmt.Print(constants.Cyan + prompt + constants.Reset)

	// Read the next line of input
	scanner.Scan()

	// Get the text from the scanner
	input := scanner.Text()

	// Check for any errors during scanning
	//if err := scanner.Err(); err != nil {
	//	fmt.Println("Error reading input:", err)
	//	return ""
	//}
	return input
}

func GetChoice() int {
	fmt.Print("Enter choice: ")
	var choice int
	fmt.Scanln(&choice)
	fmt.Println()
	return choice
}

func PromptID(prompt string) (primitive.ObjectID, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(constants.Cyan + prompt + constants.Reset)
	scanner.Scan()
	input := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
	uid, err := primitive.ObjectIDFromHex(input)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return uid, err
}
