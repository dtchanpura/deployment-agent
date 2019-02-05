package config

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net"
	"os"

	"github.com/dtchanpura/deployment-agent/constants"
	"github.com/dtchanpura/deployment-agent/utils"
	"github.com/google/uuid"
	yaml "gopkg.in/yaml.v2"
)

// UpdateConfiguration For updating the configuration.
func UpdateConfiguration(cfgFile string, configuration Configuration, overwrite bool) error {
	configBytes, err := yaml.Marshal(configuration)
	if err != nil {
		return err
	}
	// fmt.Println(string(configBytes[:]))
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
		updateProjects(cfgFile, project)
		// fmt.Println(err)
		return nil
	}
	return errors.New(constants.ErrorProjectAlreadyExists)
}

func updateProjects(cfgFile string, projects ...Project) {
	StoredProjects = append(StoredProjects, projects...)
	configuration := Configuration{
		ServeConfig:    StoredServe,
		ProjectConfigs: StoredProjects,
	}
	UpdateConfiguration(cfgFile, configuration, true)
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
	ok := false
	for _, token := range project.Tokens {
		// token.WhitelistedNetwork
		// fmt.Println(tokenHash)
		// fmt.Println(generateHash(project.Name + project.Secret + token.Token))
		if token.containsIP(clientIP) {
			if tokenHash == generateHash(project.Name+project.Secret+token.Token) {
				return true
			}
		}
		// fmt.Println(token)
	}
	return ok
}

// GetHash for getting hash for given token with index
func (project *Project) GetHash(index int) string {
	return generateHash(project.Name + project.Secret + project.Tokens[index].Token)
}

// NewProject For creating new Project
func NewProject(ipCIDRs ...string) Project {
	tokens := []TokenDetail{}
	for _, ipCIDR := range ipCIDRs {
		tokens = append(tokens, NewToken(ipCIDR))
	}
	projectUUID := uuid.New().String()
	secret := generateRandomString(16, 5)
	return Project{UUID: projectUUID, Secret: secret, Tokens: tokens}
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
	return Project{}, errors.New(constants.ErrorNoProjectFound)
}

// FindProject for finding the Project from StoredProjects
func FindProject(name, projectUUID string) (Project, error) {
	// fmt.Println(name)
	// fmt.Println(projectUUID)
	for _, project := range StoredProjects {
		if project.UUID == projectUUID || project.Name == name {
			return project, nil
		}
	}
	return Project{}, errors.New(constants.ErrorNoProjectFound)
}

// DecodeProjectConfiguration for decoding the project configuration from viper.AllSettings()
func DecodeProjectConfiguration(settingsMap map[string]interface{}) error {
	// config.StoredProjects
	// projects := viper.AllSettings()["projects"].([]interface{})

	projects := []Project{}
	if prjs, ok := settingsMap["projects"]; ok {
		for _, prj := range prjs.([]interface{}) {
			p := prj.(map[interface{}]interface{})
			projectStruct := Project{
				Tokens: []TokenDetail{},
			}
			for key, value := range p {
				switch key.(string) {
				case "name":
					projectStruct.Name = value.(string)
				case "error_hook":
					projectStruct.ErrorHook = value.(string)
				case "pre_hook":
					projectStruct.PreHook = value.(string)
				case "post_hook":
					projectStruct.PostHook = value.(string)
				// case "remote_path":
				// projectStruct.RemotePath = value.(string)
				case "work_dir":
					projectStruct.WorkDir = value.(string)
				case "uuid":
					projectStruct.UUID = value.(string)
				case "secret":
					projectStruct.Secret = value.(string)
				case "hooks":
					projectStruct.Hooks = []Hook{}
					for _, v := range value.([]interface{}) {
						projectStruct.Hooks = append(projectStruct.Hooks, decodeHook(v.(map[interface{}]interface{})))
					}
				}
			}

			if tokens, hasTokens := p["tokens"]; hasTokens {
				// tokens := p["tokens"].([]interface{})
				for _, token := range tokens.([]interface{}) {
					token := token.(map[interface{}]interface{})
					tokenStruct := TokenDetail{}
					for tokenK, tokenV := range token {
						switch tokenK.(string) {
						case "whitelistnet":
							tokenStruct.WhitelistedNetwork = tokenV.(string)
						case "token":
							tokenStruct.Token = tokenV.(string)
						}
					}
					projectStruct.Tokens = append(projectStruct.Tokens, tokenStruct)
				}
			}
			projects = append(projects, projectStruct)
		}
	}
	StoredProjects = projects
	return nil
}

func decodeHook(h map[interface{}]interface{}) Hook {
	hook := Hook{}
	for k, v := range h {
		switch k {
		case "file_path":
			hook.FilePath = v.(string)
		case "max_args":
			hook.MaxArgs = v.(int)
		}
	}
	return hook
}

// NewHooks for getting a new []Hook to be assigned to project.Hooks
func NewHooks(hooks ...string) []Hook {
	Hooks := make([]Hook, 0)
	for _, hook := range hooks {
		Hooks = append(Hooks, Hook{
			FilePath: hook,
		})
	}
	return Hooks
}

// Execution part

// ExecuteHooks for executing hooks one after another
func (project *Project) ExecuteHooks(args ...string) {
	isSuccess := true
	// Following is the replacement for above code.
	if len(project.Hooks) > 0 {
		for _, hook := range project.Hooks {
			maxArgs := len(args)
			if hook.MaxArgs != -1 {
				maxArgs = hook.MaxArgs
			}
			if hook.FilePath != "" {
				// TODO: Change this project.WorkDir
				// hook.MaxArgs is for limiting number of arguments
				err := utils.ExecuteScript(project.WorkDir, hook.FilePath, args[:maxArgs]...)
				if err != nil {
					isSuccess = false
					fmt.Println(err)
				}
			}
		}
	}
	if project.ErrorHook != "" && !isSuccess {
		fmt.Println("Error occurred in running hook(s)")
		err := utils.ExecuteScript(project.WorkDir, project.ErrorHook, project.ErrorHookArgs...)
		if err != nil {
			fmt.Println(err)
		}
	}

}
