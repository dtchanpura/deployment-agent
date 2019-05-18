package listener

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// StartListener for starting the gin server on given host:port
func StartListener(host string, port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/reload/:uuid/:token", webHookHandler)
	router.POST("/reload/:uuid/:token", webHookHandler)
	router.GET("/version", versionHandler)
	addr := fmt.Sprintf("%s:%v", host, port)
	log.Println("Server started at", addr)
	log.Fatalln(router.Run(addr))
}
