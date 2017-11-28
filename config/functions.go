package config

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"cgit.dcpri.me/deployment-agent/errorconstants"
	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

// UpdateConfiguration For updating the configuration.
func UpdateConfiguration(cfgFile string, configuration Configuration, overwrite bool) error {
	configBytes, err := yaml.Marshal(configuration)
	if err != nil {
		return err
	}
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) || overwrite {
		err = ioutil.WriteFile(cfgFile, configBytes, 0644)
		if err != nil {
			return err
		}
	} else {
		return errors.New(errorconstants.ErrorFileExists)
	}
	return nil
}

// UpdateProject for adding or modifying the project
func UpdateProject(cfgFile string, project Project) {
	_, err := FindProject(project.UUID)
	if err.Error() == errorconstants.ErrorNoProjectFound {
		StoredProjects = append(StoredProjects, project)
		updateProjects(cfgFile, StoredProjects)
	}
}

func updateProjects(cfgFile string, projects []Project) {
	configuration := Configuration{
		ServeConfig:    StoredServe,
		ProjectConfigs: projects,
	}
	UpdateConfiguration(cfgFile, configuration, true)
}

func generateHash(input string) string {
	h := sha256.New()
	h.Write([]byte(input))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}

// ValidateToken is a Project function to check if the token is valid.
func (project *Project) ValidateToken(clientIP string, tokenHash string) bool {
	// Iterate through tokens to find if the token exists or hash matches
	for _, token := range project.Tokens {
		if token.containsIP(clientIP) {
			return tokenHash == generateHash(project.Name+project.Secret+token.Token)
		}
		// fmt.Println(token)
	}
	return false
}

// NewProject For creating new Project, TODO to be used in Add command
func NewProject() *Project {
	return &Project{UUID: uuid.New().String()}
}

// TokenDetail function
func (tokenDetail *TokenDetail) containsIP(clientIP string) bool {
	_, ipNet, err := net.ParseCIDR(tokenDetail.WhitelistedNetwork)
	if err != nil {
		fmt.Println(err)
		return false
	}
	ipAddress := net.ParseIP(clientIP)
	return ipNet.Contains(ipAddress)
}

// FindProject for finding the Project from StoredProjects
func FindProject(projectUUID string) (Project, error) {
	for _, project := range StoredProjects {
		if project.UUID == projectUUID {
			return project, nil
		}
	}
	return *NewProject(), errors.New(errorconstants.ErrorNoProjectFound)
}
