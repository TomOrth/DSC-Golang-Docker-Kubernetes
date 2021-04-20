// Go files need a package at the top
// and files with the main method need to have package main at the top
package main

import "fmt"

// This is how you declare a variable in go
// It reads as "I am making a variable called x that is a 64 bit int with value 4"
var x int64 = 4

// Here is some common types besides int64
var a string = "Hello"
var b float64 = 450
var c bool = false

func main() {
	// This is an alternative form for infering the type from
	y := 4.5
	fmt.Println(a, b, c, x, y)
}
