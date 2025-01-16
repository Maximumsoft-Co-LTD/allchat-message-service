package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResData(ctx *gin.Context, status int, message string, errorText string, data any) {
	rsp := gin.H{"code": status, "message": message, "error": errorText, "payload": data}
	ctx.JSON(http.StatusOK, rsp)
}
