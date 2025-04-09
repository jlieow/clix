package util

import (
    "testing"
    "fmt"
		"os"
)

func TestSetEnv(t *testing.T) {

	var vars = []RunFunctionVars{
		{
			Key:   "zzzzzzzzzzzzzzzzzz",
			Value: "zzzzzzzzzzzzzzzzzz",
		},
	}

	setEnv(vars)

	value, exists := os.LookupEnv("zzzzzzzzzzzzzzzzzz")
	if exists {
		// Environment variable exists, value is non-empty (can be empty string though)
		fmt.Println("Value:", value)
	} else {
		t.Errorf("Environment variable not set")
	}

}

func TestLoadEnvFile(t *testing.T) {

	var vars = []RunFunctionVars{
		{
			Key:   "dir",
			Value: "../testdata/.env",
		},
	}

	loadEnvFile(vars)

	value, exists := os.LookupEnv("testtesttesttest")
	if exists {
		// Environment variable exists, value is non-empty (can be empty string though)
		fmt.Println("Value:", value)
	} else {
		t.Errorf("Environment variable not set")
	}

}

// TestRunPythonWithValidFile checks 
func TestRunPythonWithValidFile(t *testing.T) {

	var vars = []RunFunctionVars{
		{
			Key:   "path",
			Value: "../testdata/main.py",
		},
	}

	err := runPython(vars)

	if err != nil {
		t.Errorf(`Error: %v`, err)
	}
}

func TestRunPythonWithInvalidFile(t *testing.T) {

	var vars = []RunFunctionVars{
		{
			Key:   "path",
			Value: "../testdata/main.js",
		},
	}

	err := runPython(vars)
	fmt.Println(err)

	if err == nil {
		t.Errorf(`Error: %v`, err)
	}
}

func TestRunPythonWithMissingFile(t *testing.T) {

	var vars = []RunFunctionVars{
		{
			Key:   "path",
			Value: "main.py",
		},
	}

	err := runPython(vars)
	fmt.Println(err)

	if err == nil {
		t.Errorf(`Error: %v`, err)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestRunJavaScriptWithValidFile(t *testing.T) {
	var vars = []RunFunctionVars{
		{
			Key:   "path",
			Value: "../testdata/main.js",
		},
	}

	err := runJavaScript(vars)

	if err != nil {
		t.Errorf(`Error: %v`, err)
	}
}
func TestRunJavaScriptWithInvalidFile(t *testing.T) {
	var vars = []RunFunctionVars{
		{
			Key:   "path",
			Value: "../testdata/main.py",
		},
	}

	err := runJavaScript(vars)
	fmt.Println(err)

	if err == nil {
		t.Errorf(`Error: %v`, err)
	}
}

func TestRunJavaScriptWithMissingFile(t *testing.T) {
	var vars = []RunFunctionVars{
		{
			Key:   "path",
			Value: "main.js",
		},
	}

	err := runJavaScript(vars)
	fmt.Println(err)

	if err == nil {
		t.Errorf(`Error: %v`, err)
	}
}