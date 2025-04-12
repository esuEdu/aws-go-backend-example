package main

import (
	"log"

	"github.com/esuEdu/aws-go-backend-example/config"
	"github.com/esuEdu/aws-go-backend-example/internal/handlers"
	"github.com/esuEdu/aws-go-backend-example/internal/repo"
	"github.com/esuEdu/aws-go-backend-example/internal/routes"
	"github.com/esuEdu/aws-go-backend-example/internal/storage"
	"github.com/gin-gonic/gin"
)

func main() {

	// s3 Setup
	cfg := config.LoadS3Config()
	s3Client := config.InitS3Client(&cfg)

	// postgre setup

	db := config.InitDB()

	uploader := storage.NewS3Uploader(s3Client, cfg.Bucket)
	docRepo := repo.NewDocRepo(db)
	uploadHandler := handlers.NewUploadHandler(uploader, docRepo)

	// Create a new Gin router
	r := gin.Default()
	routes.SetupRoutes(r, uploadHandler)

	log.Println("Server running at port :8080")
	r.Run(":8080")
}
