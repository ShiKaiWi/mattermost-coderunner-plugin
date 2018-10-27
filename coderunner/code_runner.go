package coderunner

import "errors"

// Lang defines language type
type Lang string

const (
	// GO defines golang type
	GO Lang = "golang"
	// JS defines javscript type
	JS Lang = "javascript"
)

// ErrNotSupport defines the error not supporting given language
var ErrNotSupport = errors.New("not support such language")

// CodeRunner defines the interface for running some code
type CodeRunner interface {
	Run(code string) (string, error)
}

func buildCodeRunner(lang Lang) CodeRunner {
	switch lang {
	case GO:
		return &goCodeRunner{}
	case JS:
		return &jsCodeRunner{}
	default:
		panic("unsupport language type")
	}
}

var muxer *codeRunnerMuxer

func init() {
	muxer = &codeRunnerMuxer{
		rm: make(map[Lang]CodeRunner),
	}
	muxer.rm[GO] = buildCodeRunner(GO)
	muxer.rm[JS] = buildCodeRunner(JS)
}

type codeRunnerMuxer struct {
	rm map[Lang]CodeRunner
}

// GetCodeRunner returns a code runner which can run `lang` code
// CodeRunner won't be nil if error is nil
// error can be ErrNotSupport
func GetCodeRunner(lang Lang) (CodeRunner, error) {
	r, ok := muxer.rm[lang]
	if !ok {
		return nil, ErrNotSupport
	}
	return r, nil
}
