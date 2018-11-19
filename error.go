// Package goerror implements functions to manipulate errors.
// The traditional `error` handling method in golang is given below
// which applied recursively up the call stack results in error reports
// without context or debugging information.

// if err != nil {
//         return err
// }

// The package mainly contain two functions
//  GetErrorInfo
//  New

// GetErrorInfo

//  func GetErrorInfo(err error) error

// which accepts the error and found the detailed information about the error.
// GetErrorInfo returns the error wrapped up with the file name,
// line number and function name of that particular error.

// if err != nil {
//         return goerror.GetErrorInfo(err)
// }

// New

//  func New(err string) error

// Which is used to create a new errors.
// New help us to create custom user defined errors.

// err := goerror.New("New custom error")

package goerror

import (
	"fmt"
	"runtime"
)

type errorValue struct {
	s string
}

// GetErrorInfo returns the corresponding error informations.
// Including the file name , line number, function name
func GetErrorInfo(err error) error {
	if err != nil {
		pc, file, line, ok := runtime.Caller(1)
		if ok {
			funcName := runtime.FuncForPC(pc).Name()
			f := format(file, "file")
			fn := format(funcName, "function")
			errorMsg := fmt.Sprintf("file name=> %s : line number=> %d : function name=> %s : error=> %s", f, line, fn, err)
			return &errorValue{errorMsg}
		}
		return nil
	}
	return nil
}

// GetErrorType accepts an error
// and returns the type of the error
func GetErrorType(err error) error {
	if err != nil {
		errorMsg := fmt.Sprintf("error=> %s : error type=> %T", err, err)
		return &errorValue{errorMsg}
	}
	return nil
}

// New returns a new error.
func New(err string) error {
	return &errorValue{err}
}

func (e *errorValue) Error() string {
	return e.s
}
