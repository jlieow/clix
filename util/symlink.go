package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// GetGoModuleName reads the go.mod file in the current directory
// and returns the module name defined in it.
func GetGoModuleName() (string, error) {
	file, err := os.Open("go.mod")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", os.ErrNotExist
}

// GetGoPath returns the GOPATH environment variable.
// If not set, it falls back to the default ($HOME/go).
func GetGoPath() string {
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return gopath
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home + "/go"
}

func CreateSymLinksFromConfig() {

	// Check if the program is running as root (UID 0)
	if syscall.Geteuid() != 0 {
		fmt.Println("This program requires root privileges. Please run it with 'sudo'.")
		return
	}

	go_path := GetGoPath()

	clix_path := fmt.Sprintf("%s/bin/%s", go_path, StaticModuleName)

	src := clix_path

	var symlink_path string

	// Determine the correct path based on the OS
	switch runtime.GOOS {
	case "windows":
		symlink_path = filepath.Join(os.Getenv("USERPROFILE"), "bin")
	default:
		symlink_path = "/usr/local/bin"
	}

	// If User has defined a symlink_path in settings.json use it
	settings_symlink_path_value, err := GetSettingsValue("symlink_path")
	if err != nil {
		log.Fatalf(err.Error())
	}

	if settings_symlink_path_value != "" {
		symlink_path = settings_symlink_path_value
	}

	// Create target directory if needed (Windows case)
	if runtime.GOOS == "windows" {
		os.MkdirAll(symlink_path, 0755)
	}

	list_of_commands := GetListConfigAlias()

	log.Println(list_of_commands)

	for _, command := range list_of_commands {
		dst := filepath.Join(symlink_path, command)

		// Remove old symlink if it exists
		os.Remove(dst)

		// Creates symlinks
		create_err := os.Symlink(src, dst)
		if create_err != nil {
			log.Fatalf("Error creating symlink: %v", create_err)
		}

		log.Println("Symlink for " + command + " created successfully at path: " + dst)
	}
}
