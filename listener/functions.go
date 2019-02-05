package listener

import (
	"fmt"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/dtchanpura/deployment-agent/utils"
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
		err := utils.ExecuteScript(project.WorkDir, project.PreHook, project.PreHookArgs...)
		if err != nil {
			fmt.Printf("error occurred: %v\n", err)
			isSuccess = false
		}
	}
	if project.PostHook != "" {
		err := utils.ExecuteScript(project.WorkDir, project.PostHook, project.PostHookArgs...)
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
				err := utils.ExecuteScript(project.WorkDir, hook.FilePath, args[:maxArgs]...)
				if err != nil {
					isSuccess = false
					fmt.Println(err)
				}
			}
		}
	}
	if project.ErrorHook != "" && !isSuccess {
		fmt.Println("Error occurred in running prehook and/or posthook")
		err := utils.ExecuteScript(project.WorkDir, project.ErrorHook, project.ErrorHookArgs...)
		if err != nil {
			fmt.Println(err)
		}
	}
}
