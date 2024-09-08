//go:build !test
// +build !test

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"localEyes/cmd/cli"
	"localEyes/constants"
	"localEyes/internal/repositories"
	"localEyes/internal/services"
	"log"
	"sync"
)

var dbClient *sql.DB
var once sync.Once

func init() {
	GetSQLClient()
}

func main() {
	userRepo := repositories.NewMySQLUserRepository(dbClient)
	userService := services.NewUserService(userRepo)

	postRepo := repositories.NewMySQLPostRepository(dbClient)
	postService := services.NewPostService(postRepo)

	questionRepo := repositories.NewMySQLQuestionRepository(dbClient)
	questionService := services.NewQuestionService(questionRepo)

	adminService := services.NewAdminService(userRepo, postRepo, questionRepo)

	cli.RootCli(userService, postService, questionService, adminService)
}

func GetSQLClient() {
	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
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
}
