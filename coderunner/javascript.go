package coderunner

import "fmt"

const nodeImageName = "node"

type jsCodeRunner struct {
	baseCodeRunner
}

func (r *jsCodeRunner) Run(code string) (string, error) {
	nameGen := func(id int32) string {
		return fmt.Sprintf("javascript_%d.js", id)
	}
	commandGen := func(fileName string) string {
		return fmt.Sprintf(`node "%s"`, fileName)
	}
	return r.RunCode(nodeImageName, code, nameGen, commandGen)
}
