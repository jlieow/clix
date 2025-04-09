package util

import (
    "testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestRunPython(t *testing.T) {

	var vars = []RunFunctionVars{
		{
			Key:   "path",
			Value: "../testdata/main.py",
		},
	}

	err := runPython(vars)

	if err != nil {
		t.Errorf(`Hello("") = %v, want "", error`, err)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestRunJavaScript(t *testing.T) {
	var vars = []RunFunctionVars{
		{
			Key:   "path",
			Value: "../testdata/main.js",
		},
	}

	err := runJavaScript(vars)

	if err != nil {
		t.Errorf(`Hello("") = %v, want "", error`, err)
	}
}