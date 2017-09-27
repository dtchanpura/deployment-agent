package manage

import (
	"os"
	"encoding/json"
	"log"
	"fmt"
)

var RepositoryConfiguration Configuration

func init() {
	var filePath string
	var file *os.File

	if filePath = os.Getenv("CDGO_CONFIG"); filePath == "" {
		filePath = os.Getenv("HOME") + "/.local/share/cdgo.json"
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Configuration file not found. Creating an empty one.")
		file, err = os.Create(filePath)
		fileString := fmt.Sprintf("{\"token_secret\":\"%s\",\"repositories\":[]}",
			GenerateRandomString(16, 5))
		file.WriteString(fileString)
		file.Close()
		log.Println("Configuration file created. Re-run the previous command.")
		os.Exit(0)
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			log.Fatalln(err)
		}
	}

	decoder := json.NewDecoder(file)
	RepositoryConfiguration = Configuration{}
	err := decoder.Decode(&RepositoryConfiguration)
	if err != nil {
		log.Println("Error")
		log.Fatalln(err)
	}
	//fmt.Println(RepositoryConfiguration.Repositories)
	// Adding configuration path in Configuration type.
	RepositoryConfiguration.ConfigFilePath = filePath
	file.Close()

}

type Configuration struct {
	ConfigFilePath string       `json:"-"`
	TokenSecret    string       `json:"token_secret"`
	Repositories   []Repository `json:"repositories"`
}

type Repository struct {
	Name         string `json:"name"`
	Path         string `json:"path"`
	Token        string `json:"token"`
	RemoteName   string `json:"remote_name"`
	PostHookPath string `json:"post_hook_path"`
	LastUpdated  string `json:"last_updated"`
}
