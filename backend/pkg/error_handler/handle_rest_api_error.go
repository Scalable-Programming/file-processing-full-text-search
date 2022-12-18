package handle_rest_api_error

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRestApiError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
