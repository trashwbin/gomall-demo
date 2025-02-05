package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

// MustHandleError is a function used to handle errors.
// When an error occurs (i.e., the error object is not nil), it logs the error and terminates the program.
// This function has no return value.
func MustHandleError(err error) {
	// If the error is not nil, log the error and terminate the program
	if err != nil {
		hlog.Fatal(err)
	}
}
