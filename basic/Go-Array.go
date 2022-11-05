/*
Go also Has array Like other programing language which can store fixed size sequential collection of elements
of the same data-type.package Basic

Instead of declaring individual variable, sunch as num0, num1, num2..., You declare one array variable such as
Number.
Element can be accessed by indexing and slicing.
 indexing start from `0(Zero)` Like Python.
*/
package main

import "fmt"

func main() {
	/* n is an array of 10 integers */
	var n [10]int
	var i, j int

	/* initialize element of array n to 0*/

	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}
