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
	config.Flags().BoolP("write", "w", false, "Opens up the config file for editing.")
}

var config = &cobra.Command{
	Use:   "config",
	Short: "Performs operations on the config file.",
	Long:  `Performs operations on the config file.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: returnConfigFile,
}

func returnConfigFile(cmd *cobra.Command, args []string) {

	// Read the config file using os.ReadFile instead of ioutil.ReadFile
	configPath := util.GetConfigFilePath()

	write_flag, err := cmd.Flags().GetBool("write")

	if write_flag {
		// Open the file in the chosen editor
		if err := util.OpenFileInEditor(configPath); err != nil {
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