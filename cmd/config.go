package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(config)
	config.Flags().BoolP("write", "w", false, "Opens up the config file for editing.")
}

var config = &cobra.Command{
	Use:   "config",
	Short: "Creates symlinks based on the config file.",
	Long:  `Creates symlinks based on the config file located at xxx`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: returnConfigFile,
}

// Task represents a task to perform before running the command
type Hooks struct {
	Task       string            `json:"task"`
	EnvVars    map[string]string `json:"env_vars,omitempty"`
	Message    string            `json:"message,omitempty"`
	FilePath   string            `json:"file_path,omitempty"`
	RunCommand string            `json:"run command,omitempty"`
}

// Command represents the structure for each command to execute
type Command struct {
	Alias       string  `json:"alias"`
	Command     string  `json:"command"`
	Description string  `json:"description"`
	PreHooks    []Hooks `json:"prehooks"`
	PostHooks   []Hooks `json:"posthooks"`
}

// Config represents the top-level structure with commands
type Config struct {
	Commands map[string]Command `json:"commands"`
}

// Execute the task
func executeTask(task Hooks) {
	switch task.Task {
	case "set_env":
		// Set environment variables
		for key, value := range task.EnvVars {
			os.Setenv(key, value)
			fmt.Printf("Setting environment variable: %s=%s\n", key, value)
		}
	case "log_message":
		// Log a message
		fmt.Println(task.Message)
	case "check_file":
		// Check if file exists
		if _, err := os.Stat(task.FilePath); os.IsNotExist(err) {
			fmt.Printf("Error: file not found: %s\n", task.FilePath)
		} else {
			fmt.Printf("File exists: %s\n", task.FilePath)
		}
	default:
		fmt.Println("Unknown task:", task.Task)
	}
}

// getConfigFilePath determines the platform-specific path for the config file
func getConfigFilePath() string {
	// Get the home directory of the current user
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		return ""
	}

	homeDir := usr.HomeDir
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

func returnConfigFile(cmd *cobra.Command, args []string) {

	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	configPath := getConfigFilePath()

	write_flag, err := cmd.Flags().GetBool("write")

	if write_flag {
		// Open the file in the chosen editor
		if err := openFileInEditor(configPath); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("File opened successfully for editing.")
		}
		return
	}

	fileContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	fmt.Printf(string(fileContent))
}

// createConfigFile creates a configuration file at the determined path
func createConfigFile() {
	// Get the config file path
	configPath := getConfigFilePath()

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
			"sample_command": {
				Alias:       "",
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

func openFileInEditor(filePath string) error {
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

func main() {
	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	configFile := "config.json"
	fileContent, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Unmarshal the JSON into a Config struct
	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Process each command
	for _, cmd := range config.Commands {
		// Execute all tasks before running the actual command
		for _, task := range cmd.PreHooks {
			executeTask(task)
		}

		// After tasks are executed, run the CLI command
		fmt.Printf("Executing command: %s\n", cmd.Command)
		// Here you would actually run the command, for example using exec.Command
	}
}
