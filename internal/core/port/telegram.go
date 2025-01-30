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

type TelegramHTTPReq interface {
	GetUserImageTelegramUrl(token string, id int64) (string, error)
	GetFileTelegramFromUrl(path string) ([]byte, error)
}
