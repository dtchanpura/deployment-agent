package manage

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	git "gopkg.in/src-d/go-git.v4"
)

// Function for Initializing test configuration path
func testInit(t *testing.T, removeConfig bool) string {
	// wd, err := os.Getwd()
	// if err != nil {
	// t.Fatal(err)
	// }
	testConfigurationPath := os.Getenv("HOME") + "/.config/cd-go/test/config.json"
	// testConfigurationPath := fmt.Sprintf("/test/config.json", wd)
	if removeConfig {
		err := os.RemoveAll(testConfigurationPath)
		if err != nil {
			t.Fatal(err)
		}
	}

	// os.Setenv("CDGO_CONFIG", testConfigurationPath)

	// First Initializing Configuration file with Test Path
	initializeConfigFile(testConfigurationPath)
	return testConfigurationPath
}

// Getting the Configuration struct from the test file.
func getConfigurationFromPath(t *testing.T, testConfigurationPath string) Configuration {
	file, err := os.Open(testConfigurationPath)
	if err != nil {
		t.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	testConfiguration := Configuration{}
	err = decoder.Decode(&testConfiguration)
	if err != nil {
		t.Fatal(err)
	}
	return testConfiguration
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

func TestInitializeConfigFile(t *testing.T) {
	// First Initializing Configuration file with Test Path
	testInit(t, true)
	// Again Initializing for checking existing path

	testConfigurationPath := testInit(t, false)
	assertEqual(t, len(getConfigurationFromPath(t, testConfigurationPath).Repositories), 0, "")
}

func TestAddConfiguration(t *testing.T) {
	mockName := "test-repo"
	mockPostHook := "repo/hook"
	mockwlip := "127.0.0.1"
	testConfigurationPath := testInit(t, false)
	repoPath := createTestRepository(testConfigurationPath)
	AddConfiguration(mockName, repoPath, mockPostHook, []string{mockwlip})
	testConfiguration := getConfigurationFromPath(t, testConfigurationPath)
	assertEqual(t, len(testConfiguration.Repositories), 1, "")
	assertEqual(t, testConfiguration.Repositories[0].Name, mockName, "")
	assertEqual(t, testConfiguration.Repositories[0].Path, repoPath, "")
	assertEqual(t, testConfiguration.Repositories[0].PostHookPath, mockPostHook, "")
	assertEqual(t, len(testConfiguration.Repositories[0].WhiteListedIPs), 1, "")
	assertEqual(t, testConfiguration.Repositories[0].WhiteListedIPs[0], mockwlip, "")
}

func TestPullRepository(t *testing.T) {

	testConfigurationPath := testInit(t, false)
	// fmt.Println(testConfigurationPath)
	repoPath := createTestRepository(testConfigurationPath)
	err := PullRepository(repoPath, "")
	if err != nil {
		t.Log(err)
	}
}

func createTestRepository(testConfigurationPath string) string {
	repoName := "test-repo"
	folderPathStrings := strings.Split(testConfigurationPath, string(os.PathSeparator))
	folderPath := strings.Join(folderPathStrings[:len(folderPathStrings)-1], string(os.PathSeparator))
	repoPath := strings.Join([]string{folderPath, repoName}, string(os.PathSeparator))
	bareRepoPath := repoPath + ".git"
	// Initializing bare repo.
	git.PlainInit(bareRepoPath, true)
	// Initializing non-bare repo for pulling and cloning

	git.PlainClone(repoPath, false, &git.CloneOptions{URL: bareRepoPath})
	f, _ := os.Create(repoPath + string(os.PathSeparator) + "first")
	f.Close()
	return repoPath
}
