package controller_get_files

import (
	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file"
	file_repository "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/repositories"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFiles() ([]file.File, error) {
	return file_repository.GetFiles(bson.D{})
}
