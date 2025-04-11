package settings

import (
	"fmt"
	"log"
	"os"

	"clix/cmd"
	"clix/util"

	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(settings)
	settings.Flags().BoolP("edit", "e", false, "Opens up the settings.json in a terminal for editing.")
	settings.Flags().BoolP("gui", "g", false, "Open the settings.json in a GUI for editing.")
	settings.Flags().StringP("setting", "s", "", "Returns the setting as a string.")
}

var settings = &cobra.Command{
	Use:   "settings",
	Short: "Performs operations on the settings file.",
	Long:  `Performs operations on the settings file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: command,
}

func command(cmd *cobra.Command, args []string) {

	// Read the settings file using os.ReadFile instead of ioutil.ReadFile
	settingsPath := util.GetSettingsFilePath()

	edit_flag, err := cmd.Flags().GetBool("edit")
	gui_flag, err := cmd.Flags().GetBool("gui")
	setting_flag, err := cmd.Flags().GetString("setting")

	// If edit flag seen, open the file in editor based on the operating system
	if edit_flag {
		if err := util.OpenFileInEditor(settingsPath); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("File opened successfully for editing.")
		}
		return
	}

	// If GUI flag seen, open the file in a GUI for editing
	if gui_flag {
		config_file_path := "file://" + util.GetConfigFilePath()
		settings_file_path := "file://" + util.GetSettingsFilePath()
		util.Gui(config_file_path, settings_file_path, util.StaticSettingsJson)

		// if err := util.OpenSettingsJsonInGui(settings_file_path); err != nil {
		// 	fmt.Println("Error:", err)
		// } else {
		// 	fmt.Println("File opened successfully for editing.")
		// }
		return
	}

	// If GUI flag seen, open the file in a GUI for editing
	if setting_flag != "" {

		setting_value, err := util.GetSettingsValue(setting_flag)
		if err != nil {
			log.Fatalf(err.Error())
		}

		if setting_value == "" {
			setting_value = "Null"
		}

		fmt.Println(setting_value)
		return
	}

	fileContent, err := os.ReadFile(settingsPath)
	if err != nil {
		log.Fatalf("Error reading settings file: %v", err)
	}

	fmt.Printf(string(fileContent))
}
