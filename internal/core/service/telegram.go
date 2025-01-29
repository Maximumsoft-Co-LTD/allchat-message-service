package service

import (
	"allchat-message-service/internal/core/domain"
	"allchat-message-service/internal/core/port"
	"allchat-message-service/internal/core/util"
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

type TelegramService struct {
	tRepo port.TelegramRepository
	rRepo port.RoomRepository
	pRepo port.PlatformRepository
	cache port.CacheRepository
	amqp  port.AmqpPublisher
}

func NewTelegramService(tRepo port.TelegramRepository, rRepo port.RoomRepository, pRepo port.PlatformRepository, cache port.CacheRepository, amqp port.AmqpPublisher) *TelegramService {
	return &TelegramService{
		tRepo,
		rRepo,
		pRepo,
		cache,
		amqp,
	}
}

func (s *TelegramService) Webhook(c context.Context, id string, body domain.TelegramWebhook) error {
	// // Webhook from telegram
	// fmt.Println("service webhook")
	// if err := s.rRepo.CreateRoom(c, domain.Room{}); err != nil {
	// 	return err
	// }
	// webhookByte, err := util.Serialize(1)
	// if err != nil {
	// 	return err
	// }
	// if err := s.cache.Set(c, "webhook", webhookByte, 1*time.Minute); err != nil {
	// 	return err
	// }
	// cacheData, err := s.cache.Get(c, "webhook")
	// if err != nil {
	// 	return err
	// }
	// data := 0
	// if err := util.Deserialize(cacheData, &data); err != nil {
	// 	return err
	// }
	// fmt.Println("webhook data", data)

	// receive := time.Now().Unix()
	platform, _ := s.pRepo.GetInfoPlatformByToken(c, id)

	if !reflect.ValueOf(body.Message).IsZero() || !reflect.ValueOf(body.EditedMessage).IsZero() { // กรณีข้อความทั้วไป
		reqMessage := body.Message
		if !reflect.ValueOf(body.EditedMessage).IsZero() {
			reqMessage = body.EditedMessage
		}
		//TODO: ทำถึงฟังก์ชันนี้
		imageUrl, _ := s.getUserImageTelegram(c, platform, reqMessage.From)

		name := reqMessage.Chat.FirstName
		if name == "" {
			name = reqMessage.Chat.Title
		}
		chatID := strconv.Itoa(int(reqMessage.Chat.ID))
		roomID := platform.PageID + "_" + chatID

		fmt.Println("tmp : ", imageUrl, roomID)

		// member := model.MemberAccount{
		// 	Type:          "",
		// 	NickName:      name,
		// 	FirstName:     name,
		// 	LastName:      "",
		// 	FullName:      name,
		// 	PhoneNumber:   "",
		// 	Email:         "",
		// 	ProfileUrl:    imageUrl,
		// 	SocialID:      strconv.Itoa(int(reqMessage.From.ID)),
		// 	PageID:        platform.PageID,
		// 	RoomID:        roomID,
		// 	StatusMessage: "",
		// 	Platform:      PlatTelegram,
		// 	LastDate:      time.Now(),
		// 	Datetime:      time.Unix(int64(reqMessage.Date), 0),
		// 	LastMessage:   GetLastMessageTelegram(body),
		// 	BusinessCode:  platform.BusinessCode,
		// 	Prefix:        platform.Prefix,
		// 	PageName:      platform.Name,
		// 	PageImg:       platform.ImgUrl,
		// 	IsRead:        false,
		// }

		// saveLastMsgIDTelegramRedis(roomID, reqMessage.MessageID)

		//TODO: มีต่ออีก

	}

	return nil
}

func (s *TelegramService) TestPublishQ(data any) error {
	// Publish data to queue
	fmt.Println("service publish q")
	byteData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if err := s.amqp.Publish("raw_data", "raw_data", byteData); err != nil {
		return err
	}
	return nil
}

func (s *TelegramService) InsertWebhookRawData(data any) error {
	if err := s.tRepo.InsertWebhookRawData(data); err != nil {
		return err
	}
	return nil
}

func (s *TelegramService) getUserImageTelegram(c context.Context, platform domain.Platform, form domain.TelegramFrom) (string, error) {
	key := strconv.Itoa(int(form.ID))
	// redisConnect := service.RedisConnect()
	// redisConnect.GetRedis(key, &imageUrlRedis)
	// defer redisConnect.CloseRedis()

	cacheData, err := s.cache.Get(c, key)
	if err != nil {
		return "", err
	}
	imageUrlRedis := ""
	if err := util.Deserialize(cacheData, &imageUrlRedis); err != nil {
		return "", err
	}

	if imageUrlRedis != "" {
		return imageUrlRedis, nil // if there is data in redis return data
	}

	imageUrl, err := util.GetUserImageTelegramUrl(platform.AccessToken, form.ID)
	if err != nil {
		imageUrl = ""
	} else {
		//TODO:กลับมาทำต่อ
		// imageByte, err := util.GetFileTelegramFromUrl(imageUrl)
		// if err != nil {
		// 	imageUrl = ""
		// } else {
		// 	//TODO: ทำถึงฟังก์ชันนี้
		// 	url, _ := helper.UploadFile(strconv.Itoa(int(form.ID)), imageByte, platform.PageID, "")
		// 	imageUrl = url
		// }
	}
	//TODO: ทำต่อ
	// go SaveuserImageTelegramToRadis(imageUrl, key)
	return imageUrl, nil
}
