package util

import "github.com/gin-gonic/gin"

func ResData(status int, message string, errorText string, data interface{}) (jsonObj any) {
	return gin.H{"code": status, "message": message, "error": errorText, "payload": data}
}
