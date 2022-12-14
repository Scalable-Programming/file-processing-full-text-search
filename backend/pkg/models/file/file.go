package file

import (
	"time"

	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file_status"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	Id          primitive.ObjectID `bjson:"_id"`
	CreatedAt   time.Time          `bjson:"createdAt"`
	ContentType string             `bjson:"contentType"`
	Name        string             `bjson:"name"`
	Size        int                `bjson:"size"`
	Status      int                `bjson:"status"`
}

func NewFile(contentType string, name string, size int) *File {
	newFile := new(File)

	newFile.Id = primitive.NewObjectID()
	newFile.ContentType = contentType
	newFile.Name = name
	newFile.Size = size
	newFile.Status = int(file_status.Pending)
	newFile.CreatedAt = time.Now()

	return newFile
}
