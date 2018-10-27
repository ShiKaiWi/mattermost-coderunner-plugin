package coderunner

import (
	"fmt"
	"os"
)

const dockerRunNodeScriptTmpl = `docker run -v "%s":/usr/src/coderunner -w /usr/src/coderunner --rm node node "%s"`

type jsCodeRunner struct {
	idGen
}

func (r *jsCodeRunner) Run(code string) (string, error) {
	runFileName := fmt.Sprintf("javascirpt_%d.go", r.genID())
	p, err := writeToRunFile(code, runFileName)
	if err != nil {
		return "", err
	}
	defer os.Remove(p)

	script := fmt.Sprintf(dockerRunNodeScriptTmpl, runFileDirectory, runFileName)

	return runCommand(script)
}
