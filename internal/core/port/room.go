package port

import (
	"allchat-message-service/internal/core/domain"
	"context"
)

type RoomRepository interface {
	CreateRoom(c context.Context, room domain.Room) error
}
