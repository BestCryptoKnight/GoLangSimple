/* Interfaces is another Data-Type in GoLang which represents a set of method signatures.The Struct data-type
implements these interfaces to have method definations for the method signature of the interfaces. */

package main

import (
	"fmt"
)

/* define an interface */
type Circle struct {
	x, y, radius float64
}

/* Define a rectangle */
type Rectangle struct {
	width, height float64
}

/* Define a method for circle */
func (circle Circle) area() float64 {
	return math.pi * circle.radius * circle.radius
}

/* Define a method for rectangle */
func getArea(shape Shape) float64 {
	return shape.area()
}
func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Printf("Circle Area: %f \n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}
