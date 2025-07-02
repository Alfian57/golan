package router

import (
	"github.com/Alfian57/belajar-golang/internal/handler"
	"github.com/Alfian57/belajar-golang/internal/repository"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/gin-gonic/gin"
)

func RegisterV1Route(router *gin.RouterGroup) {

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/users", userHandler.GetAllUsers)
	router.POST("/users", userHandler.CreateUser)
	router.GET("/users/:id", userHandler.GetUserByID)
	router.PUT("/users/:id", userHandler.UpdateUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
}
