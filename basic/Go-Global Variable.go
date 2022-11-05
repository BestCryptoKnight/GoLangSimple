/*
There are two type of variable in Go .
1. Global Vaaribale --> Outside the function or Block
2. Local Variable   --> inside the function or block
*/
package main

import "fmt"

/* global variable declaration */
var res int

func main() {
	/* local Variable declaration */
	var a, b int

	a = 10
	b = 30
	res = a + b
	fmt.Printf("Sum of %d and %d are %d.", a, b, res)
}

/*
 A program can have the same name for local and globa variable but local varible inside the fucntion takes
prefence.
*/
