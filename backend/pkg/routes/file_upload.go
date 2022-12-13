package file_upload_route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(router *gin.Engine) {
	router.POST("/file", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, nil)
	})
}
