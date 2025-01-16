package service

import (
	"allchat-message-service/internal/core/domain"
	"allchat-message-service/internal/core/port"
	"context"
	"fmt"
)

// type TelegramService struct {
// 	tRepo port.TelegramRepository
// 	rRepo port.RoomRepository
// 	cache port.CacheRepository
// }

// func NewTelegramService(tRepo port.TelegramRepository, rRepo port.RoomRepository, cache port.CacheRepository) *TelegramService {
// 	return &TelegramService{
// 		tRepo,
// 		rRepo,
// 		cache,
// 	}
// }

type TelegramService struct {
	tRepo port.TelegramRepository
	rRepo port.RoomRepository
}

func NewTelegramService(tRepo port.TelegramRepository, rRepo port.RoomRepository) *TelegramService {
	return &TelegramService{
		tRepo,
		rRepo,
	}
}

func (s *TelegramService) Webhook(c context.Context, body domain.TelegramWebhookReq2) error {
	// Webhook from telegram
	fmt.Println("service webhook")
	if err := s.rRepo.CreateRoom(c, domain.Room{}); err != nil {
		return err
	}
	return nil
}
