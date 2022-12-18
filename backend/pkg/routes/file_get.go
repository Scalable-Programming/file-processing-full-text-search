package file_upload_route

import (
	"net/http"

	controller_get_file "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/controllers/get_file"
	error_handler "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/error_handler"
	"github.com/gin-gonic/gin"
)

func GetFileRoute(router *gin.Engine) {
	router.GET("/file/:file_id", func(c *gin.Context) {
		fileId := c.Params.ByName("file_id")

		files, err := controller_get_file.GetFileById(fileId)
		if err != nil {
			error_handler.HandleRestApiError(c, err)
			return
		}

		c.IndentedJSON(http.StatusOK, files)
	})
}
