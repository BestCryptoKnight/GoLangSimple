/*
Go- Programe For Boolean Data Type
*/
package main

import "fmt"

func main() {
	str1 := "PythoN"
	str2 := "Python"
	str3 := "PythoN"

	result1 := str1 == str2
	result2 := str1 == str3

	fmt.Println(result1) //False
	fmt.Println(result2) // True

	/* fmt.printf(result1,result2)
	To Check Data Type Of Result1 and Result 2
	*/
}
