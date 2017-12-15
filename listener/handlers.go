package listener

import (
	"fmt"
	"net/http"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/gin-gonic/gin"
)

func webHookHandler(c *gin.Context) {
	uuid := c.Param("uuid")
	token := c.Param("token")
	clientIP := c.ClientIP()
	response := generateResponse(uuid, token, clientIP)
	c.Status(response.StatusCode)
	// c.JSON(response.StatusCode, response)
	c.JSON(http.StatusOK, response)
}

func generateResponse(uuid, token, clientIP string) Response {
	response := Response{StatusCode: http.StatusOK, Ok: false, Message: ""}
	//fmt.Println(reponame, token)
	// repo := findProject(uuid)
	result := validateToken(uuid, token, clientIP)
	if result {
		var isUpToDate bool
		// c.Writer.Write([]byte("Token Valid\n"))
		repo, err := config.FindProjectWithUUID(uuid)
		if err != nil {
			fmt.Println(err) // this will never occur as
		}
		if !isUpToDate {
			go executeHooks(repo)
			response.Ok = true
		}
	} else {
		response.StatusCode = http.StatusUnauthorized
		response.Message = "Unauthorized"
		response.Ok = false
	}
	return response
}
