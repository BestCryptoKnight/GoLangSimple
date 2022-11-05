package main

import "fmt"

func main() {
	fmt.Println("Enter 3 Numbers For Compare:")
	var first int
	fmt.Scan(&first)
	var second int
	fmt.Scan(&second)
	var third int
	fmt.Scan(&third)
	/* Check For Conditions */
	if first >= second && first >= third {
		fmt.Println(first, " Is the largest Among Three Number.")
	}
	if second >= first && second >= third {
		fmt.Println(second, " Is the largest Among Three Number")
	}
	if third >= first && third >= second {
		fmt.Println(third, "Is the the largest among three number")
	}
}
