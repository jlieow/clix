package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Settings represents the structure for the settings file
type Settings struct {
	Symlink_Path string `json:"symlink_path"` // Specifies the path to create the symlink
	Setting1     string `json:"setting1"`
	Setting2     string `json:"setting2"`
}

// GetSettingsFilePath determines the platform-specific path for the config file
func GetSettingsFilePath() string {

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
		// For Linux/macOS, store in ~/.config/myapp/settings.json
		configDir = filepath.Join(homeDir, ".config", "clix", StaticSettingsJsonFileName)
	case "windows":
		// For Windows, store in C:\Users\<username>\AppData\Roaming\MyApp\settings.json
		configDir = filepath.Join(homeDir, "AppData", "Roaming", "CliX", StaticSettingsJsonFileName)
	default:
		fmt.Println("Unsupported OS")
		return ""
	}

	return configDir
}

func GetSettingsValue(key string) (string, error) {
	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	settingsPath := GetSettingsFilePath()

	fileContent, err := os.ReadFile(settingsPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Declare a Config object to hold the parsed data
	// Parse the JSON string into the config object
	var data map[string]string
	err = json.Unmarshal([]byte(fileContent), &data)
	if err != nil {
		return "", fmt.Errorf("Error parsing JSON: ", err)
	}

	val, ok := data[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}

	return val, nil
}

// CreateSettingsFile creates a configuration file at the determined path
func CreateSettingsFile() {
	// Get the config file path
	settingsPath := GetSettingsFilePath()

	// Ensure the config file path is valid
	if settingsPath == "" {
		return
	}

	// Create the directory if it doesn't exist
	dir := filepath.Dir(settingsPath)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	// File exists, skip the rest of the function
	if _, err := os.Stat(settingsPath); err == nil {
		// fmt.Println("settings.json already exists, skipping creation.")
		return
	}

	// Create the configuration file
	file, err := os.Create(settingsPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Set write permissions (rw for owner, and read for others)
	// 0644: rw-r--r--
	err = os.Chmod(settingsPath, 0644)
	if err != nil {
		fmt.Println("failed to set file permissions:", err)
		return
	}

	// Example content to write to the settings file
	// Create an empty Settings struct
	settings_json := Settings{
		Symlink_Path: "",
		Setting1:     "",
		Setting2:     "",
	}

	content, err := json.MarshalIndent(settings_json, "", "  ")
	if err != nil {
		fmt.Println("error marshaling settings to JSON:", err)
		return
	}

	// Write content to the settings file
	_, err = file.WriteString(string(content))
	if err != nil {
		fmt.Println("Error writing to settings file:", err)
		return
	}

	fmt.Printf("Settings file created successfully at: %s\n", settingsPath)
}
