package interfaces

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByUId(UId primitive.ObjectID) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByUsernamePassword(username string, password string) (*models.User, error)
	FindAdminByUsernamePassword(username string, password string) (*models.Admin, error)
	GetAllUsers() ([]*models.User, error)
	DeleteByUId(UId primitive.ObjectID) error
	UpdateActiveStatus(UId primitive.ObjectID, status bool) error
	ClearNotification(UId primitive.ObjectID) error
	PushNotification(UId primitive.ObjectID, title string) error
}
