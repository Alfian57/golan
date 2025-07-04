package dto

import "github.com/google/uuid"

type CreateTodoRequest struct {
	Todo string `form:"todo" binding:"required,min=3,max=100"`
}

type UpdateTodoRequest struct {
	ID   uuid.UUID `form:"id"`
	Todo string    `form:"todo" binding:"required,min=3"`
}
