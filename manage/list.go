package manage

import (
	"fmt"

	"cgit.dcpri.me/deployment-agent/config"
)

// ListProjects for listing the projects
func ListProjects() {
	fmt.Println("Following are the projects with their UUIDs")
	for index, project := range config.StoredProjects {
		fmt.Printf("%2d. \"%s\", UUID: %s\n", index+1, project.Name, project.UUID)
	}
}
