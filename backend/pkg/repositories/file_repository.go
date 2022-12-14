package file_repository

import (
	"context"
	"log"
	"time"

	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file"
	file_status "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file_status"
	mongo_db "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var fileCollection = mongo_db.Db.Collection("file")

func CreateMongoIndex() {
	var indexModel = mongo.IndexModel{
		Keys: bson.D{
			{Key: "createdAt", Value: 1},
			{Key: "status", Value: 1},
		},
	}
	_, err := fileCollection.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		panic(err)
	}
}

func GetFiles(filter bson.D) ([]file.File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := fileCollection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.Background())

	files := []file.File{}

	err = cursor.All(ctx, &files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func InsertNewFile(contentType string, name string, size int, filePath string) (*file.File, error) {
	newFile := file.NewFile(contentType, name, size, filePath)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := fileCollection.InsertOne(ctx, newFile)

	if err != nil {
		return nil, err
	}

	return newFile, nil
}

func UpdateFile(id primitive.ObjectID, updates bson.M) (file.File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updates["lastUpdatedAt"] = time.Now()

	result := fileCollection.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": updates})

	updatedFile := file.File{}
	decodeErr := result.Decode(&updatedFile)

	return updatedFile, decodeErr
}

func UpdateStatus(id primitive.ObjectID, status file_status.FileStatus) {
	UpdateFile(id, bson.M{"status": status})
}
