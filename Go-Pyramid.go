/* Go Programe to print Pyramid to change the height of pyramid change the value of
row.

*/

package main

import "fmt"

func main() {
	var rows int = 5 // Number of rows in pyramid
	var k int = 0
	for i := 1; i <= rows; i++ {
		k = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("*")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println(" ")
	}
}
