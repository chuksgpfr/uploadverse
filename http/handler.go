package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chuksgpfr/uploadverse"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	FileService uploadverse.FileService
}

func (h *Handler) UploadFileHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not upload file")
		return
	}

	uploadedFile, err := h.FileService.UploadFile(file)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not upload file")
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
		"data":    &uploadedFile,
	})
	return
}

func (h *Handler) GetFileHandler(ctx *gin.Context) {
	fileId, ok := ctx.Params.Get("fileId")

	if !ok {
		ctx.JSON(http.StatusBadRequest, "No file Id was passed")
		return
	}

	fileByte, err := h.FileService.GetFile(fileId)

	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, "Could not get file")
		return
	}

	contentType := http.DetectContentType(fileByte)

	ctx.Data(http.StatusOK, fmt.Sprintf("%s; charset=utf-8", contentType), fileByte)
	return
}
