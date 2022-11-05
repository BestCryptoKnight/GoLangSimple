/*
The range keyword is used in for loop to iterate over itemsof an array, slice channel or map.
with array and slices , it returns the index of the item as integer.
*/

package main

import "fmt"

func main() {
	/* Create a slices */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}

	/* print the numbers */
	for i := range numbers {
		fmt.Println("At index", i, "and Number is", numbers[i])
	}

	/* create a map like dict in python */
	countryCapital := map[string]string{"India": "New Delhi", "Russia": "Moscow", "Itlay": "Rome", "Japan": "Tokyo"}

	/* print map using keys */

	/* for country := range countryCapital {
	fmt.Println("Capital of ", country, "is", countryCapital[country])
	*/

	for country, capital := range countryCapital {
		fmt.Println("Capital of ", country, "is", capital)

	}
}
