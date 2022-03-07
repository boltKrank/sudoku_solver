package main

import (
	"fmt"
	"strings"
)

// TODO: Make pointers for each struct, make an array of such pointers, and parse those points

// Test data:

var puzzleInput = "2,8,0,4,5,0,6,0,0,0,5,0,6,9,0,0,0,7,0,0,0,1,8,0,0,3,0,9,0,0,0,0,4,1,0,0,7,2,0,9,0,1,0,5,6,0,0,4,7,0,0,0,0,2,0,6,0,0,4,9,0,0,0,8,0,0,0,1,5,0,6,0,0,0,3,0,7,6,0,2,1"

// Constants:

const width_size = 9
const height_size = 9
const section_size = 9
const row_type = 1
const column_type = 2
const section_type = 3

// Structs:

type number struct {
	row_num     int
	column_num  int
	section_num int
	finished    bool
	possibles   []string
	// members in row
	// members in column
	// members in section
}

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

	// Make it an array of number structs

	// Show row (1-9)
	fmt.Println(showPart(puzzleString, row_type, 1))

	// Show column (1-9)
	fmt.Println(showPart(puzzleString, column_type, 1))

	// Show section (1-9)
	fmt.Println(showPart(puzzleString, section_type, 1))

	gridDisplay(puzzleString)

}

//parameters: Array, row, column, section.
func showPart(puzzleString []string, partType int, partArea int) string {

	var parType = ""

	if partType == row_type {
		parType = "ROW"
	} else if partType == column_type {
		parType = "COLUMN"
	} else if partType == section_type {
		parType = "SECTION"
	}

	return parType

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
