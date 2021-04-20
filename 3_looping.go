package main

import "fmt"

func main() {
	// Similar for loop syntax to other languages
	for i := 0; i < 5; i++ {
		fmt.Println("Hello!")
	}

	// There is no while loop keyword. The for construct can be used

	j := 1
	for j < 9 || j%2 != 0 {
		fmt.Println("We are in a while loop this will loop for a bit")
		j++
	}
	fmt.Println("We are out of the loop")
}
