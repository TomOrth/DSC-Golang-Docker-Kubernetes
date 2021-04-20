// Notes: add more on private vs public functions
package main

import "fmt"

// Functions are preceded with the keyword "func"
// We have the return type after the function name, similar to variable declartion
func doSomeWork() int {
	var x int = 0
	for i := 0; i < 5; i++ {
		x += i
	}
	return x
}

// Types must be specified for the function parameters
func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

// Functions can return multiple values, specified as a pair of elements
func passAndReturn(x string, y int) (string, int) {
	return x, y
}

func main() {
	fmt.Println(doSomeWork())
	fmt.Println(min(5, 6))
	fmt.Println(passAndReturn("hello", 6))
}
