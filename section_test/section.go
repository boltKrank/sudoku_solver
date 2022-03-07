package main

import (
	"fmt"
	"strings"
)

var puzzleInput = "2,8,0,4,5,0,6,0,0,0,5,0,6,9,0,0,0,7,0,0,0,1,8,0,0,3,0,9,0,0,0,0,4,1,0,0,7,2,0,9,0,1,0,5,6,0,0,4,7,0,0,0,0,2,0,6,0,0,4,9,0,0,0,8,0,0,0,1,5,0,6,0,0,0,3,0,7,6,0,2,1"

type number struct {
	row_num     int
	column_num  int
	section_num int
	value       string
	possibles   []string //Make this a MAP
}

func main() {

	var row int = 1
	var column int = 1
	var iter int = 0

	numbers := strings.Split(puzzleInput, ",")

	var grid [8][8]number

	for h := 0; h <= 6; h += 3 {

		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				for k := 1; k <= 3; k++ {
					fmt.Print("[", row, ",", column, "]: ", j+h, "  ")
					grid[row-1][column-1].row_num = row - 1
					grid[row-1][column-1].column_num = column - 1
					grid[row-1][column-1].section_num = j + h
					grid[row-1][column-1].value = numbers[iter]
					iter++
					row++
				}
			}
			fmt.Print("\n")
			row = 1
			column++
		}
	}
}
