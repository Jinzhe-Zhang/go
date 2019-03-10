package main

import "fmt"

func main() {
	array := [...]int{1, 2, 45, 67, 543, 2, 1}
	var temp int
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}

		}
	}
	fmt.Print(array)
}
