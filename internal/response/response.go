package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteDataResponse(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{"data": data})
}

func WriteMessageResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(statusCode, gin.H{"message": message})
}

func WriteErrorResponse(ctx *gin.Context, err error) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}
