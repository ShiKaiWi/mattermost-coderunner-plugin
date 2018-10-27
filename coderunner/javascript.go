package coderunner

type jsCodeRunner struct{}

func (r *jsCodeRunner) Run(code string) (string, error) {
	return "result", nil
}
