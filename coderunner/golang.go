package coderunner

import (
	"fmt"
	"os"
)

const dockerRunGOScriptTmpl = `docker run -v "%s":/usr/src/coderunner -w /usr/src/coderunner --rm golang go run "%s"`

type goCodeRunner struct {
	idGen
}

func (r *goCodeRunner) Run(code string) (string, error) {
	runFileName := fmt.Sprintf("golang_%d.go", r.genID())
	p, err := writeToRunFile(code, runFileName)
	if err != nil {
		return "", err
	}
	defer os.Remove(p)

	script := fmt.Sprintf(dockerRunGOScriptTmpl, runFileDirectory, runFileName)

	return runCommand(script)
}
