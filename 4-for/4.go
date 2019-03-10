package main

import "fmt"

func main() {
	a := 0
	for i := 1; i <= 100; i++ {
		a += i
	}
	fmt.Println(a)
}
