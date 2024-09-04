package repositories

import (
	"database/sql"
	"localEyes/internal/db"
	"localEyes/internal/models"
)

type MySQLPostRepository struct {
	DB *sql.DB
}

func NewMySQLPostRepository() *MySQLPostRepository {
	return &MySQLPostRepository{
		DB: db.GetSQLClient(),
	}
}

func (r *MySQLPostRepository) Create(post *models.Post) error {
	query := "INSERT INTO posts (post_id, user_id, title,type, content, likes,created_at) VALUES (?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, post.PostId, post.UId, post.Title, post.Type, post.Content, post.Likes, post.CreatedAt)
	return err
}

func (r *MySQLPostRepository) GetAllPosts() ([]*models.Post, error) {
	query := "SELECT post_id, user_id, title, type,content, likes,created_at FROM posts"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.PostId, &post.UId, &post.Title, &post.Content, &post.Likes); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (r *MySQLPostRepository) DeleteByUIdPId(UId, PId string) error {
	query := "DELETE FROM posts WHERE user_id = ? AND post_id = ?"
	_, err := r.DB.Exec(query, UId, PId)
	return err
}

func (r *MySQLPostRepository) DeleteByUId(UId string) error {
	query := "DELETE FROM posts WHERE user_id = ?"
	_, err := r.DB.Exec(query, UId)
	return err
}

func (r *MySQLPostRepository) GetPostsByFilter(filter string) ([]*models.Post, error) {
	query := "SELECT post_id, user_id, title,type, content, likes,created_at FROM posts WHERE type = ?"
	rows, err := r.DB.Query(query, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.PostId, &post.UId, &post.Title, &post.Content, &post.Likes); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (r *MySQLPostRepository) GetPostsByUId(UId string) ([]*models.Post, error) {
	query := "SELECT post_id, user_id, title,type, content, likes,created_at FROM posts WHERE user_id = ?"
	rows, err := r.DB.Query(query, UId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.PostId, &post.UId, &post.Title, &post.Content, &post.Likes); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (r *MySQLPostRepository) UpdateUserPost(PId, UId, title, content string) error {
	query := "UPDATE posts SET title = ?, content = ? WHERE post_id = ? AND user_id = ?"
	_, err := r.DB.Exec(query, title, content, PId, UId)
	return err
}

func (r *MySQLPostRepository) UpdateLike(PId string) error {
	query := "UPDATE posts SET likes = likes + 1 WHERE post_id = ?"
	_, err := r.DB.Exec(query, PId)
	return err
}
