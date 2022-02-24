package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kis9a/lambda-sls/models"
	"go.uber.org/zap"
)

func ImageUpload() func(*gin.Context) {
	return func(c *gin.Context) {
		fs, file, err := c.Request.FormFile("file")
		if err != nil {
			zap.S().Error(err)
			return
		}
		uploader := models.NewImageUploader()
		out, err := uploader.Upload(file, fs)
		zap.S().Error(err)
		c.JSON(200, gin.H{"uri": out.Location})
		return
	}
}
