package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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
// This is called by main.main(). It only needs to happen once to the rootCmd.
// Also creates a directory with config files
func Execute() {

	fmt.Println("Command used to run the program:", os.Args)

	alias := os.Args[0]
	if alias != "clix" {

		getConfigAlias(alias)

		preHooks := getConfigAliasHooks(alias, "prehook")
		runHooks(preHooks)

		command := os.Args
		command[0] = getConfigAliasCommand(alias)

		// Use exec.Command and expand the array using '...' to pass individual arguments
		// cmd := exec.Command(command[0], command[1:]...)
		cmd := exec.Command("terraform")

		// Capture the output of the command
		output, _ := cmd.Output()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// Print the output of the command
		fmt.Println(string(output))

		postHooks := getConfigAliasHooks(alias, "posthook")
		runHooks(postHooks)

		return
	}

	createConfigFile()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clix.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
