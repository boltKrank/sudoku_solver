package main

import (
	"fmt"
	"strings"
)

type number struct {
	row_num     int
	column_num  int
	section_num int
	value       string
	possibles   map[string]string
}

func main() {

	var testNumbers = "1,2,3,4,0,6,7,8,9"
	numberList := strings.Split(testNumbers, ",")

	fmt.Println("Input list: ", numberList)

	var numberStructList []number

	for n := range numberList {

		var tempNumber number

		tempNumber.row_num = 0
		tempNumber.column_num = 0
		tempNumber.section_num = 0

		if numberList[n] == "0" {
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
			tempNumber.value = numberList[n]
			tempNumber.possibles = map[string]string{
				numberList[n]: numberList[n],
			}
		}

		numberStructList = append(numberStructList, tempNumber)

	}

	fmt.Println("Number structs list: ")
	fmt.Println(numberStructList)

	fmt.Println("Broadcast until 0 disappears..")

	fmt.Println("Sweep 1")

	for i := range numberStructList {

		if numberStructList[i].value != "0" {

			var broadcastNumber = numberStructList[i].value

			for j := range numberStructList {

				if numberStructList[j].value == "0" {
					if val, ok := numberStructList[j].possibles[broadcastNumber]; ok {
						delete(numberStructList[j].possibles, val)
					}
					if len(numberStructList[j].possibles) == 1 {
						for k := range numberStructList[j].possibles {
							numberStructList[j].value = k
						}
					}
				}

			}

		}

	}

	fmt.Println("Number structs list: ")
	fmt.Println(numberStructList)

}
