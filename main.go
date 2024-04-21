package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Check if there are command-line arguments
	if len(os.Args) > 1 {
		// If there are arguments, pass them to quadchecker
		cmd := exec.Command("./quadchecker", os.Args[1:]...)
		cmd.Stdout = os.Stdout // Redirect standard output
		cmd.Stderr = os.Stderr // Redirect standard error
		err := cmd.Run()
		if err != nil {
			fmt.Println("Error running quadchecker:", err)
		}
	} else {
		// If there are no arguments, read from standard input and pass to quadchecker
		inputByte, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading from standard input:", err)
			return
		}

		inputString := string(inputByte)

		splitted_input_string := strings.Split(inputString, "\n")
		// Pass the input string to quadchecker
		cmd := exec.Command("./quadchecker", splitted_input_string[:len(splitted_input_string)-1]...)
		cmd.Stdout = os.Stdout // Redirect standard output
		cmd.Stderr = os.Stderr // Redirect standard error
		err = cmd.Run()
		if err != nil {
			fmt.Println("Error running quadchecker:", err)
		}
	}
}
