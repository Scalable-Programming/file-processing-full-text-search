package main

import (
	config "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/config"
	cors "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/middleware"
	file_repository "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/repositories"
	routes "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/routes"
	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/services/elastic_search"

	"github.com/gin-gonic/gin"
)

var AppConfig config.Config

func main() {
	router := gin.Default()
	router.Use(cors.CORSMiddleware())
	router.Static("/uploads", "./uploads")

	routes.GetFilesRoute(router)
	routes.UploadFile(router)

	file_repository.CreateMongoIndex()
	elastic_search.Connect()

	router.Run("localhost:" + config.AppConfig.Port)
}
