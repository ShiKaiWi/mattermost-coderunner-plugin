package coderunner

import (
	"sync"
	"testing"
)

func TestRunGO(t *testing.T) {
	code :=
		`package main
import "fmt"
func main() {
	fmt.Print("hello, world")
}`

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, _ := GetCodeRunner(GO)
			out, err := r.Run(code)
			if err != nil {
				t.Errorf("fail to run code, err=%v", err)
				return
			}
			if out != "hello, world" {
				t.Errorf("expect %s, got %s", "hello, world", out)
			}
		}()
	}

	wg.Wait()
}
