package storage_upload

import (
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

type ValidationContentTypeError struct{}

func (m *ValidationContentTypeError) Error() string {
	return "Invalid content type"
}

func UploadFileToLocalStorage(fileHeader multipart.FileHeader) (*string, error) {
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return nil, err
	}

	uuid := uuid.New()
	localFilePath := "./uploads/" + uuid.String()

	dst, err := os.Create(localFilePath)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	openedFile, err := fileHeader.Open()
	defer openedFile.Close()
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(dst, openedFile)
	if err != nil {
		return nil, err
	}

	fullPath := localFilePath + "/" + fileHeader.Filename

	return &fullPath, nil

}

func ValidateFileType(fileHeader multipart.FileHeader) (*string, error) {
	validContentType := "application/pdf"

	if fileHeader.Header.Get("Content-Type") != validContentType {
		return nil, &ValidationContentTypeError{}
	}

	return &validContentType, nil
}
