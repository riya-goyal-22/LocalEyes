package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"localEyes/internal/db"
	"localEyes/internal/models"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository() *MySQLUserRepository {
	return &MySQLUserRepository{
		DB: db.GetSQLClient(),
	}
}

func (r *MySQLUserRepository) Create(user *models.User) error {
	query := "INSERT INTO users (id, username, password, is_active, city, dwelling_age, tag, notification) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err := r.DB.Exec(query, user.UId, user.Username, user.Password, user.IsActive, user.City, user.DwellingAge, user.Tag, user.Notification)
	return err
}

func (r *MySQLUserRepository) FindByUId(UId string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, password, is_active, city, dwelling_age, tag, notification FROM users WHERE id = ?"
	err := r.DB.QueryRow(query, UId).Scan(&user.UId, &user.Username, &user.Password, &user.IsActive, &user.City, &user.DwellingAge, &user.Tag, &user.Notification)
	return &user, err
}

func (r *MySQLUserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, password, is_active, city, dwelling_age, tag, notification FROM users WHERE username = ?"
	err := r.DB.QueryRow(query, username).Scan(&user.UId, &user.Username, &user.Password, &user.IsActive, &user.City, &user.DwellingAge, &user.Tag, &user.Notification)
	return &user, err
}

func (r *MySQLUserRepository) FindByUsernamePassword(username, password string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, password, is_active, city, dwelling_age, tag, notification FROM users WHERE username = ? AND password = ?"
	err := r.DB.QueryRow(query, username, password).Scan(&user.UId, &user.Username, &user.Password, &user.IsActive, &user.City, &user.DwellingAge, &user.Tag, &user.Notification)
	return &user, err
}

func (r *MySQLUserRepository) FindAdminByUsernamePassword(username, password string) (*models.Admin, error) {
	var admin models.Admin
	query := "SELECT id, username, password FROM admins WHERE username = ? AND password = ?"
	err := r.DB.QueryRow(query, username, password).Scan(&admin.User.UId, &admin.User.Username, &admin.User.Password)
	return &admin, err
}

func (r *MySQLUserRepository) GetAllUsers() ([]*models.User, error) {
	query := "SELECT id, username, password, is_active, city, dwelling_age, tag, notification FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	//for rows.Next() {
	//	var user models.User
	//	if err := rows.Scan(&user.UId, &user.Username, &user.Password, &user.IsActive, &user.City, &user.DwellingAge, &user.Tag, &user.Notification); err != nil {
	//		return nil, err
	//	}
	//	users = append(users, &user)
	//}
	err = sqlx.StructScan(rows, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *MySQLUserRepository) DeleteByUId(UId string) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := r.DB.Exec(query, UId)
	return err
}

func (r *MySQLUserRepository) UpdateActiveStatus(UId string, status bool) error {
	query := "UPDATE users SET is_active = ? WHERE id = ?"
	_, err := r.DB.Exec(query, status, UId)
	return err
}

func (r *MySQLUserRepository) PushNotification(UId string, title string) error {
	query := "UPDATE users SET notification = CONCAT(COALESCE(notification, ''), ?) WHERE id != ?"
	notification := "New post: " + title + "\n"
	_, err := r.DB.Exec(query, notification, UId)
	return err
}

func (r *MySQLUserRepository) ClearNotification(UId string) error {
	query := "UPDATE users SET notification = '' WHERE id = ?"
	_, err := r.DB.Exec(query, UId)
	return err
}
