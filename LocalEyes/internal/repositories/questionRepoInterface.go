package repositories

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/internal/models"
)

type QuestionRepository interface {
	Create(question *models.Question) error
	GetAllQuestions() ([]*models.Question, error)
	DeleteOneDoc(filter interface{}) error
	DeleteByPId(PId primitive.ObjectID) error
	GetQuestionsByPId(PId primitive.ObjectID) ([]*models.Question, error)
	UpdateQuestion(QId primitive.ObjectID, answer string) error
}
