package router

import (
	"github.com/Alfian57/belajar-golang/internal/di"
	"github.com/Alfian57/belajar-golang/internal/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterV1Route(router *gin.RouterGroup) {

	authHandler := di.InitializeAuthHandler()
	userHandler := di.InitializeUserHandler()
	todoHandler := di.InitializeTodoHandler()

	router.POST("/login", authHandler.Login)
	router.POST("/register", authHandler.Register)
	router.POST("/refresh", middleware.AuthMiddleware(), authHandler.Refresh)
	router.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)

	users := router.Group("users", middleware.AuthMiddleware())
	{
		users.GET("/", userHandler.GetAllUsers)
		users.POST("/", userHandler.CreateUser)
		users.GET("/:id", userHandler.GetUserByID)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	todos := router.Group("todos", middleware.AuthMiddleware())
	{
		todos.GET("/", todoHandler.GetAlltodos)
		todos.POST("/", todoHandler.CreateTodo)
		todos.GET("/:id", todoHandler.GetTodoByID)
		todos.PUT("/:id", todoHandler.UpdateTodo)
		todos.DELETE("/:id", todoHandler.DeleteTodo)
	}
}
