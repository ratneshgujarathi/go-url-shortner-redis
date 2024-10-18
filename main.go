package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ratneshgujarathi/url-shortner-redis/api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	routes.SetupRouter((router))
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}
	serverAddr := ":" + port
	log.Fatal(router.Run(serverAddr))
}
