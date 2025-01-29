package port

import (
	"allchat-message-service/internal/core/domain"
	"context"
)

type PlatformRepository interface {
	GetInfoPlatformByToken(c context.Context, Token string) (domain.Platform, error)
}
