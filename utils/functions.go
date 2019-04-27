package utils

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/dtchanpura/deployment-agent/constants"
)

// ExecuteScript for executing script
func ExecuteScript(workdir string, execpath string, args ...string) error {
	if fileInfo, err := os.Stat(execpath); !os.IsPermission(err) && !os.IsNotExist(err) && fileInfo.Mode()&0111 != 0 {
		cmd := exec.Command(execpath, args...)
		if dirInfo, err := os.Stat(workdir); err == nil && dirInfo.IsDir() {
			cmd.Dir = workdir
		} else {
			logWithTimestampln(err.Error())
		}
		// err := cmd.Run()
		outputBytes, err := cmd.CombinedOutput()
		if err != nil {
			logWithTimestampf("Command Error: %s\n", string(outputBytes[:]))
			return err
		}
		logWithTimestampf("Command Output: %s\n", string(outputBytes[:]))
		return nil
	}
	return errors.New(constants.ErrorFileNotExecutable)
}

func logWithTimestampf(message string, args ...interface{}) {
	log.Printf(time.Now().Format("[ 2006/01/02 15:04:05 ] ")+message, args...)
}

func logWithTimestampln(message string) {
	logWithTimestampf(message + "\n")
}
