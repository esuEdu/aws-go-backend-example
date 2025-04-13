package routes

import (
	"net/http"

	"github.com/esuEdu/aws-go-backend-example/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, uploadHandler *handlers.UploadHandler) {
	router.POST("/upload", uploadHandler.UploadFile)
	router.GET("/docs", uploadHandler.ListFiles)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
