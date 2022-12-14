package file_repository

import (
	"context"
	"log"
	"time"

	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file"
	mongo_db "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
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
