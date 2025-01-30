package httprequest

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type httpReqTelegram struct {
}

func NewHTTPReqTelegram() *httpReqTelegram {
	return &httpReqTelegram{}
}

func (rt *httpReqTelegram) GetUserImageTelegramUrl(token string, id int64) (string, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return "", err
	}
	userID := int(id)
	userRequest := tgbotapi.NewUserProfilePhotos(userID)

	// Send the user profile request to Telegram
	userPhotos, err := bot.GetUserProfilePhotos(userRequest)
	if err != nil {
		return "", err
	}

	// Get the user's profile picture URL if available
	if userPhotos.TotalCount > 0 {
		// Get the largest profile photo available
		photoSize := userPhotos.Photos[0][len(userPhotos.Photos[0])-1]

		// Download the profile photo
		photoFile, err := bot.GetFile(tgbotapi.FileConfig{
			FileID: photoSize.FileID,
		})
		if err != nil {
			return "", err
		}

		// Print the photo URL to the console
		return photoFile.Link(bot.Token), nil
	} else {
		e := fmt.Sprintf("User %d has no profile photos", userID)
		return "", errors.New(e)
	}
}

func (rt *httpReqTelegram) GetFileTelegramFromUrl(path string) ([]byte, error) {
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err2 := client.Do(req)
	if err2 != nil {
		return nil, err2
	}
	defer res.Body.Close()
	fmt.Println(req.Body)
	fileBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err2
	}
	return fileBytes, nil
}
