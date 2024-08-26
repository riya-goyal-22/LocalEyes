package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"localEyes/internal/models"
	"localEyes/internal/repositories"
	"time"
)

type PostService struct {
	repo repositories.PostRepository
}

func NewPostService(repo repositories.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) CreatePost(userId primitive.ObjectID, title, content, postType string) error {
	post := &models.Post{
		UId:       userId,
		PostId:    primitive.NewObjectID(),
		Title:     title,
		Content:   content,
		Type:      postType,
		CreatedAt: time.Now(),
		Likes:     0,
	}
	return s.repo.Create(post)
}

func (s *PostService) UpdateMyPost(postId, userId primitive.ObjectID, title, content string) error {
	err := s.repo.UpdateUserPost(postId, userId, title, content)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GiveAllPosts() ([]*models.Post, error) {
	posts, err := s.repo.GetAllPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) GiveMyPosts(UId primitive.ObjectID) ([]*models.Post, error) {
	filter := bson.M{"userId": UId}
	posts, err := s.repo.GetPostsByFilter(filter)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) DeleteMyPost(UId, PId primitive.ObjectID) error {
	err := s.repo.DeleteOneDoc(bson.M{"userId": UId, "id": PId})
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) Like(PId primitive.ObjectID) error {
	err := s.repo.UpdateLike(PId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GiveFilteredPosts(filterType string) ([]*models.Post, error) {
	filter := bson.M{"type": filterType}
	posts, err := s.repo.GetPostsByFilter(filter)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (s *PostService) PostIdExist(PId primitive.ObjectID) (bool, error) {
	filter := bson.M{"id": PId}
	posts, err := s.repo.GetPostsByFilter(filter)
	if err != nil {
		return false, err
	}
	return len(posts) > 0, nil
}
