package http

import (
	"allchat-message-service/internal/adapter/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	telegramHandler TelegramHandler,
) (*Router, error) {
	if config.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	router.Use(CORS)

	webhook := router.Group("/api/webhook")
	{
		webhook.POST("/telegram", telegramHandler.Webhook)
	}
	test := router.Group("/api/test")
	{
		test.POST("/publish-q", telegramHandler.TestPublishQ)
	}
	return &Router{
		router,
	}, nil
}

func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
		return
	} else {
		c.AbortWithStatus(http.StatusOK)
		return
	}
}

func (r *Router) Serve(listenAddr string) error {
	srv := &http.Server{
		Addr:    listenAddr,
		Handler: r,
	}
	return srv.ListenAndServe()
}
