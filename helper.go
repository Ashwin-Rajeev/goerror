// helper.go contain all the helper function.
// helper functions used for formatting operations.

package goerror

import (
	"strings"
)

// format is used to format file and function names.
// based on the name format call another functions.
func format(val, name string) string {
	switch name {
	case "file":
		return splitFile(val)
	case "function":
		return splitFunction(val)
	}
	return ""
}

// split file split the file name.
// make it easily readable.
func splitFile(file string) string {
	s := strings.Split(file, "/")
	if l := len(s); l > 0 {
		return s[l-1]
	}
	return s[0]
}

// split function split the function name.
// make it easily readable.
func splitFunction(fun string) string {
	s := strings.Split(fun, ".")
	if l := len(s); l > 0 {
		return s[l-1]
	}
	return s[0]
}
