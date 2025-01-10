package server

import (
	db "allchat-message-service/_config"
	"allchat-message-service/internal/router"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	color.Green("Server starting...")
	_ = godotenv.Load()

	// database
	resource, err := db.CreateResource()
	if err != nil {
		color.Red("Connection database failure, Please check connection")
		color.Cyan(err.Error())
		logrus.Error(err)
	}
	defer resource.Close()

	// route
	r := gin.Default()
	r.Use(CORS)

	router.Router(r)
	router.Webhook(r, resource)

	// server
	port := os.Getenv("PORT")
	fmt.Printf("Server listening on port %s\n", port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("listen: %s\n", err)
	}
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
