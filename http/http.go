package http

import (
	"github.com/chuksgpfr/uploadverse/postgres"
	"github.com/gin-gonic/gin"
)

func NewServer(fileService postgres.FileService) (*gin.Engine, error) {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// router.MaxMultipartMemory = 800 << 20
	var h Handler
	h.FileService = &fileService

	router.POST("/upload", h.UploadFileHandler)
	router.GET("/file/:fileId", h.GetFileHandler)

	return router, nil
}
