package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/api/shorten", shortenURL)
	server.GET("api/shorten/:shortUrl", getShortURL)
	server.PUT("api/shorten/:shortUrl", updateShortURL)
	server.DELETE("api/shorten/:shortUrl", deleteShortURL)

	server.GET("/:shortUrl", redirectURL)
}
