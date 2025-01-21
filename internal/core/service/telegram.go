package service

import (
	"allchat-message-service/internal/core/domain"
	"allchat-message-service/internal/core/port"
	"allchat-message-service/internal/core/util"
	"context"
	"fmt"
	"time"
)

type TelegramService struct {
	tRepo port.TelegramRepository
	rRepo port.RoomRepository
	cache port.CacheRepository
}

func NewTelegramService(tRepo port.TelegramRepository, rRepo port.RoomRepository, cache port.CacheRepository) *TelegramService {
	return &TelegramService{
		tRepo,
		rRepo,
		cache,
	}
}

func (s *TelegramService) Webhook(c context.Context, body domain.TelegramWebhookReq2) error {
	// Webhook from telegram
	fmt.Println("service webhook")
	if err := s.rRepo.CreateRoom(c, domain.Room{}); err != nil {
		return err
	}
	webhookByte, err := util.Serialize(1)
	if err != nil {
		return err
	}
	if err := s.cache.Set(c, "webhook", webhookByte, 1*time.Minute); err != nil {
		return err
	}
	cacheData, err := s.cache.Get(c, "webhook")
	if err != nil {
		return err
	}
	data := 0
	if err := util.Deserialize(cacheData, &data); err != nil {
		return err
	}
	fmt.Println("webhook data", data)
	return nil
}

func (s *TelegramService) InsertWebhookRawData(data any) error {
	if err := s.tRepo.InsertWebhookRawData(data); err != nil {
		return err
	}
	return nil
}
