package manage

import (
	"os"
	"log"
	"encoding/json"
)

func AddConfiguration(name string, repoPath string, postHook string) {
	repo := Repository{Name: name, Path: repoPath, PostHookPath: postHook, Token: GenerateRandomString(16, 4)}
	RepositoryConfiguration.Repositories = append(RepositoryConfiguration.Repositories, repo)
	log.Printf("Added a hook with name: %s. Token for auth is %s", repo.Name,
		generateHash(repo.Name+repo.Token+RepositoryConfiguration.TokenSecret))
	file, err := os.OpenFile(RepositoryConfiguration.ConfigFilePath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(RepositoryConfiguration)
	file.Close()
}
