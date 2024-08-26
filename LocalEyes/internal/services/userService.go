package services

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/constants"
	"localEyes/internal/models"
	"localEyes/internal/repositories"
)

type UserService struct {
	Repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Signup(username, password string, dwellingAge int, tag string) error {
	hashedPassword := hashPassword(password)

	user := &models.User{
		UId:          primitive.NewObjectID(),
		Username:     username,
		Password:     hashedPassword,
		City:         "delhi",
		Notification: []string{},
		IsActive:     true,
		DwellingAge:  dwellingAge,
		Tag:          tag,
		//NotifyChannel: make(chan string, 5),
		//IsAdmin:       false,
	}

	return s.Repo.Create(user)
}

func (s *UserService) Login(Username, password string) (*models.User, error) {
	hashedPassword := hashPassword(password)
	user, err := s.Repo.FindByUsernamePassword(Username, hashedPassword)
	user.NotifyChannel = make(chan string, 5)
	if err != nil {
		return nil, errors.New(constants.Red + "Invalid username or password" + constants.Reset)
	}

	if !user.IsActive {
		return nil, errors.New(constants.Red + "Account is Inactive" + constants.Reset)
	}
	return user, nil
}

func (s *UserService) DeActivate(UId primitive.ObjectID) error {
	err := s.Repo.UpdateActiveStatus(UId, false)
	if err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
