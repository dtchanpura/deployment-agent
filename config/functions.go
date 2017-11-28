package config

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"

	"cgit.dcpri.me/deployment-agent/constants"
	"github.com/google/uuid"
)

// UpdateConfiguration For updating the configuration.
func UpdateConfiguration(cfgFile string, configuration Configuration, overwrite bool) error {
	configBytes, err := json.Marshal(configuration)
	if err != nil {
		return err
	}
	fmt.Println(string(configBytes[:]))
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) || overwrite {
		err = ioutil.WriteFile(cfgFile, configBytes, 0644)
		if err != nil {
			return err
		}
	} else {
		return errors.New(constants.ErrorFileExists)
	}
	return nil
}

// UpdateProject for adding or modifying the project
func UpdateProject(cfgFile string, project Project) error {
	_, err := FindProject(project.Name, project.UUID)
	if err != nil && err.Error() == constants.ErrorNoProjectFound {
		// TODO Test and remove this following line
		StoredProjects = append(StoredProjects, project)
		updateProjects(cfgFile, StoredProjects)
		return nil
	}
	return errors.New(constants.ErrorProjectAlreadyExists)
}

func updateProjects(cfgFile string, projects []Project) {
	configuration := &Configuration{
		ServeConfig:    StoredServe,
		ProjectConfigs: projects,
	}
	UpdateConfiguration(cfgFile, *configuration, true)
}

func generateHash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	// h.Write([]byte(" "))
	// return fmt.Sprintf("%x", h.Sum(nil))
	return base64.URLEncoding.EncodeToString(h.Sum([]byte("\n"))) // \n added to remove trailing '='
}

// generateRandomString for generating a random string of specified length and
// strength between 1 to 5.
func generateRandomString(length int, strength int) string {
	if strength > 5 {
		strength = 5
	}
	if strength < 1 {
		strength = 1
	}
	var tempString string
	for i := 0; i < strength; i++ {
		tempString += constants.SecretConstants[i]
	}
	bs := make([]byte, length)
	for i := range bs {
		randomIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(tempString))))
		bs[i] = tempString[randomIndex.Int64()]
	}
	return string(bs)
}

// ValidateProjectConfiguration for validating. TODO
func (project *Project) ValidateProjectConfiguration() error {
	if fileInfo, err := os.Stat(project.ErrorHook); os.IsPermission(err) || os.IsNotExist(err) || fileInfo.Mode().IsDir() {
		return errors.New("")
	}
	return nil
}

// ValidateToken is a Project function to check if the token is valid.
func (project *Project) ValidateToken(clientIP string, tokenHash string) bool {
	// Iterate through tokens to find if the token exists or hash matches
	for _, token := range project.Tokens {
		// token.WhitelistedNetwork
		if token.containsIP(clientIP) {
			return tokenHash == generateHash(project.Name+project.Secret+token.Token)
		}
		// fmt.Println(token)
	}
	return false
}

// GetHash for getting hash for given token with index
func (project *Project) GetHash(index int) string {
	return generateHash(project.Name + project.Secret + project.Tokens[index].Token)
}

// NewProject For creating new Project
func NewProject() Project {
	return Project{UUID: uuid.New().String(), Secret: generateRandomString(16, 5)}
}

// NewToken For creating new Token
func NewToken(whitelistCIDR string) TokenDetail {
	return TokenDetail{WhitelistedNetwork: whitelistCIDR, Token: generateRandomString(16, 5)}
}

// TokenDetail function
func (tokenDetail *TokenDetail) containsIP(clientIP string) bool {
	fmt.Println(tokenDetail.WhitelistedNetwork)
	_, ipNet, err := net.ParseCIDR(tokenDetail.WhitelistedNetwork)
	if err != nil {
		fmt.Println(err)
		return false
	}
	ipAddress := net.ParseIP(clientIP)
	return ipNet.Contains(ipAddress)
}

// FindProjectWithUUID for finding the Project from StoredProjects
func FindProjectWithUUID(projectUUID string) (Project, error) {
	for _, project := range StoredProjects {
		if project.UUID == projectUUID {
			return project, nil
		}
	}
	return NewProject(), errors.New(constants.ErrorNoProjectFound)
}

// FindProject for finding the Project from StoredProjects
func FindProject(name, projectUUID string) (Project, error) {
	for _, project := range StoredProjects {
		if project.UUID == projectUUID || project.Name == name {
			return project, nil
		}
	}
	return NewProject(), errors.New(constants.ErrorNoProjectFound)
}
