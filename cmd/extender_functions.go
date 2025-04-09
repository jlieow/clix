package cmd

import (
	"fmt"

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
