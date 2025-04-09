package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func helloWorld(vars []RunFunctionVars) {

	// Convert to map for easy lookup
	argMap := make(map[string]string)
	for _, v := range vars {
		argMap[v.Key] = v.Value
	}

	// Extract values in function parameter order
	name := argMap["name"]
	age := argMap["age"]

	fmt.Println("Hello " + name + " you are currently " + age)
}

func setEnv(vars []RunFunctionVars) error {
	for _, v := range vars {

		// Set environment variable
		err := os.Setenv(v.Key, v.Value)
		if err != nil {
			return fmt.Errorf("Error setting environment variable: %s", err)
		}

	}

	return nil
}

func loadEnvFile(vars []RunFunctionVars) error {

	// Convert to map for easy lookup
	argMap := make(map[string]string)
	for _, v := range vars {
		argMap[v.Key] = v.Value
	}

	// Extract values in function parameter order
	dir := argMap["dir"]
	// Check if the variable is empty and assign a default value
	if dir == "" {
		dir = ".env"
	}

	// Load the .env file from the current working directory (where the app is called)
	err := godotenv.Load(dir) // This will automatically load the .env file from the current directory
	if err != nil {
		return fmt.Errorf("Error loading .env file")
	}
	return nil
}

func printAllEnv() {
	// Get all environment variables
	envVars := os.Environ()

	// Print each environment variable
	for _, envVar := range envVars {
		fmt.Println(envVar)
	}
}

// isPythonScript checks if the given path is a Python script
func isPythonScript(path string) (bool, error) {
	// Check if file exists and is a regular file
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("file does not exist")
		}
		return false, err
	}
	if info.IsDir() {
		return false, fmt.Errorf("path is a directory, not a file")
	}

	// Check if the file has a .py extension
	if !strings.HasSuffix(info.Name(), ".py") {
		return false, nil // Not a Python file
	}

	// Optionally, check if it has a Python shebang (e.g., #!/usr/bin/env python)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}

	// Check for Python shebang at the beginning of the file
	if len(data) >= 2 && string(data[:2]) == "#!" && strings.Contains(string(data), "python") {
		return true, nil
	}

	// If no shebang found, but file ends with .py extension, return true
	return true, nil
}

func execPythonScript(scriptPath string) {
	// Run the Python script using the Python interpreter
	cmd := exec.Command("python", scriptPath)

	// Get the output (both stdout and stderr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing Python script: %v\n", err)
	}

	// Print the output
	fmt.Println("Python Script Output:\n", string(output))
}

func runPython(vars []RunFunctionVars) error {

	// Convert to map for easy lookup
	argMap := make(map[string]string)
	for _, v := range vars {
		argMap[v.Key] = v.Value
	}

	// Extract values in function parameter order
	path := argMap["path"]

	isPython, err := isPythonScript(path)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	if !isPython {
		return fmt.Errorf("Please provide a python script.")
	}

	// Execute the Python script
	execPythonScript(path)

	return nil
}

// isJavaScriptScript checks if the given path is a JavaScript file
func isJavaScriptScript(path string) (bool, error) {
	// Check if file exists and is a regular file
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, fmt.Errorf("file does not exist")
		}
		return false, err
	}
	if info.IsDir() {
		return false, fmt.Errorf("path is a directory, not a file")
	}

	// Check if the file has a .js extension
	if !strings.HasSuffix(info.Name(), ".js") {
		return false, nil // Not a JavaScript file
	}

	// Optionally, check if it has a JavaScript shebang or comment header
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}

	// Check for a JavaScript shebang at the beginning of the file (e.g., #!/usr/bin/env node)
	if len(data) >= 2 && string(data[:2]) == "#!" && strings.Contains(string(data), "node") {
		return true, nil
	}

	// If file ends with .js extension, return true
	return true, nil
}

// runJavaScriptScript runs the given JavaScript file using Node.js
func runJavaScriptScript(scriptPath string) {
	// Run the JavaScript file using Node.js
	cmd := exec.Command("node", scriptPath)

	// Get the output (both stdout and stderr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing JavaScript script: %v\n", err)
	}

	// Print the output
	fmt.Println("JavaScript Script Output:\n", string(output))
}

func runJavaScript(vars []RunFunctionVars) error {
	// Convert to map for easy lookup
	argMap := make(map[string]string)
	for _, v := range vars {
		argMap[v.Key] = v.Value
	}

	// Extract values in function parameter order
	path := argMap["path"]

	isJS, err := isJavaScriptScript(path)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	if !isJS {
		return fmt.Errorf("Please provide a JavaScript file.")
	}

	// Run the JavaScript script
	runJavaScriptScript(path)

	return nil
}
