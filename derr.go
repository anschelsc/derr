package derr

import (
	"fmt"
	"path/filepath"
	"runtime"
)

var On = true

// An Error remembers the file and line number which created it.
type Error struct {
	File string
	Line int
	Err  error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %d: %s", e.File, e.Line, e.Err)
}

// If On is set to true, returns a *Error containing e; otherwise, returns e.
func New(e error) error {
	if !On {
		return e
	}
	if r, ok := e.(*Error); ok {
		return r
	}
	ret := &Error{Err: e}
	_, ret.File, ret.Line, _ = runtime.Caller(1)
	ret.File = filepath.Base(ret.File)
	return ret
}
