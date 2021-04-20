// Lists in golang
// Lists are called slices in golang

package main

import "fmt"

func main() {
	list := []int{}
	// Append to slice
	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	//Iterate over it
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}
	// Can take range
	fmt.Println(list[0:2])

}
