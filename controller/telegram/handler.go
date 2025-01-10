package telegram

import (
	db "allchat-message-service/_config"
	"allchat-message-service/service/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlerSetWebhook(r *mongo.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		botToken := "dfghjkmghjklx"
		setWebhookTelegram(botToken)
		c.JSON(http.StatusOK, gin.H{
			`message`: `Hello World`,
		})
	}
}

func HandlerWebhook(r *mongo.Database) func(c *gin.Context) {
	return func(c *gin.Context) {
		var body RequestWebhookTelegram
		if err := c.ShouldBindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				`message`: `Failed to bind JSON`,
				`error`:   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			`message`: `Telegram Webhook`,
			"data":    body,
		})
	}
}

func VerifyWebhookTelegram(resource *db.Resource) func(c *gin.Context) {
	return func(c *gin.Context) {
		// var body interface{}
		var body RequestWebhookTelegram
		if err := c.ShouldBindJSON(&body); err != nil {
			c.AbortWithStatusJSON(http.StatusOK, util.ResData(1, "Data Error", err.Error(), nil))
			return
		}
		// jsonBody, err := json.Marshal(body)
		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusInternalServerError, helper.ResData(1, "Internal Server Error", err.Error(), nil))
		// 	return
		// }
		// println(string(jsonBody))
		// repo.CreateStatement(resource, "test_telegram", body) //<============
		//fmt.Println("dw")
		fmt.Println("1")
		// repo.CreateStatement(resource, "telegram_webhook", body)
		// go manageMessageTelegramLocal(resource, body, c.Param("id"))
		c.JSON(http.StatusOK, util.ResData(0, "Success", "", nil))
	}
}
