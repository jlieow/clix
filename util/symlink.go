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
	module_name, get_module_name_err := GetGoModuleName()
	if get_module_name_err != nil {
		log.Fatal(get_module_name_err)
	}

	clix_path := fmt.Sprintf("%s/bin/%s", go_path, module_name)

	src := clix_path

	var targetDir string

	// Determine the correct path based on the OS
	switch runtime.GOOS {
	case "windows":
		targetDir = filepath.Join(os.Getenv("USERPROFILE"), "bin")
	default:
		targetDir = "/usr/local/bin"
	}

	// Create target directory if needed (Windows case)
	if runtime.GOOS == "windows" {
		os.MkdirAll(targetDir, 0755)
	}

	list_of_commands := GetListConfigCommand()

	log.Println(list_of_commands)

	for _, command := range list_of_commands {
		dst := filepath.Join(targetDir, command)

		// Remove old symlink if it exists
		os.Remove(dst)

		// Creates symlinks
		create_err := os.Symlink(src, dst)
		if create_err != nil {
			log.Fatalf("Error creating symlink: %v", create_err)
		}

		log.Println("Symlink for " + command + " created successfully")
	}
}
