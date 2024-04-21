package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// Function shape_check checks if the input is valid for a quad function and returns the dimensions of the array if so.
func shape_check(input []string) (bool, []int) {
	base_len := len(input[0])
	// Check if the lines of the input have the same number of characters.
	for _, v := range input {
		if len(v) != base_len {
			return false, []int{0, 0}
		}
	}
	// If everything matches, the input has the shape of a quad function.
	return true, []int{base_len, len(input)}
}

func main() {
	// Get command line arguments.
	_inputString := os.Args[1:]
	// Check the shape of the input.
	fit, cord := shape_check(_inputString)
	// Join input elements into a string with newline.
	inputString := strings.Join(_inputString, "\n")

	if fit {
		quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
		matches := []string{}

		// Check if the output of each quad function matches the input.
		for _, quad := range quads {
			output := captureOutput(quad, cord[0], cord[1])
			if strings.TrimSpace(output) == strings.TrimSpace(inputString) {
				matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad, cord[0], cord[1]))
			}
		}

		if len(matches) == 0 {
			fmt.Println("Not a quad function")
		} else {
			// Print the matches, sorted.
			sort.Strings(matches)
			fmt.Println(strings.Join(matches, " || "))
		}
	} else {
		fmt.Println("Not a quad function")
	}
}

// Function captureOutput calls the quad functions with the shape and returns their output.
func captureOutput(quad string, x, y int) string {
	cmd := exec.Command("./"+quad, fmt.Sprintf("%d", x), fmt.Sprintf("%d", y))
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running command:", err)
		return ""
	}
	return string(output)
}
