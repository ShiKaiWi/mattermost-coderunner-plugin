package coderunner

import (
	"sync"
	"testing"
)

func testRunCode(lang Lang, code, expected string, t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			r, _ := GetCodeRunner(lang)
			out, err := r.Run(code)
			if err != nil {
				t.Errorf("fail to run code, err=%v", err)
				return
			}
			if out != expected {
				t.Errorf("expect %s, got %s", expected, out)
			}
		}()
	}

	wg.Wait()
}
func TestRunGO(t *testing.T) {
	code :=
		`package main
import "fmt"
func main() {
	fmt.Print("hello, world")
}`
	testRunCode(GO, code, "hello, world", t)
}

func TestRunNode(t *testing.T) {
	code :=
		`const h = "hello"
const w = "world"
const s = h+", "+w
console.log(s)
`
	testRunCode(JS, code, "hello, world\n", t)
}
