package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Comment is the handler for the comment endpoint
func Comment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome Comment! =D"})
}