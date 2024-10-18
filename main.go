package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ratneshgujarathi/url-shortner-redis/api/constants"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":", port))
}

func setupRouter(router *gin.Engine) {
	router.POST(constants.SHORT_URL_PATH)
}
