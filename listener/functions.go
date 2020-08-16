package listener

import (
	"fmt"
	"strings"

	"github.com/dtchanpura/deployment-agent/config"
	"github.com/dtchanpura/deployment-agent/utils"
)

const (
	errURLNotValid         = "invalid url: %s"
	errCredentialsNotValid = "invalid credentials: %s"
)

func getCredentials(path string) (string, string, error) {
	if !strings.HasPrefix(path, "/reload/") {
		// Not possible, as it is called from the handler having /reload/
		return "", "", fmt.Errorf(errURLNotValid, path)
	}
	path = strings.TrimPrefix(path, "/reload/")
	paths := strings.Split(path, "/")
	if len(paths) != 2 {
		return "", "", fmt.Errorf(errCredentialsNotValid, path)
	}
	return paths[0], paths[1], nil
}

func validateToken(projectUUID, token, clientIP string) bool {
	project, err := config.FindProjectWithUUID(projectUUID)
	if err != nil {
		logger.Error().Err(err).Send()
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
			logger.Error().Err(err).Send()
			isSuccess = false
		}
	}
	if project.PostHook != "" {
		err := utils.ExecuteScript(project.WorkDir, project.PostHook, project.PostHookArgs...)
		if err != nil {
			logger.Error().Err(err).Send()
			isSuccess = false
		}
	}
	// Following is the replacement for above code.
	if len(project.Hooks) > 0 {
		for _, hook := range project.Hooks {
			allowedMaxArgs := project.MaxArgs
			if hook.MaxArgs != 0 {
				allowedMaxArgs = hook.MaxArgs
			}
			maxArgs := len(args)
			if allowedMaxArgs != -1 && maxArgs >= allowedMaxArgs {
				maxArgs = allowedMaxArgs
			}
			if hook.FilePath != "" {
				// maxArgs is for limiting number of arguments
				err := utils.ExecuteScript(project.WorkDir, hook.FilePath, args[:maxArgs]...)
				if err != nil {
					isSuccess = false
					logger.Error().Err(err).Send()
				}
			}
		}
	}
	if project.ErrorHook != "" && !isSuccess {
		fmt.Println("Error occurred in running prehook and/or posthook")
		err := utils.ExecuteScript(project.WorkDir, project.ErrorHook, project.ErrorHookArgs...)
		if err != nil {
			logger.Error().Err(err).Send()
		}
	}
}
