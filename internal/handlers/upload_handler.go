package handlers

import (
	"net/http"

	"github.com/esuEdu/aws-go-backend-example/internal/models"
	"github.com/esuEdu/aws-go-backend-example/internal/repo"
	"github.com/esuEdu/aws-go-backend-example/internal/storage"
	"github.com/gin-gonic/gin"
)

type UploadHandler struct {
	uploader storage.S3Uploader
	docRepo  repo.DocRepo
}

func NewUploadHandler(uploader storage.S3Uploader, docRepo repo.DocRepo) *UploadHandler {
	return &UploadHandler{uploader: uploader, docRepo: docRepo}
}

func (h *UploadHandler) UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}

	fileID, url, err := h.uploader.UploadFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	doc := &models.Doc{
		Name: fileID,
		Url:  url,
	}

	if err := h.docRepo.Save(doc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save metadata"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "upload successful", "doc": doc})
}

func (h *UploadHandler) ListFiles(c *gin.Context) {
	docs, err := h.docRepo.List(&models.Doc{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list documents"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"documents": docs})
}
