package file_upload

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func NewFileUpload(base string,
	fileHeader *multipart.FileHeader) (string, error) {
	unique, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	fileHeader.Filename = fmt.Sprintf("%s_%s", unique, fileHeader.Filename)

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

	path := fmt.Sprintf("%s/%s", base, fileHeader.Filename)
	dst, err := os.Create(path)
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

	return fileHeader.Filename, nil
}
