package file

import (
	"time"

	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file_status"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	Id          primitive.ObjectID `json:"_id"`
	CreatedAt   time.Time          `json:"createdAt"`
	ContentType string             `json:"contentType"`
	Name        string             `json:"name"`
	Size        int                `json:"size"`
	Status      int                `json:"status"`
}

func NewFile(contentType string, name string, size int) *File {
	newFile := new(File)

	newFile.ContentType = contentType
	newFile.Name = name
	newFile.Size = size
	newFile.Status = int(file_status.Pending)
	newFile.CreatedAt = time.Now()

	return newFile
}
