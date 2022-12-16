package controller_get_files

import (
	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file"
	file_repository "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/repositories"
	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/services/elastic_search"
)

func GetFiles(search string) ([]file.File, error) {
	fileIds := []string{}
	if search != "" {
		esFileIds, err := elastic_search.Search(search)
		fileIds = esFileIds

		if err != nil {
			return nil, err
		}

		if len(fileIds) == 0 {
			return []file.File{}, nil
		}
	}

	return file_repository.GetFiles(fileIds)
}
