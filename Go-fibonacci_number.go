/* Brute Force to find Fibbonacci Number By Recursion */
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
