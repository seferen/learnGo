package main

import (
	"fmt"
)

func main() {

	var workArray [10]uint8

	for i := 0; i < 10; i++ {

		fmt.Scan(&workArray[i])

	}
	for i := 0; i < 3; i++ {
		var a, b uint8
		fmt.Scan(&a, &b)
		t := workArray[a]
		workArray[a] = workArray[b]
		workArray[b] = t
	}

	for i := 0; i < len(workArray); i++ {
		if i != (len(workArray) - 1) {
			fmt.Printf("%d ", workArray[i])
		} else {
			fmt.Printf("%d", workArray[i])
		}

	}
}
