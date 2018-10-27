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

func writeToRunFile(code, fileName string) (string, error) {
	p := filepath.Join(runFileDirectory, fileName)
	return p, ioutil.WriteFile(p, []byte(code), 0666)
}

func removeRunFile(fileName string) error {
	p := filepath.Join(runFileDirectory, fileName)
	return os.Remove(p)
}

func init() {
	if err := os.MkdirAll(runFileDirectory, 0777); err != nil {
		panic(fmt.Sprintf("fail to create run file directory, err=%v", err))
	}
}
