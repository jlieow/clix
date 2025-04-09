package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(symlinkCmd)
}

var symlinkCmd = &cobra.Command{
	Use:   "symlink",
	Short: "Creates symlinks based on the config file.",
	Long:  `Creates symlinks based on the config file located at xxx`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: createSymLinksFromConfig,
}

// GetGoModuleName reads the go.mod file in the current directory
// and returns the module name defined in it.
func getGoModuleName() (string, error) {
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
func getGoPath() string {
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		return gopath
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return home + "/go"
}

func createSymLinksFromConfig(cmd *cobra.Command, args []string) {

	go_path := getGoPath()
	module_name, get_module_name_err := getGoModuleName()
	if get_module_name_err != nil {
		log.Fatal(get_module_name_err)
	}

	clix_path := fmt.Sprintf("%s/bin/%s", go_path, module_name)

	src := clix_path
	// dst := "/usr/local/bin/q"

	list_of_commands := getListConfigCommand()

	for _, command := range list_of_commands {
		dst := "/usr/local/bin/" + command

		// Checks if file exists
		// If it does, removes symlink to remove any previous links
		if _, err := os.Stat(dst); err == nil {
			remove_err := os.Remove(dst)
			if remove_err != nil {
				log.Fatal(remove_err)
			}
		}

		// Creates symlinks
		create_err := os.Symlink(src, dst)
		if create_err != nil {
			log.Fatalf("Error creating symlink: %v", create_err)
		}

		log.Println("Symlink for " + command + " created successfully")
	}
}

func createSymLinks() {

	fmt.Println(getGoPath())
	fmt.Println(getGoModuleName())

	go_path := getGoPath()
	module_name, get_module_name_err := getGoModuleName()
	if get_module_name_err != nil {
		log.Fatal(get_module_name_err)
	}

	clix_path := fmt.Sprintf("%s/bin/%s", go_path, module_name)

	src := clix_path
	dst := "/usr/local/bin/q"

	// Removes symlink to remove any previous links
	remove_err := os.Remove(dst)
	if remove_err != nil {
		log.Fatal(remove_err)
	}

	// Creates symlinks
	create_err := os.Symlink(src, dst)
	if create_err != nil {
		log.Fatalf("Error creating symlink: %v", create_err)
	}

	log.Println("Symlink created successfully")
}
