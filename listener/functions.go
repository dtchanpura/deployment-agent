package listener

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/dtchanpura/deployment-agent/constants"
)

func validateToken(projectUUID, token, clientIP string) bool {
	project, err := config.FindProjectWithUUID(projectUUID)
	if err != nil {
		// fmt.Println(err)
		return false
	}
	return project.ValidateToken(clientIP, token)
}

func executeHooks(project config.Project, args ...string) {
	// Execute PreHook (if any)
	isSuccess := true
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
				err := executeScript(project.WorkDir, hook.FilePath, args[:maxArgs]...)
				if err != nil {
					fmt.Println(err)
				}
			}
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
		cmd := exec.Command(filePath, args...)
		if dirInfo, err := os.Stat(workDir); err == nil && dirInfo.IsDir() {
			cmd.Dir = workDir
		} else {
			fmt.Println(err)
		}
		// err := cmd.Run()
		outputBytes, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Command Error: %s\n", string(outputBytes[:]))
			return err
		}
		fmt.Printf("Command Output: %s\n", string(outputBytes[:]))
		return nil
	}
	return errors.New(constants.ErrorFileNotExecutable)
}
