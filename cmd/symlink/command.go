package symlink

import (
	"clix/cmd"
	"clix/util"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(symlinkCmd)
}

var symlinkCmd = &cobra.Command{
	Use:   "symlink",
	Short: "Creates symlinks based on the config file.",
	Long:  `Creates symlinks based on the config file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: command,
}

// // GetGoModuleName reads the go.mod file in the current directory
// // and returns the module name defined in it.
// func GetGoModuleName() (string, error) {
// 	file, err := os.Open("go.mod")
// 	if err != nil {
// 		return "", err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		line := strings.TrimSpace(scanner.Text())
// 		if strings.HasPrefix(line, "module ") {
// 			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return "", err
// 	}

// 	return "", os.ErrNotExist
// }

// // GetGoPath returns the GOPATH environment variable.
// // If not set, it falls back to the default ($HOME/go).
// func GetGoPath() string {
// 	if gopath := os.Getenv("GOPATH"); gopath != "" {
// 		return gopath
// 	}
// 	home, err := os.UserHomeDir()
// 	if err != nil {
// 		return ""
// 	}
// 	return home + "/go"
// }

func command(cmd *cobra.Command, args []string) {

	util.CreateSymLinksFromConfig()
}
