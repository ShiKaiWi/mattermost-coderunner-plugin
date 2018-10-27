package coderunner

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	runFileDirectory = "/tmp/mattermost-coderunner-plugin"
)

// writeToRunFile makes a work directory whose name same as the filename
// the work directory is returned for clearing by caller
func writeToRunFile(code, fileName string) (string, error) {
	workDir := filepath.Join(runFileDirectory, fileName)
	if err := os.MkdirAll(workDir, 0777); err != nil {
		return "", err
	}

	p := filepath.Join(workDir, fileName)
	return workDir, ioutil.WriteFile(p, []byte(code), 0666)
}

func init() {
	if err := os.MkdirAll(runFileDirectory, 0777); err != nil {
		panic(fmt.Sprintf("fail to create run file directory, err=%v", err))
	}
}
