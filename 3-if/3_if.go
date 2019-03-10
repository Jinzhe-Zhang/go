// if
package main

import "fmt"
func main() {
    var age int;
	fmt.Scan(&age);
	if age>=18 {
	fmt.Println("You are an adult.")
	}else{
	fmt.Println("You are a child.")
	}
}
