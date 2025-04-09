package cmd

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

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
