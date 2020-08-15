package listener

import (
	"fmt"
	"net"
	"net/http"
	"strconv"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/dtchanpura/deployment-agent/constants"
)

func webHookHandler(w http.ResponseWriter, r *http.Request) {
	uuid, token, err := getCredentials(r.URL.Path)
	if err != nil {
		errorHandler(err, http.StatusBadRequest, w)
		return
	}
	clientIP := getIP(r)
	args := r.URL.Query()["arg"]
	syncFlag := false
	if s := r.URL.Query().Get("sync"); s != "" {
		syncFlag, err = strconv.ParseBool(s)
		if err != nil {
			errorHandler(err, http.StatusBadRequest, w)
			return
		}
	}

	// fmt.Println(args)
	response := generateResponse(uuid, token, clientIP, syncFlag, args...)
	response.write(w)
}
func versionHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		StatusCode: http.StatusOK,
		Ok:         true,
		Version:    constants.Version,
		BuildDate:  constants.BuildDate(),
	}
	response.write(w)
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

func getIP(r *http.Request) string {
	if realip := r.Header.Get("X-Real-Ip"); realip != "" {
		return realip
	}
	if forwardedip := r.Header.Get("X-Forwarded-For"); forwardedip != "" {
		return forwardedip
	}
	if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return host
	}
	return ""
}

func errorHandler(err error, statusCode int, w http.ResponseWriter) {
	r := Response{
		Ok:         false,
		StatusCode: statusCode,
		Message:    err.Error(),
	}
	r.write(w)
}
