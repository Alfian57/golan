//go:build wireinject
// +build wireinject

package di

import (
	"github.com/Alfian57/belajar-golang/internal/handler"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/google/wire"
)

func InitializeAuthHandler() *handler.AuthHandler {
	wire.Build(handler.NewAuthHandler, service.NewAuthService, repository.NewUserRepository, repository.NewRefreshTokenRepository)
	return &handler.AuthHandler{}
}

func InitializeUserHandler() *handler.UserHandler {
	wire.Build(handler.NewUserHandler, service.NewUserService, repository.NewUserRepository)
	return &handler.UserHandler{}
}

func InitializeTodoHandler() *handler.TodoHandler {
	wire.Build(handler.NewTodoHandler, service.NewTodoService, repository.NewTodoRepository, repository.NewUserRepository)
	return &handler.TodoHandler{}
}
