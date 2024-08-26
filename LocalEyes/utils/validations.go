package utils

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"localEyes/constants"
	"localEyes/internal/repositories"
	"localEyes/internal/services"
	"strings"
)

func ValidateUsername(username string, userRepo repositories.UserRepository) bool {
	_, err := userRepo.FindByUsername(username)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return true
	}
	return false
}

func ValidatePassword(password string) bool {
	if len(password) > 6 {
		if strings.Contains(password, "@") || strings.Contains(password, "#") || strings.Contains(password, "$") || strings.Contains(password, "%") || strings.Contains(password, "^") || strings.Contains(password, "*") {
			if strings.Contains(password, "1") || strings.Contains(password, "2") || strings.Contains(password, "3") || strings.Contains(password, "4") || strings.Contains(password, "5") || strings.Contains(password, "6") || strings.Contains(password, "7") || strings.Contains(password, "8") || strings.Contains(password, "9") || strings.Contains(password, "0") {
				return true
			}
		}
	} else {
		fmt.Println(constants.Red + "Password is not strong" + constants.Reset)
	}
	return false
}

func ValidateFilter(filter string) bool {
	return filter == "food" || filter == "travel" || filter == "shopping" || filter == "other"
}

func CheckPasswordHash(username string, password string, userService *services.UserService) bool {
	_, err := userService.Repo.FindByUsernamePassword(username, password)
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false
	}
	return true
}
