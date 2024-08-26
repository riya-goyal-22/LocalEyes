package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/constants"
	"localEyes/internal/models"
	"localEyes/internal/repositories"
)

type AdminService struct {
	UserRepo repositories.UserRepository
	PostRepo repositories.PostRepository
	QuesRepo repositories.QuestionRepository
}

func NewAdminService(userRepo repositories.UserRepository, postRepo repositories.PostRepository, quesRepo repositories.QuestionRepository) *AdminService {
	return &AdminService{UserRepo: userRepo, PostRepo: postRepo, QuesRepo: quesRepo}
}

func (s *AdminService) Login(password string) (*models.Admin, error) {
	hashedPassword := hashPassword(password)
	user, err := s.UserRepo.FindAdminByUsernamePassword("admin", hashedPassword)
	if err != nil {
		return nil, errors.New(constants.Red + "Invalid username or password" + constants.Reset)
	}
	return user, nil
}

func (s *AdminService) GetAllUsers() ([]*models.User, error) {
	users, err := s.UserRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *AdminService) GetAllPosts() ([]*models.Post, error) {
	posts, err := s.PostRepo.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *AdminService) GetAllQuestions() ([]*models.Question, error) {
	questions, err := s.QuesRepo.GetAllQuestions()
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *AdminService) DeleteUser(UId primitive.ObjectID) error {
	err := s.UserRepo.DeleteByUId(UId)
	if err != nil {
		return err
	}
	err = s.PostRepo.DeleteByUId(UId)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) DeletePost(PId primitive.ObjectID) error {
	err := s.PostRepo.DeleteOneDoc(bson.M{"id": PId})
	if err != nil {
		return err
	}
	err = s.QuesRepo.DeleteByPId(PId)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) DeleteQuestion(QId primitive.ObjectID) error {
	err := s.QuesRepo.DeleteOneDoc(bson.M{"q_id": QId})
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) ReActivate(UId primitive.ObjectID) error {
	err := s.UserRepo.UpdateActiveStatus(UId, true)
	if err != nil {
		return err
	}
	return nil
}
