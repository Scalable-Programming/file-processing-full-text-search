package storage_upload

import (
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

type ValidationContentTypeError struct{}

func (m *ValidationContentTypeError) Error() string {
	return "Invalid content type"
}

func UploadFileToLocalStorage(fileHeader multipart.FileHeader) (*string, error) {
	uuid := uuid.New()
	folderName := "./uploads/" + uuid.String()
	err := os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		return nil, err
	}

	localFilePath := folderName + "/" + fileHeader.Filename

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

	return &localFilePath, nil

}

func ValidateFileType(fileHeader multipart.FileHeader) (*string, error) {
	validContentType := "application/pdf"

	if fileHeader.Header.Get("Content-Type") != validContentType {
		return nil, &ValidationContentTypeError{}
	}

	return &validContentType, nil
}

func DeleteLocalStorage(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Println("Error removing file", err)
	}
}
