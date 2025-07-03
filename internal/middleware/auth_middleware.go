package middleware

import (
	"net/http"

	"github.com/Alfian57/belajar-golang/internal/response"
	"github.com/Alfian57/belajar-golang/internal/utils/jwt"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("access_token")
		if err != nil {
			response.WriteErrorResponse(ctx, err)
		}

		isValid, err := jwt.ValidateAccessToken(accessToken)
		if err != nil {
			response.WriteErrorResponse(ctx, err)
		}

		if !isValid {
			response.WriteMessageResponse(ctx, http.StatusUnauthorized, "unauthorized")
		}

		ctx.Next()
	}
}
