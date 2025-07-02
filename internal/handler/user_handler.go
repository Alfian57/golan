package handler

import (
	"net/http"

	"github.com/Alfian57/belajar-golang/internal/dto"
	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/response"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service: s,
	}
}

func (s *UserHandler) GetAllUsers(ctx *gin.Context) {
	users, err := s.service.GetAllUsers(ctx)
	if err != nil {
		response.WriteErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	response.WriteDataResponse(ctx, http.StatusOK, users)
}

func (s *UserHandler) CreateUser(ctx *gin.Context) {
	var request dto.CreateUserRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.WriteValidationError(ctx, err)
		return
	}

	err := s.service.CreateUser(ctx, request)
	if err != nil {
		response.WriteErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusCreated, "user successfully created")
}

func (s *UserHandler) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := s.service.GetUserByID(ctx, id)
	if err != nil {
		if err == errs.ErrUserNotFound {
			response.WriteErrorResponse(ctx, http.StatusNotFound, err)
			return
		} else {
			response.WriteErrorResponse(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	response.WriteDataResponse(ctx, http.StatusOK, user)
}

func (s *UserHandler) UpdateUser(ctx *gin.Context) {
	var request dto.UpdateUserRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.WriteValidationError(ctx, err)
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.WriteErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}
	request.ID = id

	err = s.service.UpdateUser(ctx, request)
	if err != nil {
		response.WriteErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "user successfully updated")
}

func (s *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	err := s.service.DeleteUser(ctx, id)
	if err != nil {
		if err == errs.ErrUserNotFound {
			response.WriteErrorResponse(ctx, http.StatusNotFound, err)
			return
		} else {
			response.WriteErrorResponse(ctx, http.StatusInternalServerError, err)
			return
		}
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "user successfully deleted")
}
