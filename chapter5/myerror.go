package chapter5

import (
	"fmt"
	"runtime/debug"
)

type MyError struct {
	Inner      error
	Message    string
	StackTrace string
	Misc       map[string]interface{}
}

func wrap(err error, format string, args ...interface{}) MyError {
	return MyError{
		Inner:      err,
		Message:    fmt.Sprintf(format, args...),
		StackTrace: string(debug.Stack()),
		Misc:       make(map[string]interface{}),
	}
}

func (e MyError) Error() string {
	return e.Message
}
