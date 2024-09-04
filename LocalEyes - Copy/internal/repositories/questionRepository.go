package repositories

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"localEyes/internal/db"
	"localEyes/internal/models"
)

type MySQLQuestionRepository struct {
	DB *sql.DB
}

func NewMySQLQuestionRepository() *MySQLQuestionRepository {
	return &MySQLQuestionRepository{
		DB: db.GetSQLClient(),
	}
}

func (r *MySQLQuestionRepository) Create(question *models.Question) error {
	query := "INSERT INTO questions (q_id, post_id,user_id, text, replies,created_at) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, question.QId, question.PostId, question.UserId, question.Text, question.Replies, question.CreatedAt)
	return err
}

func (r *MySQLQuestionRepository) GetAllQuestions() ([]*models.Question, error) {
	query := "SELECT q_id, post_id,user_id, text, replies,created_at FROM questions"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*models.Question
	//for rows.Next() {
	//	var question models.Question
	//	if err := rows.Scan(&question.QId,&question.PostId, &question.Text, &question.Replies,&question.CreatedAt); err != nil {
	//		return nil, err
	//	}
	//	questions = append(questions, &question)
	//}
	err = sqlx.StructScan(rows, &questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *MySQLQuestionRepository) DeleteByQIdUId(questionID string) error {
	query := "DELETE FROM questions WHERE q_id = ? AND user_id = ?"
	_, err := r.DB.Exec(query, questionID)
	return err
}

func (r *MySQLQuestionRepository) DeleteByPId(PId string) error {
	query := "DELETE FROM questions WHERE post_id = ?"
	_, err := r.DB.Exec(query, PId)
	return err
}

func (r *MySQLQuestionRepository) GetQuestionsByPId(PId string) ([]*models.Question, error) {
	query := "SELECT q_id, post_id,user_id, text, replies ,created_at FROM questions WHERE post_id = ?"
	rows, err := r.DB.Query(query, PId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*models.Question
	//for rows.Next() {
	//	var question models.Question
	//	if err := rows.Scan(&question.ID, &question.PostID, &question.QuestionText, &question.Replies); err != nil {
	//		return nil, err
	//	}
	//	questions = append(questions, &question)
	//}
	err = sqlx.StructScan(rows, &questions)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (r *MySQLQuestionRepository) UpdateQuestion(QId, answer string) error {
	query := "UPDATE questions SET replies = CONCAT(replies, ?) WHERE id = ?"
	_, err := r.DB.Exec(query, answer, QId)
	return err
}
