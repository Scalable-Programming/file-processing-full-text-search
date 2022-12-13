package file_repository

import (
	"context"
	"log"
	"time"

	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/models/file"
	mongo_db "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

var fileCollection = mongo_db.Db.Collection("file")

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

	for cursor.Next(context.Background()) {
		result := struct {
			Foo string
			Bar int32
		}{}
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		file := file.File{}

		decode_err := cursor.Decode(&file)
		if decode_err != nil {
			log.Fatal(decode_err)
			return nil, err
		}

		files = append(files, file)

	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return files, nil
}
