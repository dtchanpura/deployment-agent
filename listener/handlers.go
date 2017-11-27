package listener

import (
	"net/http"

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
		repo := findProject(uuid)
		if !isUpToDate {
			go executeHooks(repo)
			response.Ok = true
		}
	} else {
		response.StatusCode = 403
		response.Message = "Invalid Token"
		response.Ok = false
	}
	return response
}
