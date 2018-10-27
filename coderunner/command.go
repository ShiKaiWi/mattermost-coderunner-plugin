package coderunner

import (
	"bytes"
	"os/exec"
)

// error message will be returned if error is not nil
func runCommand(script string) (string, error) {
	var out bytes.Buffer
	cmd := exec.Command("sh", "-c", script)
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}
