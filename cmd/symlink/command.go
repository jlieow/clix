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

func command(cmd *cobra.Command, args []string) {
	util.CreateSymLinksFromConfig()
}
