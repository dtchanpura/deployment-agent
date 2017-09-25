package manage

import (
	"os"
	"log"
	"encoding/json"
)

func AddConfiguration(name string, repoPath string, postHook string) {
	RepositoryConfiguration.Repositories = append(RepositoryConfiguration.Repositories, Repository{Name:name, Path:repoPath, PostHookPath:postHook})

	file, err:= os.OpenFile(RepositoryConfiguration.ConfigFilePath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(RepositoryConfiguration)
	file.Close()
}