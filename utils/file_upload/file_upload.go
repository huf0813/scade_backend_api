package file_upload

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func NewFileUpload(base string, fileHeader *multipart.FileHeader) (string, error) {
	randomUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	fileHeader.Filename = fmt.Sprintf("%s_%s", randomUUID, fileHeader.Filename)

	src, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(src)

	destination := fmt.Sprintf("%s/%s", base, fileHeader.Filename)
	dst, err := os.Create(destination)
	if err != nil {
		return "", err
	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dst)

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return destination, nil
}
