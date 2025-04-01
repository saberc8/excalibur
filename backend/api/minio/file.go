package minio

import (
	"aquila/global"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"net/http"
)

type Minio struct{}

func (m Minio) GetBuckets(c *gin.Context) {
	buckets, err := global.AquilaMinio.ListBuckets(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"buckets": buckets})
}

func (m Minio) UploadFile(c *gin.Context) {
	bucketName := c.Param("bucket")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	objectName := header.Filename
	contentType := header.Header.Get("Content-Type")

	_, err = global.AquilaMinio.PutObject(context.Background(), bucketName, objectName, file, header.Size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully", "file": objectName})
}
