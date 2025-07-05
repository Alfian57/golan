package dto

import "github.com/google/uuid"

type CreateTodoRequest struct {
	Todo string `form:"todo" binding:"required,min=3,max=100"`
}

type UpdateTodoRequest struct {
	ID   uuid.UUID `form:"id"`
	Todo string    `form:"todo" binding:"required,min=3"`
}

type GetTodosFilter struct {
	PaginationRequest
	Search    string `json:"search" form:"search" binding:"omitempty,max=255"`
	OrderBy   string `json:"order_by" form:"order_by" binding:"omitempty,oneof=todo created_at"`
	OrderType string `json:"order_type" form:"order_type" binding:"omitempty,oneof=ASC DESC asc desc"`
}
