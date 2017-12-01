package manage

import (
	"io/ioutil"
	"testing"

	"cgit.dcpri.me/deployment-agent/config"
)

var (
	cfgFile    string
	testConfig config.Configuration
	project    config.Project
)

func TestAddProject(t *testing.T) {
	cfgFile = "./test-config.yaml"
	ioutil.WriteFile(cfgFile, []byte{}, 0644)

	project = config.Project{
		Name: "test-project",
		UUID: "test-uuid",
	}

	AddProject(cfgFile, project)
	// testConfig.ProjectConfigs
	// os.Open(cfgFile)
}
