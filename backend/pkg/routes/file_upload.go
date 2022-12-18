package file_upload_route

import (
	"net/http"

	controller_upload_file "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/controllers/upload_file"
	error_handler "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/error_handler"

	"github.com/gin-gonic/gin"
)

func UploadFile(router *gin.Engine) {
	router.POST("/file", func(c *gin.Context) {
		file, err := c.FormFile("file")

		if err != nil {
			error_handler.HandleRestApiError(c, err)
			return
		}

		newUploadedFile, err := controller_upload_file.UploadFile(file)

		if err != nil {
			error_handler.HandleRestApiError(c, err)
			return
		}

		c.IndentedJSON(http.StatusCreated, newUploadedFile)
	})
}
