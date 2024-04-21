package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Check if the number of command-line arguments is correct
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <directory>")
		os.Exit(1)
	}

	// Get the directory path from the command-line argument
	directory := os.Args[1]

	// Walk through the directory and its subdirectories
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		// Check if the file is a regular file
		if err == nil && !info.IsDir() {
			// Make the file executable
			err = os.Chmod(path, info.Mode()|0111)
			if err != nil {
				fmt.Printf("Error making %s executable: %v\n", path, err)
			} else {
				fmt.Printf("Made %s executable.\n", path)
			}
		}
		return nil
	})
	// Check for errors during filepath.Walk
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("All files made executable.")
}
