package manage

import (
	"bytes"
	"fmt"

	"cgit.dcpri.me/deployment-agent/config"
)

// ListProjects for listing the projects
func ListProjects(projects []config.Project) string {
	b := new(bytes.Buffer)
	// b.WriteString("-w")
	fmt.Fprintln(b, "Following are the projects with their UUIDs")
	for index, project := range projects {
		fmt.Fprintf(b, "%2d. \"%s\", UUID: %s\n", index+1, project.Name, project.UUID)
	}
	return b.String()
}
