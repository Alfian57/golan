package router

import (
	"github.com/Alfian57/belajar-golang/internal/handler"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterV1Route(router *gin.RouterGroup) {

	refreshTokenRepository := repository.NewRefreshTokenRepository()

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepository, refreshTokenRepository)
	authHandler := handler.NewAuthHandler(authService)

	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, userRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	router.GET("/users", userHandler.GetAllUsers)
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUserByID)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	// router.DELETE("/users/:id/todos", userHandler.GetUserTodos)

	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
	router.POST("/refreh", authHandler.Refresh)
	router.POST("/logout", authHandler.Logout)

	router.GET("/todos", todoHandler.GetAlltodos)
	router.POST("/todos", todoHandler.CreateTodo)
	router.GET("/todos/:id", todoHandler.GetTodoByID)
	router.PUT("/todos/:id", todoHandler.UpdateTodo)
	router.DELETE("/todos/:id", todoHandler.DeleteTodo)
}
