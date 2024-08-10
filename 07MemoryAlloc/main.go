package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Memory Allocation")

	// new() - allocates memory but no INIT, returns a memory address, zeroed storage
	// make() - allocates memory and INIT, returns memory address, non-zeroed storage

	fmt.Println(runtime.NumCPU())
}
