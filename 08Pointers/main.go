package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	// not initalized - value is nil
	var intPtr *int

	fmt.Println(intPtr)

	// trying to access the actual value of an uninitialized pointer variable will result to nil pointer error
	// fmt.Println(*intPtr)

	// new initalizes with zero value and returns memory address of that value
	var intPtr2 = new(int)
	fmt.Println("Memory address -", intPtr2, "Actual value - ", *intPtr2)

	// use '&' to get memory address of a value
	var num = 23
	var numPtr = &num
	fmt.Println("Memory address of num -", numPtr, "and actual value", *numPtr)
}
