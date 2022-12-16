package controller_upload_file

import (
	"mime/multipart"

	file "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file"
	file_status "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file_status"
	file_repository "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/repositories"
	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/services/elastic_search"
	gs "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/services/ghost_script"
	pdf_reader "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/services/pdf_reader"
	storage_upload "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/services/storage_upload"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	go HandleFileProcessing(newFile.Id, *saved_file_location)

	return newFile, nil
}

func TriggerPdfParsing(path string) (string, error) {
	text, err := pdf_reader.ReadPdf(path)
	return text, err
}

func HandleFileProcessing(id primitive.ObjectID, path string) {
	file_repository.UpdateStatus(id, file_status.InProcess)
	text, err := TriggerPdfParsing(path)

	if err != nil {
		file_repository.UpdateStatus(id, file_status.Pending)
		return
	}

	elastic_search.IndexFullFileText(id.Hex(), &text)
	thumbnailPath, err := gs.GenerateThumbnail(path)

	if err != nil {
		file_repository.UpdateStatus(id, file_status.Pending)
		return
	}

	file_repository.UpdateThumbnail(id, thumbnailPath)
	file_repository.UpdateStatus(id, file_status.Processed)

	// splitPath := strings.SplitAfter(path, "/")
	// storage_upload.DeleteLocalStorage(strings.Join(splitPath[:len(splitPath)-1], ""))

}
