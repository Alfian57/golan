package dto

type LoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type RegisterRequest struct {
	Username             string `form:"username" binding:"required,min=3,max=100"`
	Password             string `form:"password" binding:"required,min=8"`
	PasswordConfirmation string `form:"password_confirmation" binding:"required,eqfield=Password"`
}
