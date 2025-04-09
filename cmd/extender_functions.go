package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func helloWorld(vars []RunFunctionVars) {

	// Convert to map for easy lookup
	argMap := make(map[string]string)
	for _, v := range vars {
		argMap[v.Key] = v.Value
	}

	// Extract values in function parameter order
	name := argMap["name"]
	age := argMap["age"]

	fmt.Println("Hello " + name + " you are currently " + age)
}

func loadEnvFile() error {
	// Load the .env file from the current working directory (where the app is called)
	err := godotenv.Load() // This will automatically load the .env file from the current directory
	if err != nil {
		return fmt.Errorf("Error loading .env file")
	}
	return nil
}

func printAllEnv() {
	// Get all environment variables
	envVars := os.Environ()

	// Print each environment variable
	for _, envVar := range envVars {
		fmt.Println(envVar)
	}
}
