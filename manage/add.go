package manage

import (
	"fmt"

	"github.com/dtchanpura/deployment-agent/config"
)

// AddProject for adding the project
func AddProject(cfgFile string, project config.Project) error {
	// Add the project to existing configuration.
	if project.Name == "" {
		project.Name = GetRandomName()
		fmt.Printf("No name provided. Using %s for this one.\n", project.Name)
	}
	err := config.UpdateProject(cfgFile, project)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	fmt.Printf("UUID for this project is: %s\n", project.UUID)
	for i, token := range project.Tokens {
		fmt.Printf("Hash to be used for %s: %s\n", token.WhitelistedNetwork, project.GetHash(i))
	}
	return nil
}
