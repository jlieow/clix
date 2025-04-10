package gui

import (
	"clix/cmd"
	"clix/util"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(guiCmd)
}

var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Launches CliX GUI",
	Long:  `Launches CliX GUI.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: command,
}

func command(cmd *cobra.Command, args []string) {

	config_file_path := "file://" + util.GetConfigFilePath()
	util.OpenConfigJsonInTabs(config_file_path)

}
