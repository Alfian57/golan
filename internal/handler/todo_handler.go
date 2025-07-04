package handler

import (
	"net/http"

	"github.com/Alfian57/belajar-golang/internal/dto"
	errs "github.com/Alfian57/belajar-golang/internal/errors"
	"github.com/Alfian57/belajar-golang/internal/response"
	"github.com/Alfian57/belajar-golang/internal/service"
	"github.com/Alfian57/belajar-golang/internal/utils/auth"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: s,
	}
}

func (h *TodoHandler) GetAlltodos(ctx *gin.Context) {
	var query dto.GetTodosFilter
	if err := ctx.ShouldBindQuery(&query); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	currentUser, exist := auth.GetCurrentUser(ctx)
	if !exist {
		response.WriteErrorResponse(ctx, errs.ErrUnauthorized)
		return
	}

	todos, err := h.service.GetAllTodos(ctx, currentUser, query)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteDataResponse(ctx, http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(ctx *gin.Context) {
	var request dto.CreateTodoRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	currentUser, exist := auth.GetCurrentUser(ctx)
	if !exist {
		response.WriteErrorResponse(ctx, errs.ErrUnauthorized)
		return
	}

	if err := h.service.CreateTodo(ctx, currentUser, request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusCreated, "todo successfully created")
}

func (h *TodoHandler) GetTodoByID(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	currentUser, exist := auth.GetCurrentUser(ctx)
	if !exist {
		response.WriteErrorResponse(ctx, errs.ErrUnauthorized)
		return
	}

	user, err := h.service.GetTodoByID(ctx, currentUser, id)
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteDataResponse(ctx, http.StatusOK, user)
}

func (h *TodoHandler) UpdateTodo(ctx *gin.Context) {
	var request dto.UpdateTodoRequest
	if err := ctx.ShouldBind(&request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	currentUser, exist := auth.GetCurrentUser(ctx)
	if !exist {
		response.WriteErrorResponse(ctx, errs.ErrUnauthorized)
		return
	}

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}
	request.ID = id

	if err := h.service.UpdateTodo(ctx, currentUser, request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "todo successfully updated")
}

func (h *TodoHandler) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	currentUser, exist := auth.GetCurrentUser(ctx)
	if !exist {
		response.WriteErrorResponse(ctx, errs.ErrUnauthorized)
		return
	}

	if _, err := uuid.Parse(id); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	if err := h.service.DeleteTodo(ctx, currentUser, id); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "todo successfully deleted")
}
