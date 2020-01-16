/* Type casting is a way to convert a variable from one data type to another data type. If User want to store a
long value into simple interger than you can type cast long to int.  */

package main

import "fmt"

func main() {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)

	fmt.Printf("Mean :%f\n", mean)
}

// Value of mean : 3.400000
