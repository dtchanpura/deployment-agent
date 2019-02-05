package utils

import (
	"errors"
	"fmt"
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
