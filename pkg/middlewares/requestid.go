package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID acts as middleware, setting up the RID header as some UUID
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		uuid, err := uuid.NewUUID()

		if err != nil {
			return
		}

		c.Header("RID", uuid.String())
		c.Next()
	}
}
