package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	router = gin.Default()
)

func init() {
	_ = godotenv.Load()
}

func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
