package ipfs

import (
	"fmt"
	"io"
	"mime/multipart"

	shell "github.com/ipfs/go-ipfs-api"
)

type IpfsService struct {
	Sh *shell.Shell
}

func (i *IpfsService) AddFile(file *multipart.FileHeader) (string, error) {
	utilFile, err := file.Open()
	defer utilFile.Close()

	if err != nil {
		return "", nil
	}

	fileId, err := i.Sh.Add(utilFile)

	return fileId, nil

}

func (i *IpfsService) GetFile(fileId string) ([]byte, error) {
	reader, err := i.Sh.Cat(fmt.Sprintf("/ipfs/%s", fileId))

	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(reader)

	return bytes, err

}
