package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/constants"
	"localEyes/internal/interfaces"
	"localEyes/internal/models"
)

type AdminService struct {
	UserRepo interfaces.UserRepository
	PostRepo interfaces.PostRepository
	QuesRepo interfaces.QuestionRepository
}

func NewAdminService(userRepo interfaces.UserRepository, postRepo interfaces.PostRepository, quesRepo interfaces.QuestionRepository) *AdminService {
	return &AdminService{UserRepo: userRepo, PostRepo: postRepo, QuesRepo: quesRepo}
}

func (s *AdminService) Login(password string) (*models.Admin, error) {
	hashedPassword := HashPassword(password)
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
	err1 := s.UserRepo.DeleteByUId(UId)
	//if err != nil {
	//	return err
	//}
	err2 := s.PostRepo.DeleteByUId(UId)
	if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
	}
	return nil
}

func (s *AdminService) DeletePost(PId primitive.ObjectID) error {
	err1 := s.PostRepo.DeleteOneDoc(bson.M{"id": PId})
	err2 := s.QuesRepo.DeleteByPId(PId)
	if err1 != nil {
		return err1
	} else if err2 != nil {
		return err2
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
