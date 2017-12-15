package manage

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dtchanpura/deployment-agent/config"
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

	project = config.Project{
		UUID: "test-uuid-new0",
	}
	project.Tokens = append([]config.TokenDetail{}, config.TokenDetail{
		Token:              "zZZzZZzZZz",
		WhitelistedNetwork: "0.0.0.0/0",
	})

	AddProject(cfgFile, project)
	os.Remove(cfgFile)
	// testConfig.ProjectConfigs
	// os.Open(cfgFile)
}
