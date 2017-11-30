package listener

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"cgit.dcpri.me/deployment-agent/config"
	"cgit.dcpri.me/deployment-agent/constants"
)

func validateToken(projectUUID, token, clientIP string) bool {
	project, err := config.FindProjectWithUUID(projectUUID)
	if err != nil {
		// fmt.Println(err)
		return false
	}
	return project.ValidateToken(clientIP, token)
}

func executeHooks(project config.Project) {
	// Execute PreHook (if any)
	var isSuccess bool
	isSuccess = true
	if project.PreHook != "" {
		err := executeScript(project.WorkDir, project.PreHook, project.PreHookArgs...)
		if err != nil {
			fmt.Printf("error occurred: %v\n", err)
			isSuccess = false
		}
	}
	if project.PostHook != "" {
		err := executeScript(project.WorkDir, project.PostHook, project.PostHookArgs...)
		if err != nil {
			fmt.Printf("error occurred: %v\n", err)
			isSuccess = false
		}
	}
	if project.ErrorHook != "" && !isSuccess {
		fmt.Println("Error occurred in running prehook and/or posthook")
		err := executeScript(project.WorkDir, project.ErrorHook, project.ErrorHookArgs...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func executeScript(workDir, filePath string, args ...string) error {
	if fileInfo, err := os.Stat(filePath); !os.IsPermission(err) && !os.IsNotExist(err) && fileInfo.Mode()&0111 != 0 {
		cmd := exec.Command(filePath)
		if dirInfo, err := os.Stat(workDir); err == nil && dirInfo.IsDir() {
			cmd.Dir = workDir
		} else {
			fmt.Println(err)
		}
		// err := cmd.Run()
		bts, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Printf("Command Output: %s", string(bts[:]))
		return nil
	}
	return errors.New(constants.ErrorFileNotExecutable)
}
