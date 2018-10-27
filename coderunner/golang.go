package coderunner

import "fmt"

const (
	golangImageName = "golang"
)

type goCodeRunner struct {
	baseCodeRunner
}

func (r *goCodeRunner) Run(code string) (string, error) {
	nameGen := func(id int32) string {
		return fmt.Sprintf("golang_%d.go", id)
	}
	commandGen := func(fileName string) string {
		return fmt.Sprintf(`go run "%s"`, fileName)
	}
	return r.RunCode(golangImageName, code, nameGen, commandGen)
}
