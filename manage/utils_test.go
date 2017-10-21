package manage

import (
	"fmt"
	"os"
	"testing"
)

func TestAddConfiguration(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	testPath := fmt.Sprintf("%s/test", wd)
	os.Setenv("CDGO_CONFIG", testPath)
	// init()

}

func TestPullRepository(t *testing.T) {

}
