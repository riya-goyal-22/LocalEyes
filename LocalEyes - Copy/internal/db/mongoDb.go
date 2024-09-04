package db

import (
	"database/sql"
	"fmt"
	"localEyes/constants"
	"log"
)

var dbClient *sql.DB

type SqlWrapper struct {
	db *sql.DB
}

func NewSqlWrapper(db *sql.DB) *SqlWrapper {
	return &SqlWrapper{db: db}
}

func (w SqlWrapper) InsertOne(query string, args ...interface{}) (sql.Result, error) {
	return w.db.Exec(query, args...)
}

func (w SqlWrapper) FindOne(query string, args ...interface{}) (sql.Result, error) {
	return w.db.Exec(query, args...)
}

func (w SqlWrapper) Find(query string, args ...interface{}) (*sql.Rows, error) {
	return w.db.Query(query, args...)
}

func (w SqlWrapper) Delete(query string, args ...interface{}) (sql.Result, error) {
	return w.db.Exec(query, args...)
}

func (w SqlWrapper) Update(query string, args ...interface{}) (sql.Result, error) {
	return w.db.Exec(query, args...)
}

func GetSQLClient() *sql.DB {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			constants.DBUser,
			constants.DBPassword,
			constants.DBHost,
			constants.DBPort,
			constants.DBName,
		)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		dbClient = db
	})
	return dbClient
}
