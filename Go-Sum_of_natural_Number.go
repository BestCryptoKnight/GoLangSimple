// Sum of n natural Number

package main

import "fmt"

func main() {
	var n, sum int
	fmt.Print("Enter a positive Number :")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Printf("Sum of %d Natural Numbers are %d:", n, sum)

}
