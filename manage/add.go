package manage

import (
	"fmt"

	"cgit.dcpri.me/deployment-agent/config"
)

// AddProject for adding the project
func AddProject(cfgFile string, project config.Project) error {
	// Add the project to existing configuration.
	err := config.UpdateProject(cfgFile, project)
	if err != nil {
		// fmt.Println(err)
		return err
	}
	for i, token := range project.Tokens {
		fmt.Printf("Hash to be used for %s: %s\n", token.WhitelistedNetwork, project.GetHash(i))
	}
	return nil
}
