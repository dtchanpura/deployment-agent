package listener

import (
	"fmt"

	"cgit.dcpri.me/deployment-agent/common"
	"cgit.dcpri.me/deployment-agent/config"
)

func validateToken(projectUUID, token, clientIP string) bool {
	project, err := common.FindProject(projectUUID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return project.ValidateToken(clientIP, token)
}

func executeHooks(project config.Project) {
	fmt.Println("hello.")
}
