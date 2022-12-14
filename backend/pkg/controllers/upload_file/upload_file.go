package controller_upload_file

import (
	"mime/multipart"

	file "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file"
	file_repository "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/repositories"
	storage_upload "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/services/storage_upload"
)

func UploadFile(uploadedFile *multipart.FileHeader) (*file.File, error) {
	content_type, err := storage_upload.ValidateFileType(*uploadedFile)
	if err != nil {
		return nil, err
	}

	saved_file_location, err := storage_upload.UploadFileToLocalStorage(*uploadedFile)
	if err != nil {
		return nil, err
	}

	newFile, err := file_repository.InsertNewFile(*content_type, uploadedFile.Filename, int(uploadedFile.Size), *saved_file_location)

	if err != nil {
		return nil, err
	}

	return newFile, nil
}
