package port

import (
	"allchat-message-service/internal/core/domain"
	"context"
)

// type TelegramRepository interface {
// }

type TelegramService interface {
	Webhook(c context.Context, body domain.TelegramWebhookReq2) error
}
