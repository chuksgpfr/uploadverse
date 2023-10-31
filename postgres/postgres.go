package postgres

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"time"

	"github.com/chuksgpfr/uploadverse"
	"github.com/chuksgpfr/uploadverse/ipfs"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FileService struct {
	DB          *gorm.DB
	IpfsService ipfs.IpfsService
}

func createDBClient(postgresDSN string) *gorm.DB {
	fmt.Println("DSN ", postgresDSN)
	db, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})

	if err != nil {
		// panic(err)
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}

func NewDBClient(postgresDSN string) *gorm.DB {
	client := createDBClient(postgresDSN)
	client.AutoMigrate(uploadverse.File{})
	return client
}

func (f *FileService) UploadFile(file *multipart.FileHeader) (*uploadverse.File, error) {
	//upload file to IPFS
	fileId, err := f.IpfsService.AddFile(file)
	if err != nil {
		return nil, err
	}

	fileSize := file.Size

	now := time.Now()

	var result *uploadverse.File

	findFilter := &uploadverse.File{
		FileId: fileId,
	}

	err = f.DB.Take(&result, findFilter).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	response := &uploadverse.File{
		FileId:    fileId,
		FileSize:  int(fileSize),
		Timestamp: now,
	}

	if err == gorm.ErrRecordNotFound {
		err = f.DB.Create(response).Error
		if err != nil {
			return nil, err
		}
	}

	return response, nil

}

func (f *FileService) GetFile(fileId string) ([]byte, error) {
	var result *uploadverse.File
	findFilter := &uploadverse.File{
		FileId: fileId,
	}

	err := f.DB.Take(&result, findFilter).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Record not found")
	}

	fileByte, err := f.IpfsService.GetFile(fileId)

	if err != nil {
		return nil, err
	}

	return fileByte, nil
}
