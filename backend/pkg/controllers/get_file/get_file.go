package controller_get_files

import (
	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file"
	file_repository "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/repositories"
	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/services/elastic_search"
)

func GetFileById(fileId string) (*file.File, error) {
	text, err := elastic_search.GetFileText(fileId)

	if err != nil {
		return nil, err
	}

	file, err := file_repository.GetFile(fileId)

	if err != nil {
		return nil, err
	}

	file.Text = *text

	return file, nil
}
