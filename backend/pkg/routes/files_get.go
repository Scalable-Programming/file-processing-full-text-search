package file_upload_route

import (
	"net/http"

	controller_get_files "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/controllers/get_files"
	error_handler "github.com/Scalable-Programming/file-processing-full-text-search/backend/pkg/error_handler"
	"github.com/gin-gonic/gin"
)

func GetFilesRoute(router *gin.Engine) {
	router.GET("/files", func(c *gin.Context) {
		searchFilter, hasSearchFilter := c.GetQuery("search")
		if !hasSearchFilter {
			searchFilter = ""
		}

		files, err := controller_get_files.GetFiles(searchFilter)
		if err != nil {
			error_handler.HandleRestApiError(c, err)
			return
		}

		c.IndentedJSON(http.StatusOK, files)
	})
}
