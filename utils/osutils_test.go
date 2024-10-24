// Copyright 2024 NetApp, Inc. All Rights Reserved.

package utils

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

var (
	execReturnValue string
	execReturnCode  int
	execPadding     int
	execDelay       time.Duration
)

// TestShellProcess is a method that is called as a substitute for a shell command,
// the GO_TEST flag ensures that if it is called as part of the test suite, it is
// skipped. GO_TEST_RETURN_VALUE flag allows the caller to specify a base64 encoded version of what should be returned via stdout,
// GO_TEST_RETURN_CODE flag allows the caller to specify what the return code should be, and
// GO_TEST_DELAY flag allows the caller to inject a delay before the function returns.
// GO_TEST_RETURN_PADDING_LENGTH flag allows the caller to specify how many bytes to return in total,
// if GO_TEST_RETURN_VALUE does not use all the bytes, the end is padded with NUL data
func TestShellProcess(t *testing.T) {
	if os.Getenv("GO_TEST") != "1" {
		return
	}
	// Print out the test value to stdout
	returnString, _ := b64.StdEncoding.DecodeString(os.Getenv("GO_TEST_RETURN_VALUE"))
	fmt.Fprintf(os.Stdout, string(returnString))
	if os.Getenv("GO_TEST_RETURN_PADDING_LENGTH") != "" {
		padLength, _ := strconv.Atoi(os.Getenv("GO_TEST_RETURN_PADDING_LENGTH"))
		padString := make([]byte, padLength-len(returnString))
		fmt.Fprintf(os.Stdout, string(padString))
	}
	code, err := strconv.Atoi(os.Getenv("GO_TEST_RETURN_CODE"))
	if err != nil {
		code = -1
	}
	// Pause for some amount of time
	delay, err := time.ParseDuration(os.Getenv("GO_TEST_DELAY"))
	if err == nil {
		time.Sleep(delay)
	}
	os.Exit(code)
}
