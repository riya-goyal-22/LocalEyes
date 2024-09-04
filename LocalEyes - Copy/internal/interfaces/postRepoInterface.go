package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/internal/models"
)

type PostRepository interface {
	Create(post *models.Post) error
	GetAllPosts() ([]*models.Post, error)
	DeleteOneDoc(filter interface{}) error
	DeleteByUId(UId primitive.ObjectID) error
	GetPostsByFilter(filter interface{}) ([]*models.Post, error)
	UpdateUserPost(PId primitive.ObjectID, UId primitive.ObjectID, title string, content string) error
	UpdateLike(PId primitive.ObjectID) error
}
