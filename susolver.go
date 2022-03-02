package main

import (
	"fmt"
	"strings"
)

// Test data:

var puzzleInput = "2,8,0,4,5,0,6,0,0,0,5,0,6,9,0,0,0,7,0,0,0,1,8,0,0,3,0,9,0,0,0,0,4,1,0,0,7,2,0,9,0,1,0,5,6,0,0,4,7,0,0,0,0,2,0,6,0,0,4,9,0,0,0,8,0,0,0,1,5,0,6,0,0,0,3,0,7,6,0,2,1"

func main() {

	fmt.Println("Sudoku solver...")

	var validatedNumbers []string = validatePuzzleInput(puzzleInput)

	convertToGrid(validatedNumbers)

}

// Make sure the inputted string is valid:
func validatePuzzleInput(puzzleString string) []string {

	numbers := strings.Split(puzzleString, ",")

	if len(numbers) != 81 {
		fmt.Println("Length of input is: ", len(numbers))
		panic("Invalid input length. Should be 81")
	}

	return numbers

}

// Converts a one line string into a 2d array
func convertToGrid(puzzleString []string) {

	// fmt.Println(puzzleString)
	gridDisplay(puzzleString)

}

func gridDisplay(gridArray []string) {

	var breakline string = "______________________"

	fmt.Println("\n", breakline, "_")

	for i := 0; i < len(gridArray); i++ {

		if (i%27 == 0) && i != 0 {
			fmt.Print("|\n |", breakline)
		}

		if (i%9 == 0) && i != 0 {
			fmt.Print("|\n")
		}

		if (i)%3 == 0 {
			fmt.Print(" |")
		}

		fmt.Print(" ", gridArray[i])

	}

	fmt.Print("|\n |", breakline, "|")

}
