package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ratneshgujarathi/url-shortner-redis/api/constants"
	"github.com/ratneshgujarathi/url-shortner-redis/api/handlers"
)

func SetupRouter(router *gin.Engine) {
	router.POST(constants.SHORT_URL_PATH, handlers.ShortenUrl)
}
