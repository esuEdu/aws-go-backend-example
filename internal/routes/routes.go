package routes

import (
	"github.com/esuEdu/aws-go-backend-example/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, uploadHandler *handlers.UploadHandler) {
	router.POST("/upload", uploadHandler.UploadFile)
	router.GET("/docs", uploadHandler.ListFiles)
}
