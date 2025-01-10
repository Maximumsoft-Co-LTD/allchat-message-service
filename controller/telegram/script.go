package telegram

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	APISetWebhookTelegram = "https://api.telegram.org/bot%s/setWebhook?url=%s"
)

func setWebhookTelegram(token string) (string, error) {
	url := fmt.Sprintf(APISetWebhookTelegram, token, "")
	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, _ := client.Do(req)
	if res.StatusCode == http.StatusOK {
		defer res.Body.Close()
		return "Webhook set successfully", nil
	} else {
		defer res.Body.Close()
		return "", errors.New("failed to set webhook")
	}
}
