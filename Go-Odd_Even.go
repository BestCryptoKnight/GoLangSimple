package main

import "fmt"

func main() {
	fmt.Print("Enter number:")
	var n int
	fmt.Scanln(&n)
	if n%2 == 0 {
		fmt.Printf("%d is Even Number", n)
	} else {
		fmt.Printf("%d Is Odd Number", n)
	}
}
