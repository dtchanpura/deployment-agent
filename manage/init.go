package manage

import "os"

// RepositoryConfiguration Global Configuration Object of type Configuration
var RepositoryConfiguration Configuration

func init() {
	var filePath string
	if filePath = os.Getenv("CDGO_CONFIG"); filePath == "" {
		filePath = os.Getenv("HOME") + "/.config/cd-go/config.json"
	}
	initializeConfigFile(filePath)
}
