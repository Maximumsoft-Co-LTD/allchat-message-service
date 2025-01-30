package http

import (
	"allchat-message-service/internal/adapter/handler/model"
	"allchat-message-service/internal/core/domain"
	"encoding/json"
)

func newTelegramWebhookReq(req model.TelegramWebhook) (domain.TelegramWebhook, error) {
	byteData, err := json.Marshal(req)
	if err != nil {
		return domain.TelegramWebhook{}, err
	}
	var data domain.TelegramWebhook
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return domain.TelegramWebhook{}, err
	}
	return data, nil
}
