package listener

import (
	"fmt"

	"cgit.dcpri.me/deployment-agent/config"
)

func validateToken(projectUUID, token, clientIP string) bool {
	project := findProject(projectUUID)
	return project.ValidateToken(clientIP, token)
}

func findProject(projectUUID string) config.Project {
	return *config.NewProject()
}

func executeHooks(project config.Project) {
	fmt.Println("hello.")
}
