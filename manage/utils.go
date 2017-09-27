package manage

import (
	"gopkg.in/src-d/go-git.v4"
	"math/big"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"os/exec"
	"log"
)

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

func ExecuteHook(hookPath string) error {
	cmd := exec.Command(hookPath)
	err := cmd.Run()
	log.Println()
	return err
}

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

func validateToken(repoName string, hashInput string) bool {
	// check hash(TokenSecret + name + token) == token
	if repo := findRepository(repoName); repo.Name == repoName {
		hashedString := generateHash(repo.Name + repo.Token + RepositoryConfiguration.TokenSecret)
		return hashedString == hashInput
	}
	return false
}
