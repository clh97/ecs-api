package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"

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

type requestBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w requestBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	fmt.Println(w.body)
	fmt.Println(b)
	return w.ResponseWriter.Write(b)
}

func RequestStructure() gin.HandlerFunc {
	return func(c *gin.Context) {
		rbw := &requestBodyWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = rbw

		c.Next()

		var jsonData interface{}

		reqData := rbw.body.String()
		json.Unmarshal([]byte(reqData), &jsonData)

		if c.Writer.Status() >= 400 {
			currentError, ok := c.Get("currentError")

			if !ok {
				fmt.Println("Obscure stuff happened... o_o")
			}

			c.JSON(c.Writer.Status(), gin.H{"shit": jsonData})

			fmt.Println(currentError)
		}
	}
}
