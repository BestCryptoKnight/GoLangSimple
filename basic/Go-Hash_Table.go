package main

import "fmt"

func main() {
	var country map[int]string
	country = make(map[int]string)
	country[1] = "India"
	country[2] = "China"
	country[3] = "Japan"
	country[4] = "Austrlia"
	country[5] = "Africa"
	for i, j := range country {
		fmt.Printf("Key:  %d country: %s\n", i, j)
	}

}
