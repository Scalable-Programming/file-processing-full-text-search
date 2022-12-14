package main

import (
	config "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/config"
	file_repository "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/repositories"
	routes "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/routes"

	"github.com/gin-gonic/gin"
)

var AppConfig config.Config

func main() {
	router := gin.Default()

	routes.GetFilesRoute(router)
	routes.UploadFile(router)

	file_repository.CreateMongoIndex()

	router.Run("localhost:" + config.AppConfig.Port)
}
