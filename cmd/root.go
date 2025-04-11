package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"clix/util"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "clix",
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true, // hides cmd as described here: https://github.com/spf13/cobra/blob/main/site/content/completions/_index.md#adapting-the-default-completion-command
		// DisableDefaultCmd: true, // removes cmd
	},
	Short: "Command line extender.",
	Long: `Designed to enhance and extend the functionality of command-line interfaces (CLI). 
CliX empowers developers, system administrators, and power users by allowing them to define, 
customize, and chain commands with ease.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
// Also creates a directory with config files
func Execute() {

	alias := os.Args[0]
	if alias != "clix" {

		err := util.GetConfigAlias(alias)
		if err != nil {
			fmt.Println(err)
			return
		}

		preHooks := util.GetConfigAliasHooks(alias, "prehook")
		util.RunHooks(preHooks)

		command := os.Args
		command[0] = util.GetConfigAliasCommand(alias)

		// Use exec.Command and expand the array using '...' to pass individual arguments
		cmd := exec.Command(command[0], command[1:]...)

		cmd.Stdin = os.Stdin   // Connects to terminal input
		cmd.Stdout = os.Stdout // Outputs to terminal
		cmd.Stderr = os.Stderr // Error output to terminal
		cmd.Run()

		// err = cmd.Run()
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// // Capture the output of the command
		// output, _ := cmd.Output()
		// // if err != nil {
		// // 	log.Fatal(err)
		// // }

		// // Print the output of the command
		// fmt.Println(string(output))

		postHooks := util.GetConfigAliasHooks(alias, "posthook")
		util.RunHooks(postHooks)

		return
	}

	util.CreateConfigFile()
	util.CreateSettingsFile()

	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clix.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
