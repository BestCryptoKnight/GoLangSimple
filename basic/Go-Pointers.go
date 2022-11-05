/*
Every variable is a a memory location and every memory location has its address defined which can be accessed
using ampersand and opeator, which denotes an addresss in memory.Cosider the following example, which will print
the address of the variables defined -
*/

package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("Address of a Variable %x\n", &a)
}
/*
What are Pointers :
					A pointer is a variable whose value is the address of another variable,
					
*/