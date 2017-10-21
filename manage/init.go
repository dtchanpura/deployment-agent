package manage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

// RepositoryConfiguration Global Configuration Object of type Configuration
var RepositoryConfiguration Configuration

func init() {
	var filePath string
	var file *os.File

	// Configuration file stored at any of following paths
	// * Path in environment variable CDGO_CONFIG
	// * Path in $HOME/.config/cd-go/config.json
	if filePath = os.Getenv("CDGO_CONFIG"); filePath == "" {
		filePath = os.Getenv("HOME") + "/.config/cd-go/config.json"
	}
	folderPathStrings := strings.Split(filePath, string(os.PathSeparator))
	folderPath := strings.Join(folderPathStrings[:len(folderPathStrings)-1], string(os.PathSeparator))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Configuration file not found. Creating an empty one.")
		_ = os.Mkdir(folderPath, 0755)
		file, err = os.Create(filePath)
		if err != nil {
			log.Println("Error while creating a file.")
			log.Fatal(err)
		}
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
