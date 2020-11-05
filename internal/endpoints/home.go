package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home is the handler for the home endpoint
func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome home! =D"})
}
