package manage

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// StartServer for starting the gin server on given host:port
func StartServer(host string, port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/repository/:repoName/:token", webHookHandler)
	addr := fmt.Sprintf("%s:%v", host, port)
	log.Println("Server started at", addr)
	log.Fatalln(router.Run(addr))
}
