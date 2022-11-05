/* Brute Force to find Fibbonacci Number By Recursion
Fibonacci is series which start from 0,1 and next number formed by addition of previous two
Number.
0, 1, 1, 2, 3, 5, 8, 13 ,21, 34, 55....
*/
package main

import "fmt"

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
func main() {

	fmt.Println(Fibonacci(10))
}
