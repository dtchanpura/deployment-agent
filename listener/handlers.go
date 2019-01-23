package listener

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/gin-gonic/gin"
)

func webHookHandler(c *gin.Context) {
	uuid := c.Param("uuid")
	token := c.Param("token")
	clientIP := c.ClientIP()
	args := c.QueryArray("arg")
	syncFlag := strings.EqualFold(c.Query("sync"), "true")

	// fmt.Println(args)
	response := generateResponse(uuid, token, clientIP, syncFlag, args...)
	c.Status(response.StatusCode)
	// c.JSON(response.StatusCode, response)
	c.JSON(http.StatusOK, response)
}

func generateResponse(uuid, token, clientIP string, syncFlag bool, args ...string) Response {
	response := Response{StatusCode: http.StatusOK, Ok: false, Message: "execution queued"}
	//fmt.Println(reponame, token)
	// repo := findProject(uuid)
	result := validateToken(uuid, token, clientIP)
	if result {
		// c.Writer.Write([]byte("Token Valid\n"))
		project, err := config.FindProjectWithUUID(uuid)
		if err != nil {
			fmt.Println(err) // this will never occur as
		}
		if !syncFlag {
			go executeHooks(project, args...)
		} else {
			executeHooks(project, args...)
			response.Message = "execution completed"
		}
		response.Ok = true
	} else {
		response.StatusCode = http.StatusUnauthorized
		response.Message = "Unauthorized"
		response.Ok = false
	}
	return response
}
