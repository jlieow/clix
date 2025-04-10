package config

import (
	"fmt"
	"log"
	"os"

	"clix/cmd"
	"clix/util"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(config)
	config.Flags().BoolP("edit", "e", false, "Opens up the config.json in a terminal for editing.")
	config.Flags().BoolP("gui", "g", false, "Open the config.json in a GUI for editing.")
}

var config = &cobra.Command{
	Use:   "config",
	Short: "Performs operations on the config file.",
	Long:  `Performs operations on the config file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: command,
}

func command(cmd *cobra.Command, args []string) {

	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	configPath := util.GetConfigFilePath()

	edit_flag, err := cmd.Flags().GetBool("edit")
	gui_flag, err := cmd.Flags().GetBool("gui")

	// If edit flag seen, open the file in editor based on the operating system
	if edit_flag {
		if err := util.OpenFileInEditor(configPath); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("File opened successfully for editing.")
		}
		return
	}

	// If GUI flag seen, open the file in a GUI for editing
	if gui_flag {
		config_file_path := "file://" + util.GetConfigFilePath()
		if err := util.OpenConfigJsonInGui(config_file_path); err != nil {
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
