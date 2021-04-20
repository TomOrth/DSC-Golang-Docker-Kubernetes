package main

import "fmt"

// Structs are a way to construct a collection of items together for a similar purpose
// Point contains 2 ints to represent a coordinate point
// An important note is if you want the struct to be used outside its file, it has to start with an uppercase letter
type Point struct {
	x int
	y int
}

// By having a function in the form "func (struct instance) <func name>", you can attach a method to the struct like objects
func (p Point) Print() {
	fmt.Println(p.x, p.y)
}

// You can define methods for regular struct instances
// Or you can setup special methods for just pointers
func (p *Point) Update(x int, y int) {
	p.x = x
	p.y = y
}

func main() {
	point1 := Point{
		4,
		5,
	}
	point1.Print()
	// Example pointer
	point2 := &Point{
		6,
		7,
	}
	point2.Print()
	point2.Update(10, 11)
	point2.Print()
}
