package coderunner

import (
	"fmt"
	"os"
	"sync/atomic"
)

const dockerRunScriptTmpl = `docker run -v "%s":/usr/src/coderunner -w /usr/src/coderunner --rm %s %s`

type baseCodeRunner struct {
	id int32
}

func (r *baseCodeRunner) genID() int32 {
	return atomic.AddInt32(&r.id, 1)
}

func (r *baseCodeRunner) RunCode(imageName, code string, nameGen func(id int32) string, commandGen func(fileName string) string) (string, error) {
	runFileName := nameGen(r.genID())
	workDir, err := writeToRunFile(code, runFileName)
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(workDir)

	script := fmt.Sprintf(dockerRunScriptTmpl, workDir, imageName, commandGen(runFileName))

	return runCommand(script)
}
