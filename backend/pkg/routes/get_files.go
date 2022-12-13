package file_upload_route

import (
	"net/http"

	"github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func GetFilesRoute(router *gin.Engine) {
	router.GET("/files", func(c *gin.Context) {
		files, err := controllers.GetFiles()

		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, files)
	})
}
