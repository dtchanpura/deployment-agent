package manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/src-d/go-git.v4"
	"log"
	"net/http"
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
			var isUpToDate bool
			c.Writer.Write([]byte("Token Valid\n"))
			repo := findRepository(repoName)
			err := PullRepository(repo.Path, repo.RemoteName)
			if err != nil {
				log.Printf("Error in Pulling: %v\n", err)
				if err == git.NoErrAlreadyUpToDate {
					isUpToDate = true
				}
			}
			if !isUpToDate {
				go ExecuteHook(repo.PostHookPath)
			} else {
				c.Writer.Write([]byte("Repository already up-to-date"))
			}
		} else {
			c.Writer.Write([]byte("Invalid"))
		}

	})
	addr := fmt.Sprintf("%s:%v", host, port)
	log.Println("Server started at", addr)
	log.Fatalln(router.Run(addr))
}
