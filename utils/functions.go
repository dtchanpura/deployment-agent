package utils

import (
	"errors"
	"log"
	"os"
	"os/exec"

	"github.com/dtchanpura/deployment-agent/constants"
)

// ExecuteScript for executing script
func ExecuteScript(workdir string, execpath string, args ...string) error {
	if fileInfo, err := os.Stat(execpath); !os.IsPermission(err) && !os.IsNotExist(err) && fileInfo.Mode()&0111 != 0 {
		cmd := exec.Command(execpath, args...)
		if dirInfo, err := os.Stat(workdir); err == nil && dirInfo.IsDir() {
			cmd.Dir = workdir
		} else {
			log.Println(err)
		}
		// err := cmd.Run()
		outputBytes, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Command Error: %s\n", string(outputBytes[:]))
			return err
		}
		log.Printf("Command Output: %s\n", string(outputBytes[:]))
		return nil
	}
	return errors.New(constants.ErrorFileNotExecutable)
}
