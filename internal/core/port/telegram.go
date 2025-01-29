package port

import (
	"allchat-message-service/internal/core/domain"
	"context"
)

type TelegramRepository interface {
	InsertWebhookRawData(data any) error
}

type TelegramService interface {
	Webhook(c context.Context, id string, body domain.TelegramWebhook) error
	TestPublishQ(data any) error
	InsertWebhookRawData(data any) error
}
