package manage

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"log"
)

func StartServer(host string, port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/repository/:reponame/:token", func(c *gin.Context) {
		repoName := c.Param("reponame")
		token := c.Param("token")
		c.Status(http.StatusOK)
		//fmt.Println(reponame, token)
		result := validateToken(repoName, token)
		if result {
			c.Writer.Write([]byte("Valid"))
		} else {
			c.Writer.Write([]byte("Invalid"))
		}

	})
	addr := fmt.Sprintf("%s:%v", host, port)
	log.Fatalln(router.Run(addr))
}
