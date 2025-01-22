package http

import (
	"allchat-message-service/internal/core/domain"
	"allchat-message-service/internal/core/port"
	"fmt"

	"github.com/gin-gonic/gin"
)

type TelegramHandler struct {
	svc port.TelegramService
}

func NewTelegramHandler(svc port.TelegramService) *TelegramHandler {
	return &TelegramHandler{
		svc,
	}
}

type SetWebhookRequest struct {
	BotToken string `json:"bot_token"`
}

func (h *TelegramHandler) Webhook(c *gin.Context) {
	var req domain.TelegramWebhookReq2
	// Set webhook
	if err := c.ShouldBind(&req); err != nil {
	}

	fmt.Println("handle webhook")
	// telegram := domain.Telegram
	h.svc.Webhook(c, req)
	ResData(c, 200, "success", "", nil)
}

func (h *TelegramHandler) TestPublishQ(c *gin.Context) {
	fmt.Println("test publish q")
	data := "hello world"
	if err := h.svc.TestPublishQ(data); err != nil {
		fmt.Println("publish q err", err)
	}
	ResData(c, 200, "success", "", nil)
}
