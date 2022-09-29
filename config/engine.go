package config

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func RunMode() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed load env file.")
	}

	GinMode := os.Getenv("GIN_MODE")
	SetMode(GinMode)

}

func SetMode(mode string) {
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		panic("mode unavailable. (debug, release, test)")
	}
}
