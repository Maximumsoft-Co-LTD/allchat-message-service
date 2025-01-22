package pubsub

import (
	"allchat-message-service/internal/core/port"
	"fmt"
)

type TelegramSubscriber struct {
	tsv port.TelegramService
}

func NewTelegramSubscriber(tsv port.TelegramService) *TelegramSubscriber {
	return &TelegramSubscriber{
		tsv,
	}
}

func (h *TelegramSubscriber) InsertWebhookRawData(data []any) {
	fmt.Println("receive telegram raw data")
	// telegram := domain.Telegram
	if err := h.tsv.InsertWebhookRawData(data); err != nil {
		fmt.Println("insert telegram webhook data err", err)
		return
	}
}
