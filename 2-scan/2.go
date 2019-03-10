package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Print(a, b, "\n")
	fmt.Scanf("%d", &a)
	fmt.Scan(&b)
	fmt.Print(a, b, "\n")
}
