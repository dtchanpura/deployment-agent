package manage

import (
	"encoding/json"
	"log"
	"os"
)

// AddConfiguration function for adding a new configuration in RepositoryConfiguration
// and also writing the same in Configuration File.
func AddConfiguration(name string, repoPath string, postHook string, whitelistedIPs []string) {
	repo := Repository{Name: name, Path: repoPath, PostHookPath: postHook, Token: GenerateRandomString(16, 4), WhiteListedIPs: whitelistedIPs}
	RepositoryConfiguration.Repositories = append(RepositoryConfiguration.Repositories, repo)
	log.Printf("Added a hook with name: %s. Token for auth is %s", repo.Name,
		generateHash(repo.Name+repo.Token+RepositoryConfiguration.TokenSecret))
	file, err := os.OpenFile(RepositoryConfiguration.ConfigFilePath, os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println("Error Opening File.")
		log.Fatalln(err)
	}
	encoder := json.NewEncoder(file)
	err = encoder.Encode(RepositoryConfiguration)
	if err != nil {
		log.Println("Error Encoding JSON.")
		log.Fatalln(err)
	}
	file.Close()
}
