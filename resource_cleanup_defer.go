package main

import (
	"fmt"
	"os"
)

// writeToFile creates a file, writes text to it, and uses defer to ensure cleanup.
func writeToFile(filename, content string) error {
	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	// Defer closing the file to ensure it happens after the function returns
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Printf("Error closing file: %v\n", closeErr)
		}
	}()

	// Write content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	// Optionally sync to ensure data is written to disk
	if err := file.Sync(); err != nil {
		return fmt.Errorf("failed to sync file: %w", err)
	}

	fmt.Printf("Successfully wrote to %s\n", filename)
	return nil
}

func main() {
	filename := "resource_cleanup_defer.txt"
	content := "Hello, Go! This is a test file created on Sunday, May 25, 2025.\n"

	err := writeToFile(filename, content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File operations completed.")
}
