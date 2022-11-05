/* NumCPU returns the number of logical CPUs usable by the current process */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
}
