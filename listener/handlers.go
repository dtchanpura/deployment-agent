package listener

import (
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
	l := logger.Debug().Fields(map[string]interface{}{"uuid": uuid, "token": token, "ip": clientIP, "valid": result})
	if result {
		project, err := config.FindProjectWithUUID(uuid)
		if err != nil {
			logger.Error().Err(err).Send() // this will never occur as
			return response
		}
		if !syncFlag {
			go executeHooks(project, args...)
		} else {
			executeHooks(project, args...)
			response.Message = "execution completed"
		}
		dateVec.WithLabelValues(project.Name, uuid, clientIP, response.Message).SetToCurrentTime()
		l.Str("name", project.Name).Msg(response.Message)
		response.Ok = true
		return response
	}

	response.StatusCode = http.StatusUnauthorized
	response.Message = "Unauthorized"
	response.Ok = false
	l.Msg(response.Message)
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
	logger.Error().Err(err).Send()
	r.write(w)
}
