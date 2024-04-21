package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./quadA width height")
		return
	}

	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid width")
		return
	}

	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid height")
		return
	}

	QuadE(width, height)
}

func QuadE(x, y int) {
	if Validate(x, y) {
		PrintTop(x, "A", "C", "B")
		if y > 2 {
			PrintBody(x, y, "B")
		}
		if y > 1 {
			PrintBottom(x, y, "C", "A", "B")
		}
	}
}

func Validate(x int, y int) bool {
	if x <= 0 || y <= 0 {
		return false
	}
	return true
}

func PrintTop(x int, leftCorner, rightCorner, midChar string) {
	for i := 0; i < x; i++ {
		if i == 0 {
			fmt.Print(leftCorner)
		} else if i == (x - 1) {
			fmt.Print(rightCorner)
		} else if i > 0 && i < (x-1) {
			fmt.Print(midChar)
		}
	}
	fmt.Println()
}

func PrintBody(x, y int, side string) {
	for i := 0; i < (y - 2); i++ {
		if x > 1 {
			fmt.Print(side)
		} else {
			fmt.Println(side)
		}
		if x > 1 {
			for j := 0; j < (x - 2); j++ {
				fmt.Print(" ")
			}
			fmt.Println(side)
		}
	}
}

func PrintBottom(x int, y int, leftCorner, rightCorner, midChar string) {
	if y == 1 {
		return
	}
	for i := 0; i < x; i++ {
		if i == 0 {
			fmt.Print(leftCorner)
		} else if i == (x - 1) {
			fmt.Print(rightCorner)
		} else if i > 0 && i < (x-1) {
			fmt.Print(midChar)
		}
	}
	fmt.Println()
}
