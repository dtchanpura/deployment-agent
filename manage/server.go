package manage

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)
func StartServer(port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/repository/:reponame/:token", func(c *gin.Context) {
		reponame := c.Param("reponame")
		token := c.Param("token")
		c.Status(http.StatusOK)
		fmt.Println(reponame, token)

	})
	addr := fmt.Sprintf(":%v", port)
	router.Run(addr)
}