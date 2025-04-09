package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"reflect"
)

// isZero checks if a reflect.Value is zero (empty)
func isZero(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

func runHooks(hooks []Hooks) {
	for _, hook := range hooks {

		if hook.RunCommand != "" && hook.RunFunction != "" {
			fmt.Println("Error! Only one of run_command or run_function should be set!")
			return
		}

		if hook.RunCommand != "" {
			runCommand(hook.RunCommand)
		}

		if hook.RunFunction != "" {
			runFunction(hook.RunFunction, hook.RunFunctionVars)
		}
	}
}

func runCommand(command string) {
	fmt.Println(command)

	// Run the editor command to open the file
	cmd := exec.Command(command)
	// Get the output from the command
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	// Print the output
	fmt.Println(string(output))
}

func runFunction(command string, vars []RunFunctionVars) {
	fmt.Println(command)

	switch command {
	case "set_env":
		setEnv(vars)
	case "load_env_file":
		loadEnvFile(vars)
	case "print_all_env":
		printAllEnv()
	case "hello_world":
		helloWorld(vars)
	case "run_python":
		runPython(vars)
	case "run_javascript":
		runJavaScript(vars)
	}
}
