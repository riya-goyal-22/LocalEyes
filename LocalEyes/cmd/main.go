package main

import (
	"localEyes/cmd/cli"
	"localEyes/internal/repositories"
	"localEyes/internal/services"
)

func main() {
	userRepo := repositories.NewMongoUserRepository()
	userService := services.NewUserService(userRepo)

	postRepo := repositories.NewMongoPostRepository()
	postService := services.NewPostService(postRepo)

	questionRepo := repositories.NewMongoQuestionRepository()
	questionService := services.NewQuestionService(questionRepo)

	adminService := services.NewAdminService(userRepo, postRepo, questionRepo)

	cli.RootCli(userService, postService, questionService, adminService)
}
