package pkg

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	viper "github.com/spf13/viper"
)

// InitServer initializes http listening and sets up endpoints
func InitServer() {
	gin.ForceConsoleColor()

	router := InitRouter()

	addr := getAddr()

	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func getAddr() string {
	return fmt.Sprintf("%s:%d", viper.Get("server.addr"), viper.Get("server.port"))
}
