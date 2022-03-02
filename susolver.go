package main

import "fmt"

// Test data:

var puzzleInput = "2,8,0,4,5,0,6,0,0,0,5,0,6,9,0,0,0,7,0,0,0,1,8,0,0,3,0,9,0,0,0,0,4,1,0,0,7,2,0,9,0,1,0,5,6,0,0,4,7,0,0,0,0,2,0,6,0,0,4,9,0,0,0,8,0,0,0,1,5,0,6,0,0,0,3,0,7,6,0,2,1;"

func main() {

	fmt.Println("Sudoku solver...")

	// Check inputted data for validity
	// validatePuzzleInput(puzzleInput)

	convertToGrid(puzzleInput)

}

// Make sure the inputted string is valid:
func validatePuzzleInput(puzzleString string) {

}

// Converts a one line string into a 2d array
func convertToGrid(puzzleString string) {

	fmt.Println(puzzleString)

}
