package router

import (
	db "allchat-message-service/_config"
	"allchat-message-service/controller/telegram"

	"github.com/gin-gonic/gin"
)

func Webhook(r *gin.Engine, resource *db.Resource) {

	webhook := r.Group("/webhook")

	webhook.POST("/telegram-bot-update/:id", telegram.VerifyWebhookTelegram(resource)) // logs jaeger
}
