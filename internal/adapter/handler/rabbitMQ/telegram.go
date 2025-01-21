package rabbitmq

import (
	"allchat-message-service/internal/core/port"
	"fmt"
)

type TelegramSubscribe struct {
	tsv port.TelegramService
}

func NewTelegramSubscriber(tsv port.TelegramService) *TelegramSubscribe {
	return &TelegramSubscribe{
		tsv,
	}
}

func (h *TelegramSubscribe) InsertWebhookRawData(data any) {
	fmt.Println("receive telegram raw data")
	// telegram := domain.Telegram
	if err := h.tsv.InsertWebhookRawData(data); err != nil {
		fmt.Println("insert telegram webhook data err", err)
		return
	}
}
