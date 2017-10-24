package manage

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	git "gopkg.in/src-d/go-git.v4"
)

func webHookHandler(c *gin.Context) {
	repoName := c.Param("repoName")
	token := c.Param("token")
	clientIP := c.ClientIP()
	response := generateResponse(repoName, token, clientIP)
	c.Status(response.StatusCode)
	// c.JSON(response.StatusCode, response)
	c.JSON(http.StatusOK, response)
}

func generateResponse(repoName, token, clientIP string) Response {
	response := Response{StatusCode: 200, Ok: false, Message: ""}
	//fmt.Println(reponame, token)
	result := validateToken(repoName, token, clientIP)
	if result {
		var isUpToDate bool
		// c.Writer.Write([]byte("Token Valid\n"))
		repo := findRepository(repoName)
		err := PullRepository(repo.Path, repo.RemoteName)
		if err != nil {
			if err == git.NoErrAlreadyUpToDate {
				response.StatusCode = 200
				response.Message = fmt.Sprintf("Already up-to-date.")
				response.Ok = true
				isUpToDate = true
			} else {
				response.Ok = false
				response.Message = fmt.Sprintf("Error in Pulling: %v", err)
				response.StatusCode = 500
				// log.Printf("Error in Pulling: %v\n", err)
			}
		}
		if !isUpToDate {
			go ExecuteHook(repo.PostHookPath)
			response.Ok = true
		}
	} else {
		response.StatusCode = 403
		response.Message = "Invalid Token"
		response.Ok = false
	}
	return response
}
