package main

import (
	"fmt"
	"strings"
)

// Test data:

var puzzleInput = "2,8,0,4,5,0,6,0,0,0,5,0,6,9,0,0,0,7,0,0,0,1,8,0,0,3,0,9,0,0,0,0,4,1,0,0,7,2,0,9,0,1,0,5,6,0,0,4,7,0,0,0,0,2,0,6,0,0,4,9,0,0,0,8,0,0,0,1,5,0,6,0,0,0,3,0,7,6,0,2,1"

// Structs:

type number struct {
	row_num     int
	column_num  int
	section_num int
	value       string
	possibles   map[string]string
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
func convertToGrid(puzzleString []string) []number {

	// Convert to []number
	var row int = 1
	var column int = 1
	var iter int = 0
	var numberStructList []number

	//var grid [8][8]number

	for h := 0; h <= 6; h += 3 {

		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				for k := 1; k <= 3; k++ {
					//fmt.Print("[", row, ",", column, "]: ", j+h, "  ", "iter: ", iter, " value: ", puzzleString[iter], " ")

					var tempNumber number

					tempNumber.row_num = row
					tempNumber.column_num = column
					tempNumber.section_num = j + h

					if puzzleString[iter] == "0" {
						tempNumber.value = "0"
						tempNumber.possibles = map[string]string{
							"1": "1",
							"2": "2",
							"3": "3",
							"4": "4",
							"5": "5",
							"6": "6",
							"7": "7",
							"8": "8",
							"9": "9",
						}
					} else {
						tempNumber.value = puzzleString[iter]
						tempNumber.possibles = map[string]string{
							puzzleString[iter]: puzzleString[iter],
						}
					}

					numberStructList = append(numberStructList, tempNumber)

					iter++
					column++
				}
			}
			fmt.Print("\n")
			column = 1
			row++
		}
	}

	fmt.Println("numberStructList size: ", len(numberStructList))

	return numberStructList

}

func gridDisplay(gridArray []number) {

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

		fmt.Print(" ", gridArray[i].value)

	}

	fmt.Print("|\n |", breakline, "|")

}

func solveGrid(numberGrid []number) []number {

	var tempNumberGrid []number = numberGrid

	for i := range tempNumberGrid {

		if tempNumberGrid[i].value != "0" {

			var temp_row_num = tempNumberGrid[i].row_num
			var temp_column_num = tempNumberGrid[i].column_num
			var temp_section_num = tempNumberGrid[i].section_num
			var temp_value_num = tempNumberGrid[i].value

			// parse rows 	for h := 0; h <= 6; h += 3 {

			for j := range tempNumberGrid {

				if tempNumberGrid[j].value == "0" {
					if (tempNumberGrid[j].row_num == temp_row_num) || (tempNumberGrid[j].column_num == temp_column_num) || (tempNumberGrid[j].section_num == temp_section_num) {

						if val, ok := tempNumberGrid[j].possibles[temp_value_num]; ok {
							delete(tempNumberGrid[j].possibles, val)
						}
						if len(tempNumberGrid[j].possibles) == 1 {
							for k := range tempNumberGrid[j].possibles {
								tempNumberGrid[j].value = k
							}
						}

					}
				}

			}

			// fmt.Println("Row-Col-Sec-Val: ", temp_row_num, "-", temp_column_num, "-", temp_section_num, "-", temp_value_num)

		}

	}

	return tempNumberGrid

}

func main() {

	fmt.Println("Sudoku solver...")

	// string returned as []string
	var validatedNumbers []string = validatePuzzleInput(puzzleInput)

	// []string returned as []number (struct)
	var numberStructGrid []number = convertToGrid(validatedNumbers)

	// Text display of numbers
	gridDisplay(numberStructGrid)

	// Testing single sweeps:
	fmt.Println("\nTest 1:")
	numberStructGrid = solveGrid(numberStructGrid)
	gridDisplay(numberStructGrid)

	fmt.Println("\nTest 2:")
	numberStructGrid = solveGrid(numberStructGrid)
	gridDisplay(numberStructGrid)

	fmt.Println("\nTest 3:")
	numberStructGrid = solveGrid(numberStructGrid)
	gridDisplay(numberStructGrid)

	fmt.Println("\nTest 4:")
	numberStructGrid = solveGrid(numberStructGrid)
	gridDisplay(numberStructGrid)

}
