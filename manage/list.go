package manage

import (
	"bytes"
	"fmt"

	"cgit.dcpri.me/deployment-agent/config"
)

// ListProjects for listing the projects
func ListProjects(projects []config.Project) string {
	b := new(bytes.Buffer)
	if len(projects) > 0 {
		fmt.Fprintln(b, "Following are the projects with their UUIDs")
		for index, project := range projects {
			fmt.Fprintf(b, "%2d. \"%s\", UUID: %s\n", index+1, project.Name, project.UUID)
		}
	} else {
		fmt.Fprintln(b, "No projects found. Add new one using add command.")
	}
	return b.String()
}
