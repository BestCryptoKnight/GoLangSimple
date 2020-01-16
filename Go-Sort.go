/* Golang Program with example of array sort Functions */
package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{50, 90, 30, 20, 10}
	fmt.Println("Unsorted Array :", num) // Unsorted Array
	sort.Ints(num)
	fmt.Println("Sorted Array   :", num) //Sorted Array

	text := []string{"Us", "Uk", "India", "Japan", "China"}
	fmt.Println(text) // Unsorted array
	sort.Strings(text)
	fmt.Println(text) // Sorted Array
}
