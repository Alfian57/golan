package router

import (
	"github.com/Alfian57/belajar-golang/internal/handler"
	"github.com/Alfian57/belajar-golang/internal/middleware"
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

	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
	router.POST("/refreh", authHandler.Refresh)
	router.POST("/logout", authHandler.Logout)

	users := router.Group("users", middleware.AuthMiddleware())
	{
		users.GET("/", userHandler.GetAllUsers)
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
	// router.DELETE("/users/:id/todos", userHandler.GetUserTodos)

	todos := router.Group("todos", middleware.AuthMiddleware())
	{
		todos.GET("/", todoHandler.GetAlltodos)
		todos.POST("/", todoHandler.CreateTodo)
		todos.GET("/:id", todoHandler.GetTodoByID)
		todos.PUT("/:id", todoHandler.UpdateTodo)
		todos.DELETE("/:id", todoHandler.DeleteTodo)
	}
}
