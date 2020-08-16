package utils

import (
	"errors"
	"os"
	"os/exec"

	"github.com/rs/zerolog"

	"github.com/dtchanpura/deployment-agent/constants"
)

var (
	logger = zerolog.New(os.Stdout)
)

// ExecuteScript for executing script
func ExecuteScript(workdir string, execpath string, args ...string) error {
	if fileInfo, err := os.Stat(execpath); !os.IsPermission(err) && !os.IsNotExist(err) && fileInfo.Mode()&0111 != 0 {
		logger.Debug().Fields(map[string]interface{}{"workdir": workdir, "script": execpath, "args": args}).Send()
		cmd := exec.Command(execpath, args...)
		if dirInfo, err := os.Stat(workdir); err == nil && dirInfo.IsDir() {
			cmd.Dir = workdir
		} else {
			logger.Error().Err(err).Send()
		}
		// err := cmd.Run()
		outputBytes, err := cmd.CombinedOutput()
		if err != nil {
			logger.Error().Str("output", string(outputBytes[:])).Msg("Command Error")
			return err
		}
		logger.Info().Str("output", string(outputBytes[:])).Msg("Command Output")
		return nil
	}
	return errors.New(constants.ErrorFileNotExecutable)
}
