package port

import (
	"allchat-message-service/internal/core/domain"
	"context"
)

type TelegramRepository interface {
	InsertWebhookRawData(data any) error
}

type TelegramService interface {
	Webhook(c context.Context, body domain.TelegramWebhookReq2) error
	InsertWebhookRawData(data any) error
}
