package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/internal/models"
	"localEyes/internal/repositories"
	"time"
)

type QuestionService struct {
	repo repositories.QuestionRepository
}

func NewQuestionService(repo repositories.QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) AskQuestion(userId, postId primitive.ObjectID, content string) error {
	question := &models.Question{
		QId:       primitive.NewObjectID(),
		PostId:    postId,
		UserId:    userId,
		Text:      content,
		Replies:   make([]string, 0),
		CreatedAt: time.Now(),
	}
	return s.repo.Create(question)
}

func (s *QuestionService) DeleteQuesByPId(postId primitive.ObjectID) error {
	err := s.repo.DeleteByPId(postId)
	if err != nil {
		return err
	}
	return nil
}

func (s *QuestionService) DeleteUserQues(UId, QId primitive.ObjectID) error {
	err := s.repo.DeleteOneDoc(bson.M{"q_id": QId, "user_id": UId})
	if err != nil {
		return err
	}
	return nil
}

func (s *QuestionService) GetPostQuestions(PId primitive.ObjectID) ([]*models.Question, error) {
	questions, err := s.repo.GetQuestionsByPId(PId)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *QuestionService) AddAnswer(QId primitive.ObjectID, answer string) error {
	err := s.repo.UpdateQuestion(QId, answer)
	if err != nil {
		return err
	}
	return nil
}
