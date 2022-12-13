package main

import (
	routes "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.GetFilesRoute(router)
	routes.GetFilesRoute(router)

	router.Run("localhost:8080")
}
