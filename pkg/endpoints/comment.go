package endpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PublicComment is the handler for the anon comment endpoint
func PublicComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome Comment! =D"})
}

// PrivateComment is the handler for the anon comment endpoint
func PrivateComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Welcome Comment! =D"})
}
