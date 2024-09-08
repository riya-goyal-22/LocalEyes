package utils

import (
	"bufio"
	"errors"
	"fmt"
	"localEyes/constants"
	"localEyes/internal/interfaces"
	"os"
	"strconv"
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
	return input
}

func GetChoice() int {
	fmt.Print("Enter choice: ")
	var choice int
	fmt.Scanln(&choice)
	fmt.Println()
	return choice
}

func PromptIntInput(prompt string) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(constants.Cyan + prompt + constants.Reset)
	scanner.Scan()
	input := scanner.Text()
	if err := scanner.Err(); err != nil {
		return 0, errors.New("Error reading input")
	}
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return num, nil
}

//func PromptPassword(prompt string) string {
//	prompt1 := promptui.Prompt{
//		Label:     prompt,
//		Mask:      '*',
//		IsConfirm: false,
//	}
//	result, err := prompt1.Run()
//	if err != nil {
//		fmt.Println("Prompt failed:", err)
//		return ""
//	}
//	return result
//}

func PromptPassword(promptInstance interfaces.PromptInterface) string {
	result, err := promptInstance.Run()
	if err != nil {
		return ""
	}
	return result
}
