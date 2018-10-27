package coderunner

import (
	"fmt"
	"sync/atomic"
)

const dockerRunScriptTmpl = `docker run -v "%s":/usr/src/coderunner -w /usr/src/coderunner --rm golang go run "%s"`

type goCodeRunner struct {
	id int32
}

func (r *goCodeRunner) Run(code string) (string, error) {
	runFileName := fmt.Sprintf("golang_%d.go", atomic.AddInt32(&r.id, 1))
	p, err := writeToRunFile(code, runFileName)
	if err != nil {
		return "", err
	}
	defer removeRunFile(p)

	script := fmt.Sprintf(dockerRunScriptTmpl, runFileDirectory, runFileName)

	return runCommand(script)
}
