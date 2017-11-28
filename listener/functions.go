package listener

import (
	"fmt"

	"cgit.dcpri.me/deployment-agent/config"
)

func validateToken(projectUUID, token, clientIP string) bool {
	project, err := config.FindProjectWithUUID(projectUUID)
	if err != nil {
		// fmt.Println(err)
		return false
	}
	return project.ValidateToken(clientIP, token)
}

func executeHooks(project config.Project) {
	fmt.Println("hello.")
}
