package manage

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/src-d/go-git.v4"
)

// PullRepository for pulling a repository with repoPath and remoteName
func PullRepository(repoPath string, remoteName string) error {
	if remoteName == "" {
		remoteName = "origin"
	}
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		return err
	}
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	err = w.Pull(&git.PullOptions{RemoteName: remoteName})
	if err != nil {
		return err
	}
	return nil
}

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

// ExecuteHook for executing hook for given hookPath
// It should be a executable file with shebang (`#!`) in first line.
// For Example: `#!/bin/bash`
func ExecuteHook(hookPath string) {
	cmd := exec.Command(hookPath)
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

// GenerateRandomString for generating a random string of specified length and
// strength between 1 to 5.
func GenerateRandomString(length int, strength int) string {
	if strength > 5 {
		strength = 5
	}
	if strength < 1 {
		strength = 1
	}
	var tempString string
	for i := 0; i < strength; i++ {
		tempString += TokenConstants[i]
	}
	bs := make([]byte, length)
	for i := range bs {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(tempString))))
		bs[i] = tempString[randomIndex.Int64()]
	}
	return string(bs)
}

func initializeConfigFile(filePath string) {
	var file *os.File

	// Configuration file stored at any of following paths
	// * Path in environment variable CDGO_CONFIG
	// * Path in $HOME/.config/cd-go/config.json

	folderPathStrings := strings.Split(filePath, string(os.PathSeparator))
	folderPath := strings.Join(folderPathStrings[:len(folderPathStrings)-1], string(os.PathSeparator))

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Println("Configuration file not found. Creating an empty one.")
		_ = os.MkdirAll(folderPath, 0755)
		file, err = os.Create(filePath)
		if err != nil {
			log.Println("Error while creating a file.")
			log.Fatal(err)
		}
		fileString := fmt.Sprintf(`{"config_file_path":"%s","token_secret":"%s","repositories":[]}`, filePath, GenerateRandomString(16, 5))
		file.WriteString(fileString)
		file.Close()
		log.Println("Configuration file created. Re-run the previous command.")
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			log.Fatalln(err)
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
		// RepositoryConfiguration.ConfigFilePath = filePath
		file.Close()
	}
}

func findRepository(repoName string) Repository {
	for _, repo := range RepositoryConfiguration.Repositories {
		if repo.Name == repoName {
			return repo
		}
	}
	return Repository{}
}

func generateHash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

func validateToken(repoName string, hashInput string, clientIP string) bool {
	// check hash(TokenSecret + name + token) == token
	if repo := findRepository(repoName); repo.Name == repoName {
		hashedString := generateHash(repo.Name + repo.Token + RepositoryConfiguration.TokenSecret)
		isValidIP := false
		if len(repo.WhiteListedIPs) > 0 {
			for _, ip := range repo.WhiteListedIPs {
				isValidIP = ip == clientIP
			}
		} else {
			isValidIP = true
		}
		return (hashedString == hashInput) && isValidIP
	}
	return false
}
