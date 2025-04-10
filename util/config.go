package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

// RunFunctionVars represents the variables associated with a run_function
type RunFunctionVars struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Task represents a task to perform before running the command
type Hooks struct {
	RunCommand      string            `json:"run_command,omitempty"`
	RunFunction     string            `json:"run_function,omitempty"`
	RunFunctionVars []RunFunctionVars `json:"run_function_vars,omitempty"`
}

// Command represents the structure for each command to execute
type Command struct {
	Command     string  `json:"command"`
	Description string  `json:"description"`
	PreHooks    []Hooks `json:"prehooks"`
	PostHooks   []Hooks `json:"posthooks"`
}

// Config represents the top-level structure with commands
type Config struct {
	Commands map[string]Command `json:"commands"`
}

// GetConfigFilePath determines the platform-specific path for the config file
func GetConfigFilePath() string {

	// Get the home directory of the user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error retrieving home directory:", err)
		return ""
	}

	var configDir string

	// Determine the correct path based on the OS
	switch runtime.GOOS {
	case "linux", "darwin":
		// For Linux/macOS, store in ~/.config/myapp/config.json
		configDir = filepath.Join(homeDir, ".config", "clix", "config.json")
	case "windows":
		// For Windows, store in C:\Users\<username>\AppData\Roaming\MyApp\config.json
		configDir = filepath.Join(homeDir, "AppData", "Roaming", "CliX", "config.json")
	default:
		fmt.Println("Unsupported OS")
		return ""
	}

	return configDir
}

func GetListConfigCommand() []string {
	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	configPath := GetConfigFilePath()

	log.Printf("configPath: " + configPath)

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Declare a Config object to hold the parsed data
	var config Config

	// Parse the JSON string into the config object
	err = json.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	var keys []string
	for k := range config.Commands {
		keys = append(keys, k)
	}

	return keys
}

func GetConfigAlias(command string) error {
	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	configPath := GetConfigFilePath()

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	// Declare a Config object to hold the parsed data
	var config Config

	// Parse the JSON string into the config object
	err = json.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		return fmt.Errorf("Error parsing JSON: %v", err)
	}

	// Extract command "q"
	if _, exists := config.Commands[command]; exists {
		fmt.Println("Running clix...")
	} else {
		return fmt.Errorf("Command not found in config")
	}

	return nil
}

func GetConfigAliasCommand(command string) string {
	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	configPath := GetConfigFilePath()

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Declare a Config object to hold the parsed data
	var config Config

	// Parse the JSON string into the config object
	err = json.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	// Return an empty slice of Hooks
	return config.Commands[command].Command
}

func GetConfigAliasHooks(command string, hooktype string) []Hooks {
	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	configPath := GetConfigFilePath()

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Declare a Config object to hold the parsed data
	var config Config

	// Parse the JSON string into the config object
	err = json.Unmarshal([]byte(fileContent), &config)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	// Returns correct hook
	if hooktype == "prehook" {
		return config.Commands[command].PreHooks
	} else if hooktype == "posthook" {
		return config.Commands[command].PostHooks
	}

	// Return an empty slice of Hooks
	return []Hooks{}
}

// CreateConfigFile creates a configuration file at the determined path
func CreateConfigFile() {
	// Get the config file path
	configPath := GetConfigFilePath()

	// Ensure the config file path is valid
	if configPath == "" {
		return
	}

	// Create the directory if it doesn't exist
	dir := filepath.Dir(configPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	if _, err := os.Stat(configPath); err == nil {
		// File exists, skip the rest of the function
		fmt.Println("File already exists, skipping the rest of the function.")
		return
	}

	// Create the configuration file
	file, err := os.Create(configPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Set write permissions (rw for owner, and read for others)
	// 0644: rw-r--r--
	err = os.Chmod(configPath, 0644)
	if err != nil {
		fmt.Println("failed to set file permissions:", err)
		return
	}

	// Example content to write to the config file
	// Create an empty Config struct
	config_json := Config{
		Commands: map[string]Command{
			"alias": {
				Command:     "",
				Description: "",
				PreHooks:    []Hooks{},
				PostHooks:   []Hooks{},
			},
		},
	}

	content, err := json.MarshalIndent(config_json, "", "  ")
	if err != nil {
		fmt.Println("error marshaling config to JSON:", err)
		return
	}

	// Write content to the config file
	_, err = file.WriteString(string(content))
	if err != nil {
		fmt.Println("Error writing to config file:", err)
		return
	}

	fmt.Printf("Config file created successfully at: %s\n", configPath)
}

func OpenFileInEditor(filePath string) error {
	var editorCommand string

	// Choose the editor based on the operating system
	if runtime.GOOS == "windows" {
		// On Windows, use notepad or any other editor installed
		editorCommand = "notepad"
	} else {
		// On Unix-like systems (Linux/macOS), use vim or nano
		editorCommand = "vim"
	}

	fmt.Println("here:", filePath)

	// Run the editor command to open the file
	cmd := exec.Command(editorCommand, filePath)
	// Pass on stdin and stdout from the calling program which, provided it was run from a terminal (likely for a command line program) will start vim for you and return control when the user has finished editing the file.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}

	return nil
}
