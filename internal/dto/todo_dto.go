package dto

import "github.com/google/uuid"

type GetTodosFilter struct {
	Page      int    `json:"page" binding:"omitempty,number"`
	Limit     int    `json:"limit" binding:"omitempty,number"`
	Search    string `json:"search" binding:"omitempty"`
	OrderType string `json:"order_type" binding:"omitempty,oneof=ASC DESC"`
}

type CreateTodoRequest struct {
	Todo string `form:"todo" binding:"required,min=3,max=100"`
}

type UpdateTodoRequest struct {
	ID   uuid.UUID `form:"id"`
	Todo string    `form:"todo" binding:"required,min=3"`
}
