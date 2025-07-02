package handler

import (
	"net/http"

	"github.com/Alfian57/belajar-golang/internal/dto"
	"github.com/Alfian57/belajar-golang/internal/response"
	"github.com/Alfian57/belajar-golang/internal/service"
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
	todos, err := h.service.GetAllTodos(ctx)
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

	if err := h.service.CreateTodo(ctx, request); err != nil {
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

	user, err := h.service.GetTodoByID(ctx, id)
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

	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}
	request.ID = id

	if err := h.service.UpdateTodo(ctx, request); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "todo successfully updated")
}

func (h *TodoHandler) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	if err := h.service.DeleteTodo(ctx, id); err != nil {
		response.WriteErrorResponse(ctx, err)
		return
	}

	response.WriteMessageResponse(ctx, http.StatusOK, "todo successfully deleted")
}
