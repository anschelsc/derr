package derr

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// An Error remembers the file and line number which created it.
type Error struct {
	File string
	Line int
	Err  error
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s %d: %s", e.File, e.Line, e.Err)
}

func New(e error) *Error {
	if r, ok := e.(*Error); ok {
		return r
	}
	ret := &Error{Err: e}
	_, ret.File, ret.Line, _ = runtime.Caller(1)
	ret.File = filepath.Base(ret.File)
	return ret
}
