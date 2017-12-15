package manage

import (
	"fmt"
	"testing"

	"github.com/dtchanpura/deployment-agent/config"
)

func TestListProjects(t *testing.T) {
	name := GetRandomName()
	projects := []config.Project{
		{
			Name: name,
			UUID: "d6e0c569-0359-4f6c-bdeb-21ac53f42b00",
		},
	}

	output := fmt.Sprintf(
		`Following are the projects with their UUIDs
 1. "%s", UUID: d6e0c569-0359-4f6c-bdeb-21ac53f42b00
`,
		projects[0].Name)
	if ListProjects(projects) != output {
		fmt.Println(ListProjects(projects))
		fmt.Println(output)
		t.Errorf("strings don't match.")
	}

	noProjectsOutput := "No projects found. Add new one using add command.\n"
	if ListProjects([]config.Project{}) != noProjectsOutput {
		fmt.Println(ListProjects([]config.Project{}))
	}
	// fmt.Println(ListProjects(projects))
}
