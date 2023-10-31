package uploadverse

import (
	"mime/multipart"
	"time"
)

type File struct {
	FileId    string    `json:"fileId"`
	FileSize  int       `json:"fileSize"`
	Timestamp time.Time `json:"timestamp"`
}

type FileService interface {
	UploadFile(file *multipart.FileHeader) (*File, error)
	GetFile(fileId string) ([]byte, error)
}
