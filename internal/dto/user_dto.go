package dto

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username             string `form:"username" binding:"required,min=3"`
	Password             string `form:"password" binding:"required,min=8"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required,eqfield=Password"`
}

type UpdateUserRequest struct {
	ID       uuid.UUID `form:"id"`
	Username string    `form:"username" binding:"required,min=3"`
}
