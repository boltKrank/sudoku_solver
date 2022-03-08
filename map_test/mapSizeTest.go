package main

import "fmt"

func main() {

	var myMap9 map[string]string = map[string]string{
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

	var myMap1 map[string]string = map[string]string{
		"2": "2",
	}

	fmt.Println("myMap9 len() is: ", len(myMap9))
	fmt.Println("myMap1 len() is: ", len(myMap1))

	fmt.Println("Contents of myMap9 are: ", myMap9)
	fmt.Println("Contents of myMap1 are: ", myMap1)

	var map1key string

	for k := range myMap1 {
		map1key = k
	}

	fmt.Println("Map1key: ", map1key)

}
