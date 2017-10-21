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
	"testing"

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
