package file

import (
	"time"

	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file_status"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	Id            primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	CreatedAt     time.Time          `bson:"createdAt" json:"createdAt"`
	ContentType   string             `bson:"contentType" json:"contentType"`
	FilePath      string             `bson:"filePath" json:"filePath"`
	LastUpdatedAt time.Time          `bson:"lastUpdatedAt" json:"lastUpdatedAt"`
	Name          string             `bson:"name" json:"name"`
	Size          int                `bson:"size" json:"size"`
	Status        int                `bson:"status" json:"status"`
	Thumbnail     string             `bson:"thumbnail" json:"thumbnail,omitempty"`
}

func NewFile(contentType string, name string, size int, filePath string) *File {
	newFile := new(File)

	newFile.Id = primitive.NewObjectID()
	newFile.ContentType = contentType
	newFile.Name = name
	newFile.Size = size
	newFile.Status = int(file_status.Pending)
	newFile.CreatedAt = time.Now()
	newFile.LastUpdatedAt = time.Now()
	newFile.FilePath = filePath

	return newFile
}
